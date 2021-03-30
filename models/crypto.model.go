package models

import "gopkg.in/mgo.v2/bson"

type Crypto struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Likes    int           `bson:"likes" json:"likes"`
	Dislikes int           `bson:"dislikes" json:"dislikes"`
}
