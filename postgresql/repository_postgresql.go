package postgresql

import (
	"fliper/entities"
	"fmt"
	"log"
)

type RepositoryPostgreSQL struct{}

func (s RepositoryPostgreSQL) Create(feature entities.Feature) (entities.Feature, error) {
	return entities.Feature{}, fmt.Errorf("TOTO")
}

func (s RepositoryPostgreSQL) Update(feature entities.Feature) (entities.Feature, error) {
	log.Println("")
	return entities.Feature{}, nil
}

func (s RepositoryPostgreSQL) FindByID(id string) (*entities.Feature, error) {
	feat := entities.NewFeature(id, "licz zieugc lxizeg")
	return feat, nil
}

func (s RepositoryPostgreSQL) FindAll() ([]*entities.Feature, error) {
	features := []*entities.Feature{
		entities.NewFeature("Feat 1", "licz zieugc lxizeg"),
		entities.NewFeature("Feat 2", "dzi eziuf iohbczouie"),
	}
	return features, nil
}
