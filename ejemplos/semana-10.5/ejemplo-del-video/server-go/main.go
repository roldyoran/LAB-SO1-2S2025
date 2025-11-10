package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server_go/go_grpc_server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedWeatherTweetServiceServer
}

func (s *server) SendTweet(ctx context.Context, req *pb.WeatherTweetRequest) (*pb.WeatherTweetResponse, error) {
	fmt.Printf("Go gRPC recibió: %+v\n", req)
	return &pb.WeatherTweetResponse{Status: "Tweet recibido correctamente en Go ✅"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error escuchando: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWeatherTweetServiceServer(grpcServer, &server{})

	fmt.Println("Go gRPC server escuchando en :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Error en Serve: %v", err)
	}
}