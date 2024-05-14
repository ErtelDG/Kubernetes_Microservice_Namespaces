// main.go
package main

import (
	"log"
	"net"

	"github.com/erteldg/grpcnamespaceservice/pkg/config"
	"github.com/erteldg/grpcnamespaceservice/pkg/model"
	pb "github.com/erteldg/grpcnamespaceservice/pkg/proto"
	"google.golang.org/grpc"
)

func main() {

	clientset, err := config.GetClientset()
	if err != nil {
		log.Fatalf("Failed to create clientset: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterNamespaceServiceServer(s, &model.Server{Clientset: clientset})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
