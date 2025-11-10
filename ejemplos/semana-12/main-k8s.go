package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Clima struct {
	Municipio   string  `json:"municipio"`
	Temperatura int `json:"temperatura"`
	Humedad     int     `json:"humedad"`
	Clima       string  `json:"clima"`
}

var municipios = []string{"Mixco", "Guatemala", "Villa Nueva", "Amatitlán", "Antigua"}
var climas = []string{"Soleado", "Nublado", "Lluvioso", "Ventoso"}

func generarClima() Clima {
	return Clima{
		Municipio:   municipios[rand.Intn(len(municipios))],
		Temperatura: 15 + rand.Intn(15),
		Humedad:     50 + rand.Intn(50),
		Clima:       climas[rand.Intn(len(climas))],
	}
}

// Enviar a Kafka
func enviarKafka(msg Clima, kafkaBroker string) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaBroker, "clima", 0)
	if err != nil {
		return err
	}
	defer conn.Close()

	data, _ := json.Marshal(msg)
	_, err = conn.WriteMessages(kafka.Message{Value: data})
	return err
}

// Enviar a RabbitMQ
func enviarRabbit(msg Clima, rabbitmqURL string) error {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("clima", false, false, false, false, nil)
	if err != nil {
		return err
	}

	data, _ := json.Marshal(msg)
	return ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// Configuración desde variables de entorno
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}
	
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		rabbitmqURL = "amqp://guest:guest@localhost:5672/"
	}
	
	fmt.Printf("🚀 Enviando datos de clima...\n")
	fmt.Printf("📡 Kafka broker: %s\n", kafkaBroker)
	fmt.Printf("🐰 RabbitMQ URL: %s\n", rabbitmqURL)

	for {
		clima := generarClima()
		data, _ := json.MarshalIndent(clima, "", "  ")
		fmt.Println("📦 Nuevo dato:", string(data))

		if rand.Intn(2) == 0 {
			fmt.Println("➡️  Enviando a Kafka")
			if err := enviarKafka(clima, kafkaBroker); err != nil {
				log.Println("❌ Error Kafka:", err)
			} else {
				fmt.Println("✅ Enviado a Kafka exitosamente")
			}
		} else {
			fmt.Println("➡️  Enviando a RabbitMQ")
			if err := enviarRabbit(clima, rabbitmqURL); err != nil {
				log.Println("❌ Error RabbitMQ:", err)
			} else {
				fmt.Println("✅ Enviado a RabbitMQ exitosamente")
			}
		}

		time.Sleep(3 * time.Second)
	}
}

// Comando para consumir desde Kafka:
// kubectl exec -it <kafka-pod-name> -n clima-app -- kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic clima --from-beginning