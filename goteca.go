package goteca

type Movie struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Dao interface {
	Find(id string) (*Movie, error)
	FindAll() ([]*Movie, error)
	Save(object *Movie) (*Movie, error)
	Delete(id string) error
}

type Service interface {
	Get(ID string) (*Movie, error)
	GetAll() ([]*Movie, error)
	Save(movie *Movie) (*Movie, error)
	Delete(ID string) error
}

type serviceImpl struct {
	dao Dao
}

func NewService(dao Dao) Service {
	return &serviceImpl{
		dao: dao,
	}
}

func (me *serviceImpl) Get(ID string) (*Movie, error) {
	return me.dao.Find(ID)
}

func (me *serviceImpl) GetAll() ([]*Movie, error) {
	return me.dao.FindAll()
}

func (me *serviceImpl) Save(movie *Movie) (*Movie, error) {
	return me.dao.Save(movie)
}

func (me *serviceImpl) Delete(ID string) error {
	return me.dao.Delete(ID)
}
