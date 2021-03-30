package services

import (
	"backend/app/pb"
	"backend/models"
	"context"
	"time"
)

type CryptoGrpcServer struct {
	pb.UnimplementedCryptoServiceServer
	Cryptos *models.Cryptos
}

func (c *CryptoGrpcServer) CreateCrypto(ctx context.Context, in *pb.Crypto) (*pb.CryptoResult, error) {
	crypto := models.NewCrypto()
	crypto.Name = in.Name
	c.Cryptos.Add(crypto)

	return &pb.CryptoResult{
		Id:   string(crypto.Id),
		Name: crypto.Name,
	}, nil
}

func (c *CryptoGrpcServer) ListCryptos(req *pb.Empty, stream pb.CryptoService_ListCryptosServer) error {
	for _, crypto := range c.Cryptos.Crypto {
		time.Sleep(time.Second * 5)
		stream.Send(&pb.CryptoResult{Name: crypto.Name, Id: crypto.Id})
	}

	return nil
}

func NewCryptoGrpcServer(cryptos *models.Cryptos) *CryptoGrpcServer {
	return &CryptoGrpcServer{
		Cryptos: cryptos,
	}
}
