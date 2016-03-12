package repository

import "gopkg.in/mgo.v2"

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
	return nil, nil
}

func (me *GenericRepository) Find(collection string, id string) (object interface{}, err error) {
	return nil, nil
}

func (me *GenericRepository) Delete(collection string, id string) (err error) {
	return nil
}
