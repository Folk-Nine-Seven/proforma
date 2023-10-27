package project

import (
	"folk/proforma/core/entities/location"

	"github.com/google/uuid"
)

type (
	Project struct {
		Id        string              `json:"id"`
		Name      string              `json:"name" binding:"required"`
		Locations []location.Location `json:"locations" binding:"dive"`
	}
)

func New(name string) *Project {
	return &Project{
		Id:   uuid.New().String(),
		Name: name,
	}
}

func (proj *Project) AddLocation(location *location.Location) {
	if location == nil {
		return
	}

	proj.Locations = append(proj.Locations, *location)
}
