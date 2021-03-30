package models

import "gopkg.in/mgo.v2/bson"

type Crypto struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	IdStr    string        `bson:"_idstr,omitempty" json:"idstr"`
	Name     string        `bson:"name" json:"name"`
	Likes    int           `bson:"likes" json:"likes"`
	Dislikes int           `bson:"dislikes" json:"dislikes"`
}

func NewCrypto() *Crypto {
	return &Crypto{
		IdStr: bson.NewObjectId().Hex(),
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
