package main

import (
	"context"
	"log"
	"time"

	pb "grpc/proto" // Importa el paquete generado para el servicio gRPC

	"google.golang.org/grpc" // Importa la librería gRPC de Google
)

func main() {
	// Conectar con el servidor gRPC en localhost en el puerto 50051
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err) // Manejo de error si la conexión falla
	}
	defer conn.Close() // Asegura que la conexión se cierre al finalizar

	// Crear un cliente para el servicio TweetService definido en el proto
	client := pb.NewTweetServiceClient(conn)

	// Crear un contexto con un tiempo límite de 1 segundo para la solicitud
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // Asegura que el contexto se cancele al finalizar

	// Enviar una solicitud al servidor con los datos del tweet
	resp, err := client.SendTweet(ctx, &pb.TweetRequest{
		Description: "Hola desde Go con gRPC!", // Descripción del tweet
		Country:     "Guatemala",               // País desde donde se envía
		Weather:     "Soleado",                 // Clima asociado al tweet
	})
	if err != nil {
		log.Fatalf("Error al enviar tweet: %v", err) // Manejo de error si la solicitud falla
	}

	// Imprimir la respuesta del servidor
	log.Printf("Respuesta del servidor: %s", resp.Status)
}
