package postgresql

import (
	"fliper/entities"
	"gorm.io/gorm"
)

type RepositoryPostgreSQL struct {
	db gorm.DB
}

func NewRepositoryPostgreSQL(db gorm.DB) *RepositoryPostgreSQL {
	return &RepositoryPostgreSQL{db: db}
}

func (s RepositoryPostgreSQL) Create(feature *entities.Feature) (*entities.Feature, error) {
	model := &FeatureModel{
		ID:          feature.ID,
		Title:       feature.Title,
		Description: feature.Description,
		Enabled:     feature.Enabled,
	}
	res := s.db.Create(model)
	if res.Error != nil {
		return nil, res.Error
	}

	return model.ToEntity(), nil
}

func (s RepositoryPostgreSQL) Update(feature *entities.Feature) (*entities.Feature, error) {
	model := &FeatureModel{
		ID:          feature.ID,
		Title:       feature.Title,
		Description: feature.Description,
		Enabled:     feature.Enabled,
	}
	res := s.db.Updates(model)
	if res.Error != nil {
		return nil, res.Error
	}

	return model.ToEntity(), nil
}

func (s RepositoryPostgreSQL) FindByID(id string) (*entities.Feature, error) {
	model := &FeatureModel{ID: id}
	res := s.db.First(model)
	if res.Error != nil {
		return nil, res.Error
	}

	return model.ToEntity(), nil
}

func (s RepositoryPostgreSQL) FindAll() ([]*entities.Feature, error) {
	model := new([]FeatureModel)
	res := s.db.Find(model)
	if res.Error != nil {
		return nil, res.Error
	}

	var features []*entities.Feature
	for _, val := range *model {
		features = append(features, val.ToEntity())
	}

	return features, nil
}

func (s RepositoryPostgreSQL) DeleteByID(id string) error {
	model := &FeatureModel{ID: id}
	res := s.db.Delete(model)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
