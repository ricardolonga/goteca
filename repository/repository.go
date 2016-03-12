package repository

type Repository interface {
	Save(collection string, object interface{}) (savedObject interface{}, err error)
	FindAll(collection string) (objects []interface{}, err error)
	Find(collection string, id string) (object interface{}, err error)
	Delete(collection string, id string) (err error)
}

