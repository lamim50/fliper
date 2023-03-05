package service

import (
	"fliper/entities"
	"log"
)

type Service struct {
	Repository entities.Repository
}

func NewService(repository entities.Repository) *Service {
	return &Service{Repository: repository}
}

func (s Service) Create(title string, description string) (entities.Feature, error) {
	feature := entities.NewFeature(title, description)
	res, err := s.Repository.Create(*feature)

	if err != nil {
		log.Println(err)
		return entities.Feature{}, err
	}

	return res, nil

}

func (s Service) FindByID(id string) (*entities.Feature, error) {
	features, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (s Service) FindAll() ([]*entities.Feature, error) {
	features, err := s.Repository.FindAll()
	if err != nil {
		return nil, err
	}
	return features, nil
}
