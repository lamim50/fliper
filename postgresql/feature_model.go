package postgresql

import (
	"fliper/entities"
)

type FeatureModel struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Enabled     bool
}

func (r *FeatureModel) ToEntity() *entities.Feature {
	return &entities.Feature{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Enabled:     r.Enabled,
	}
}
