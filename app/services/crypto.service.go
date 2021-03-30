package services

import (
	"backend/app/pb"
	"backend/models"
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CryptoGrpcServer struct {
	pb.UnimplementedCryptoServiceServer
	Cryptos *models.Cryptos
}

func (c *CryptoGrpcServer) CreateCrypto(ctx context.Context, in *pb.Crypto) (*pb.CryptoResult, error) {
	// var database *mgo.Database
	// var err error
	var crypto models.Crypto

	crypto.IdStr = bson.NewObjectId().Hex()
	crypto.Id = bson.ObjectIdHex(crypto.IdStr)
	crypto.Name = in.Name
	crypto.Likes = 0
	crypto.Dislikes = 0

	// database, err = db.Connect()
	// if err != nil {
	// 	return nil, err
	// }

	// defer database.Session.Close()

	// cryptoRepository := repositories.CryptoDAO{
	// 	Db:         database,
	// 	Collection: "cryptos",
	// }

	// err = cryptoRepository.Create(&crypto)
	// if err != nil {
	// 	return nil, err
	// }

	return &pb.CryptoResult{
		Id:   crypto.IdStr,
		Name: crypto.Name,
	}, nil
}

func (c *CryptoGrpcServer) ListCryptos(req *pb.Empty, stream pb.CryptoService_ListCryptosServer) error {
	for _, crypto := range c.Cryptos.Crypto {
		time.Sleep(time.Second * 5)
		stream.Send(&pb.CryptoResult{Name: crypto.Name, Id: crypto.IdStr})
	}

	return nil
}

func NewCryptoGrpcServer(cryptos *models.Cryptos) *CryptoGrpcServer {
	return &CryptoGrpcServer{
		Cryptos: cryptos,
	}
}
