package organizations

import (
	"folk/proforma/core/interfaces/database"
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

var ds database.OrganizationStore

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

func GetOrganization(id string) (*model.Organization, error) {
	return ds.Describe(id)
}

func GetOrganizations() ([]model.Organization, error) {
	return ds.List()
}

func CreateOrganization(name, description string) (*model.Organization, error) {
	return ds.Create(name, description)
}

func SetDataStore(dataStore database.OrganizationStore) {
	ds = dataStore
}
