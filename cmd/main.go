package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	newsv1 "github.com/supLano/go-grpc-proto/api/news/v1"
	internal_grpc "github.com/supLano/go-grpc-proto/internal/grpc"
)	

func main() {
	// Create a listener on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the NewsService implementation
	newsv1.RegisterNewsServiceServer(server, &internal_grpc.Server{})

	healthServer := health.NewServer()
	healthgrpc.RegisterHealthServer(server, healthServer)

	// Start serving
	log.Println("Server started on :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
