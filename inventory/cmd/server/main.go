package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/nexarise/microservices-factory/inventory/internal/service"
	inventoryv1 "github.com/nexarise/microservices-factory/shared/pkg/proto/inventory/v1"
)

const (
	port = ":50051"
)

func main() {
	log.Println("Starting Inventory Service...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	inventoryService := service.NewInventoryService()
	inventoryv1.RegisterInventoryServiceServer(grpcServer, inventoryService)

	reflection.Register(grpcServer)

	log.Printf("Inventory Service listening on %s", port)
	log.Println("Available methods:")
	log.Println("\t - GetPart: получение детали по UUID")
	log.Println("\t - ListParts: получение списка деталей с фильтрацией")
	log.Println("Для тестирования используйте grpcurl или любой gRPC клиент")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
