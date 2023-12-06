package stores

import "folk/proforma/core/model"

type OrganizationStore interface {
	CreateOrganization(newOrg model.Organization) (*model.Organization, error)
	DeleteOrganization(id string) error
	UpdateOrganization(id string, changes *model.Organization) (*model.Organization, error)
	DescribeOrganization(id string) (*model.Organization, error)
	ListOrganizations() ([]model.Organization, error)
}
