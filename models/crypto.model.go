package models

import "gopkg.in/mgo.v2/bson"

type Crypto struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name" json:"name"`
	Likes    int    `bson:"likes" json:"likes"`
	Dislikes int    `bson:"dislikes" json:"dislikes"`
}

func NewCrypto() *Crypto {
	return &Crypto{
		Id: bson.NewObjectId().Hex(),
	}
}

type Cryptos struct {
	Crypto []*Crypto `bson:"cryptos" json:"cryptos"`
}

func (c *Cryptos) Add(crypto *Crypto) {
	c.Crypto = append(c.Crypto, crypto)
}

func NewCryptos() *Cryptos {
	return &Cryptos{}
}
