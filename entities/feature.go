package entities

import "github.com/google/uuid"

type Feature struct {
	ID          string
	Title       string
	Description string
	Enabled     bool
}

func NewFeature(title string, description string) *Feature {
	id := uuid.NewString()
	return &Feature{ID: id, Title: title, Description: description, Enabled: false}
}

func (f *Feature) Update(title string, description string) {
	f.Title = title
	f.Description = description
}
