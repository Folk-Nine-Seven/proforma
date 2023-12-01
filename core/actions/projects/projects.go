package project

import (
	"folk/proforma/core/model"

	"github.com/google/uuid"
)

type (
	NewProjectInput struct {
		Name      string           `json:"name" binding:"required"`
		Tasks     []model.Task     `json:"tasks" binding:"dive"`
		Locations []model.Location `json:"locations" binding:"dive"`
	}
)

func New(input NewProjectInput) *model.Project {
	return &model.Project{
		Id:        uuid.New().String(),
		Name:      input.Name,
		Tasks:     input.Tasks,
		Locations: input.Locations,
	}
}

func Get(projectId string) model.Project {
	return model.Project{}
}

func AddLocation(proj *model.Project, location *model.Location) {
	if location == nil {
		return
	}

	proj.Locations = append(proj.Locations, *location)
}
