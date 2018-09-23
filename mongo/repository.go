package mongo

import (
	"github.com/ricardolonga/goteca"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoImpl struct {
	session *mgo.Session
}

func NewDao(mongoSession *mgo.Session) goteca.Dao {
	return &mongoImpl{session: mongoSession}
}

func (me *mongoImpl) Save(object *goteca.Movie) (*goteca.Movie, error) {
	session := me.session.Copy()

	object.Id = bson.NewObjectId().Hex()

	if err := session.DB("goteca").C("movies").Insert(object); err != nil {
		return nil, err
	}

	return object, nil
}

func (me *mongoImpl) FindAll() ([]*goteca.Movie, error) {
	session := me.session.Copy()

	movies := make([]*goteca.Movie, 0)

	if err := session.DB("goteca").C("movies").Find(nil).All(&movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (me *mongoImpl) Find(id string) (*goteca.Movie, error) {
	session := me.session.Copy()

	movie := &goteca.Movie{}

	if err := session.DB("goteca").C("movies").FindId(id).One(movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (me *mongoImpl) Delete(ID string) error {
	session := me.session.Copy()

	if err := session.DB("goteca").C("movies").RemoveId(ID); err != nil {
		return err
	}

	return nil
}
