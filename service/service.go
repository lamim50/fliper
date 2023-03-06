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

func (s Service) Create(title string, description string) (*entities.Feature, error) {
	feature := entities.NewFeature(title, description)
	_, err := s.Repository.Create(feature)

	if err != nil {
		return nil, err
	}

	return feature, nil
}

func (s Service) Update(id string, title string, description string) (*entities.Feature, error) {
	feature, err := s.Repository.FindByID(id)
	log.Println(feature)
	feature.Update(title, description)
	log.Println(feature)

	_, err = s.Repository.Update(feature)
	log.Println(feature)

	if err != nil {
		return nil, err
	}
	return feature, nil
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

func (s Service) DeleteByID(id string) error {
	err := s.Repository.DeleteByID(id)
	return err
}
