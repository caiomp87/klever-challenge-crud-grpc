package repositories

import (
	"backend/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CryptoDAO struct {
	Db         *mgo.Database
	Collection string
}

func (dao CryptoDAO) FindAll() (cryptos []models.Crypto, err error) {
	err = dao.Db.C(dao.Collection).Find(bson.M{}).All(&cryptos)

	return
}

func (dao CryptoDAO) FindById(id string) (crypto models.Crypto, err error) {
	err = dao.Db.C(dao.Collection).FindId(bson.ObjectIdHex(id)).One(&crypto)

	return
}

func (dao CryptoDAO) Create(crypto *models.Crypto) (err error) {
	err = dao.Db.C(dao.Collection).Insert(&crypto)

	return err
}

func (dao CryptoDAO) Update(crypto *models.Crypto) error {
	err := dao.Db.C(dao.Collection).UpdateId(crypto.Id, &crypto)

	return err
}

func (dao CryptoDAO) Delete(id string) error {
	err := dao.Db.C(dao.Collection).RemoveId(bson.ObjectIdHex(id))

	return err
}

func (dao CryptoDAO) AddLike(crypto *models.Crypto) error {
	err := dao.Db.C(dao.Collection).UpdateId(crypto.Id, &crypto)

	return err
}
