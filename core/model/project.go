package model

type (
	Project struct {
		Id        string     `json:"id"`
		Name      string     `json:"name" binding:"required"`
		Locations []Location `json:"locations,omitempty" binding:"dive"`
		Tasks     []Task     `json:"tasks,omitempty" binding:"dive"`
		Metadata
	}
)
