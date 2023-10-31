package organization

import (
	"folk/proforma/core/model"

	"github.com/google/uuid"
)

type (
	NewOrganizationInput struct {
		Name        string
		Description string
		Projects    []model.Project
	}
)

func New(input NewOrganizationInput) *model.Organization {
	projects := make(map[string]model.Project)
	for _, project := range input.Projects {
		projects[project.Id] = project
	}
	return &model.Organization{
		Id:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		Projects:    projects,
	}
}

func Get(organizationId string) model.Organization {
	return model.Organization{}
}
