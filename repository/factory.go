package repository

func New() Repository {
	return &MongoRepository{}
}
