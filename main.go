// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/trigo-William-HOANG/project-root/invoicer"
)

// server is used to implement the gRPC service.
type server struct {
	pb.UnimplementedInvoicerServer
}

// InvoiceResponseJSON is a custom struct to be returned as JSON.
type InvoiceResponseJSON struct {
	Id       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Paid     bool    `json:"paid"`
}

// GetInvoice method will return the invoice in JSON format.
func (s *server) GetInvoice(ctx context.Context, in *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	// Example data
	invoiceData := InvoiceResponseJSON{
		Id:       "12345",
		Amount:   500.0,
		Currency: "USD",
		Paid:     true,
	}

	// Convert the Go struct to JSON
	jsonData, err := json.Marshal(invoiceData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Return the JSON string inside the response
	return &pb.InvoiceResponse{
		InvoiceJson: string(jsonData),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterInvoicerServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
