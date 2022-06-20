package xdatabase

type Entity[ID comparable] interface {
	EntityID() ID
}

type Repository[T Entity[ID], ID comparable] interface {
	Save(T) error
	SaveAll([]T) error

	Update(T) error
	UpdateAll([]T) error

	Delete(T) error
	DeleteByIdInBatch([]ID) error

	Find(T) ([]T, error)
	FindByIdInBatch([]ID) ([]T, error)
}
