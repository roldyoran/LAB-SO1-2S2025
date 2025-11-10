package main

import (
	"context"
	"fmt"
	"log"
	time "time"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "clima",
		GroupID: "clima-consumer-group",
	})

	fmt.Println("📥 Esperando mensajes de Kafka...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("❌ Error leyendo mensaje:", err)
			continue
		}
		fmt.Printf("✅ [Kafka] Mensaje recibido: %s\n", string(m.Value))
		time.Sleep(2 * time.Second) // Simular procesamiento
	}
}
