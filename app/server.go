package app

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGrpcServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Could not connect", err)
	}

	grpcServer := grpc.NewServer()
	log.Println("gRPC server has been started on port :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Could not connect", err)
	}
}
