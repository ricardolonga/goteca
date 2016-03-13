package repository

import (
	"gopkg.in/mgo.v2"
	"camlistore.org/third_party/labix.org/v2/mgo/bson"
	"gitlab.com/ricardolonga/goteca/entity"
)

const DB = "goteca"

type GenericRepository struct {
	MongoSession *mgo.Session
}

func (me *GenericRepository) Save(collection string, object interface{}) (savedObject interface{}, err error) {
	me.MongoSession.Refresh()

	object.(*entity.Movie).Id = bson.NewObjectId().Hex()

	if err := me.MongoSession.DB(DB).C(collection).Insert(object); err != nil {
		return nil, err
	}

	return object, nil
}

func (me *GenericRepository) FindAll(collection string) (objects []interface{}, err error) {
	me.MongoSession.Refresh()

	movies := make([]interface{}, 0)

	if err := me.MongoSession.DB(DB).C(collection).Find(nil).All(&movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (me *GenericRepository) Find(collection string, id string) (object interface{}, err error) {
	me.MongoSession.Refresh()

	if err := me.MongoSession.DB(DB).C(collection).FindId(id).One(&object); err != nil {
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
