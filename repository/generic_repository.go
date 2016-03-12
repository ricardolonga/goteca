package repository

import (
	"gopkg.in/mgo.v2"
	"camlistore.org/third_party/labix.org/v2/mgo/bson"
)

const DB = "goteca"

type GenericRepository struct {
	MongoSession *mgo.Session
}

func (me *GenericRepository) Save(collection string, object interface{}) (savedObject interface{}, err error) {
	me.MongoSession.Refresh()

	if err := me.MongoSession.DB(DB).C(collection).Insert(object); err != nil {
		return nil, err
	}

	return object, nil
}

func (me *GenericRepository) FindAll(collection string) (objects []interface{}, err error) {
	me.MongoSession.Refresh()

	if err := me.MongoSession.DB(DB).C(collection).Find(bson.M{}).All(objects); err != nil {
		return nil, err
	}

	return objects, nil
}

func (me *GenericRepository) Find(collection string, id string) (object interface{}, err error) {
	me.MongoSession.Refresh()

	if err := me.MongoSession.DB(DB).C(collection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(object); err != nil {
		return nil, err
	}

	return object, nil
}

func (me *GenericRepository) Delete(collection string, id string) (err error) {
	me.MongoSession.Refresh()

	if err := me.MongoSession.DB(DB).C(collection).RemoveId(id); err != nil {
		return err
	}

	return nil
}
