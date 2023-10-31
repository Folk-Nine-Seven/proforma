package project

import (
	"folk/proforma/core/model"

	"github.com/google/uuid"
)

func New(name string) *model.Project {
	return &model.Project{
		Id:   uuid.New().String(),
		Name: name,
	}
}

func AddLocation(proj *model.Project, location *model.Location) {
	if location == nil {
		return
	}

	proj.Locations = append(proj.Locations, *location)
}
