package stores

import "folk/proforma/core/model"

type ProjectStore interface {
	CreateProject(orgId string, newProj model.Project) (*model.Project, error)
	DescribeProject(orgId, projId string) (model.Project, error)
	UpdateProject(orgId, projId string, proj model.Project) (model.Project, error)
	DeleteProject(orgId, projId string) error
	ListProjects(orgId string) ([]model.Project, error)
}
