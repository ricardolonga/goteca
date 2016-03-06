package repository

type Repository interface {
	Save()
	Find()
	Delete()
}

