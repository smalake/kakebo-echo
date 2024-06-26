package repository

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id int) (T, error)
	Create(entity T) (T, error)
	Update(entity T) (T, error)
	Delete(id int) error
}
