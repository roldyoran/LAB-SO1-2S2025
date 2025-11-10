package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Configurar cliente
	addr := func() string { if v := os.Getenv("VALKEY_SERVICE_URL"); v != "" { return v }; return "localhost:6379" }()

	var rdb *redis.Client
	if len(addr) >= 8 && addr[:8] == "redis://" {
		opt, err := redis.ParseURL(addr)
		if err != nil {
			log.Fatal("URL inválida en VALKEY_SERVICE_URL:", err)
		}
		rdb = redis.NewClient(opt)
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "", // sin contraseña por defecto
			DB:       0,  // base de datos por defecto
		})
	}

	// Verificar conexión
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error conectando a Valkey:", err)
	}
	fmt.Println("Conexión exitosa:", pong)

	// Ejemplos de operaciones
	ejemploSetGet(rdb)
	ejemploHashes(rdb)
}

func ejemploSetGet(rdb *redis.Client) {
	fmt.Println("\n--- SET/GET - Weather Tweet Status ---")
	
	// Almacenar temperatura actual de Guatemala
	if err := rdb.Set(ctx, "temperature:guatemala", "25", 0).Err(); err != nil {
		log.Fatal(err)
	}

	temp, err := rdb.Get(ctx, "temperature:guatemala").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temperatura en Guatemala:", temp, "°C")
}


func ejemploHashes(rdb *redis.Client) {
	fmt.Println("\n--- HASHES - Datos Meteorológicos Actuales ---")
    
	// helper para guardar lectura: actualiza el hash (estado actual)
	saveReading := func(municipality string, temp int, hum int, weather string) {
		ts := time.Now().Format("2006-01-02 15:04:05")

		// actualizar estado actual en HASH
		if err := rdb.HSet(ctx,
			fmt.Sprintf("municipality:%s", municipality),
			"name", municipality,
			"temperature", temp,
			"humidity", hum,
			"weather", weather,
			"last_update",  ts,
		).Err(); err != nil {
			log.Fatal(err)
		}
	}

	// Almacenar datos meteorológicos de Guatemala (estado + historial)
	saveReading("guatemala", 99, 26, "sunny")

	// Almacenar datos meteorológicos de Mixco
	saveReading("mixco", 22, 70, "cloudy")

	// Almacenar datos meteorológicos de Amatitlán
	saveReading("amatitlan", 28, 55, "rainy")


	// Establecer expiraciones correctas por clave (estado actual)
	rdb.Expire(ctx, "municipality:guatemala", 1*time.Hour)
	rdb.Expire(ctx, "municipality:mixco", 1*time.Hour)
	rdb.Expire(ctx, "municipality:amatitlan", 1*time.Hour)

	// Obtener datos completos de Guatemala (estado actual)
	guatemalaWeather, err := rdb.HGetAll(ctx, "municipality:guatemala").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Clima en Guatemala (actual):", guatemalaWeather)
	
	// Obtener solo la temperatura de Mixco
	mixcoTemp, err := rdb.HGet(ctx, "municipality:mixco", "temperature").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temperatura en Mixco:", mixcoTemp, "°C")
	
	// Obtener múltiples campos de Amatitlán
	amatitlanData, err := rdb.HMGet(ctx, "municipality:amatitlan", "temperature", "humidity", "weather").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Datos de Amatitlán - Temp:", amatitlanData[0], "°C, Humedad:", amatitlanData[1], "%, Clima:", amatitlanData[2])
}
