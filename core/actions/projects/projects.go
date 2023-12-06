package projects

import (
	"folk/proforma/core/interfaces/stores"
	"folk/proforma/core/model"

	"github.com/google/uuid"
)

type (
	NewProjectInput struct {
		Description string           `json:"description"`
		Name        string           `json:"name" binding:"required"`
		Tasks       []model.Task     `json:"tasks" binding:"dive"`
		Locations   []model.Location `json:"locations" binding:"dive"`
	}
)

func New(input NewProjectInput) *model.Project {
	return &model.Project{
		Id:          uuid.New().String(),
		Description: input.Description,
		Name:        input.Name,
		Tasks:       input.Tasks,
		Locations:   input.Locations,
	}
}

func Get(projectId string) model.Project {
	return model.Project{}
}

func Create(orgId string, proj model.Project, ds stores.ProjectStore) (*model.Project, error) {
	return ds.CreateProject(orgId, proj)
}

func AddLocation(proj *model.Project, location *model.Location) {
	if location == nil {
		return
	}

	proj.Locations = append(proj.Locations, *location)
}
