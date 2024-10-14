package main

import (
	"context"
	"log"
	"net"

	"github.com/maximilienandile/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s MyInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("Document invoice"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &MyInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("cannot serve: %s", err)
	}
}
