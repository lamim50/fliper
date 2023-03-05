package entities

import "github.com/google/uuid"

type Feature struct {
	id          string
	title       string
	description string
	enabled     bool
}

func NewFeature(title string, description string) *Feature {
	id := uuid.NewString()
	return &Feature{id: id, title: title, description: description, enabled: false}
}

func (f Feature) Id() string {
	return f.id
}

func (f Feature) Title() string {
	return f.title
}

func (f Feature) Description() string {
	return f.description
}

func (f Feature) Enabled() bool {
	return f.enabled
}
