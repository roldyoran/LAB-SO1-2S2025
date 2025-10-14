package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("❌ Error conectando a RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("❌ Error abriendo canal:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("clima", false, false, false, false, nil)
	if err != nil {
		log.Fatal("❌ Error declarando cola:", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("❌ Error al consumir mensajes:", err)
	}

	fmt.Println("📥 Esperando mensajes de RabbitMQ...")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("✅ [RabbitMQ] Mensaje recibido: %s\n", d.Body)
			time.Sleep(2 * time.Second) // Simular procesamiento
		}
	}()

	<-forever
}
