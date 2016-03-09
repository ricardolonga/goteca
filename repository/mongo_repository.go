package repository

import "gopkg.in/mgo.v2"

type MongoRepository struct {
	MongoSession *mgo.Session
}

func (me *MongoRepository) Save() {}

func (me *MongoRepository) Find() {}

func (me *MongoRepository) Delete() {}
