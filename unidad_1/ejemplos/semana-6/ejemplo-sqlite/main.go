package main

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Contenedor simulado
type Container struct {
	Name   string
	CPU    float64
	Memory float64
	Status string
}

func main() {
	// Inicializar aleatoriedad
	rand.Seed(time.Now().UnixNano())

	// Conectar a SQLite (crea el archivo si no existe)
	db, err := sql.Open("sqlite3", "./containers.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear tabla si no existe
	createTable := `
	CREATE TABLE IF NOT EXISTS containers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		cpu REAL,
		memory REAL,
		status TEXT,
		created_at INTEGER
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Error creando tabla:", err)
	}

	// Canal para capturar señales (detener daemon)
	// 'sigs' es un canal que se utiliza para recibir señales del sistema operativo.
	// En este caso, se capturan señales SIGINT (Ctrl+C) y SIGTERM (terminación del proceso).
	// Estas señales permiten detener el daemon de manera controlada.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Daemon iniciado. Generando datos cada 44 segundos...")

	// 'ticker' es un temporizador que genera eventos periódicos cada 44 segundos.
	// Estos eventos se utilizan para ejecutar tareas recurrentes, como generar e insertar contenedores aleatorios.
	ticker := time.NewTicker(20 * time.Second)
	// 'defer ticker.Stop()' asegura que el temporizador se detenga correctamente cuando el programa termine.
	defer ticker.Stop()

	// Este bucle infinito permite que el daemon esté en ejecución constante,
	// esperando eventos del ticker o señales del sistema operativo.
loop:
	// El for select es una construcción que permite esperar en múltiples canales.
	for {
		select {
		case <-ticker.C:
			// Este caso se ejecuta cada vez que el ticker genera un evento (cada 44 segundos).
			// Genera contenedores aleatorios y los inserta en la base de datos.
			containers := generateRandomContainers()
			for _, c := range containers {
				_, err := db.Exec("INSERT INTO containers (name, cpu, memory, status, created_at) VALUES (?, ?, ?, ?, ?)",
					c.Name, c.CPU, c.Memory, c.Status, time.Now().UnixMilli())
				if err != nil {
					// Si ocurre un error al insertar, se registra en los logs.
					log.Println("Error insertando:", err)
				} else {
					// Si la inserción es exitosa, se registra el contenedor insertado en los logs.
					log.Printf("Insertado contenedor random: %s (CPU: %.2f, MEM: %.2f, Estado: %s)\n",
						c.Name, c.CPU, c.Memory, c.Status)
				}
			}
		case <-sigs:
			// Este caso se ejecuta cuando se recibe una señal de interrupción o terminación.
			// Detiene el daemon y sale del bucle.
			log.Println("Daemon detenido.")
			break loop
		}
	}
}

// Genera hasta 4 contenedores con datos random
func generateRandomContainers() []Container {
	names := []string{"nginx", "redis", "mysql", "golang-app", "nodejs-app", "python-app", "java-app", "ruby-app", "postgres", "mongodb", "ubuntu", "alpine"}
	statuses := []string{"running", "stopped", "paused", "restarting", "dead", "zombie"}

	n := rand.Intn(4) + 1 // mínimo 1, máximo 4
	var containers []Container
	for i := 0; i < n; i++ {
		containers = append(containers, Container{
			Name:   names[rand.Intn(len(names))],
			CPU:    rand.Float64() * 100, // 0-100%
			Memory: rand.Float64() * 512, // hasta 1GB
			Status: statuses[rand.Intn(len(statuses))],
		})
	}
	return containers
}
