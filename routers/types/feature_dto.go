package types

import "fliper/entities"

type FeatureDTO struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

func FromFeature(feature *entities.Feature) *FeatureDTO {
	return &FeatureDTO{
		Id:          feature.Id(),
		Title:       feature.Title(),
		Description: feature.Description(),
		Enabled:     feature.Enabled(),
	}
}

func FromFeatures(features []*entities.Feature) []*FeatureDTO {
	var res []*FeatureDTO
	for _, value := range features {
		res = append(res, FromFeature(value))
	}

	return res
}
