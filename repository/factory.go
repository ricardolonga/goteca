package repository
import "gopkg.in/mgo.v2"

func New(mongoSession *mgo.Session) Repository {
	return &MongoRepository{MongoSession: mongoSession}
}
