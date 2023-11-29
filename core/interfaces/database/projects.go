package database

import "folk/proforma/core/model"

type ProjectStore interface {
	Create(orgId, name, description string) (model.Project, error)
	Describe(orgId, projId string) (model.Project, error)
	Update(orgId, projId string, proj model.Project) (model.Project, error)
	Delete(orgId, projId string) error
	List(orgId string) ([]model.Project, error)
}
