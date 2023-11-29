package database

import "folk/proforma/core/model"

type OrganizationStore interface {
	Create(name, description string) (model.Organization, error)
	Delete(id string) error
	Update(id string, changes model.Organization) (model.Organization, error)
	Describe(id string) (model.Organization, error)
	List() ([]model.Organization, error)
}