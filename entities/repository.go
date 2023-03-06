package entities

type Repository interface {
	Create(feature *Feature) (*Feature, error)
	Update(feature *Feature) (*Feature, error)
	FindByID(id string) (*Feature, error)
	FindAll() ([]*Feature, error)
	DeleteByID(id string) error
}
