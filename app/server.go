package app

import (
	"backend/app/pb"
	"backend/app/services"
	"backend/models"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var CryptoList = models.NewCryptos()

func StartGrpcServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Could not connect", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	log.Println("gRPC server has been started on port :50051")

	cryptoService := services.NewCryptoGrpcServer(CryptoList)
	pb.RegisterCryptoServiceServer(grpcServer, cryptoService)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Could not connect", err)
	}
}
