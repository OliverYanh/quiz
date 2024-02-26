package main

import (
	"fmt"
	"net"

	pb "github.com/oliveryanh/quiz/internal/quiz/generated"
	server "github.com/oliveryanh/quiz/internal/quiz/grpc"

	"google.golang.org/grpc"
)

const (
	port = ":50051" // Port on which the gRPC server listens
)

func main() {
	// Create a listener on the specified port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}

	// Create a new instance of the QuizServer
	svc := server.NewQuizServer()

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the QuizServiceServer with the gRPC server
	pb.RegisterQuizServiceServer(s, svc)

	// Start the gRPC server
	fmt.Println("Starting gRPC quiz server on", port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}
}
