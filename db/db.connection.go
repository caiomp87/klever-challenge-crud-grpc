package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

func Connect() (*mgo.Database, error) {
	connectionString := fmt.Sprintf("mongo://%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	db := session.DB(os.Getenv("DB_NAME"))
	return db, nil
}
