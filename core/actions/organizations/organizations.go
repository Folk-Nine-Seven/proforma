package organizations

import (
	"folk/proforma/core/interfaces/stores"
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

func GetOrganization(id string, ds stores.OrganizationStore) (*model.Organization, error) {
	return ds.DescribeOrganization(id)
}

func GetOrganizations(ds stores.OrganizationStore) ([]model.Organization, error) {
	return ds.ListOrganizations()
}

func CreateOrganization(org model.Organization, ds stores.OrganizationStore) (*model.Organization, error) {
	return ds.CreateOrganization(org)
}

func DeleteOrganization(id string, ds stores.OrganizationStore) error {
	return ds.DeleteOrganization(id)
}
