package model

type (
	Organization struct {
		Id          string             `json:"id"`
		Name        string             `json:"name" binding:"required"`
		Description string             `json:"description,omitempty"`
		Projects    map[string]Project `json:"projects,omitempty" binding:"dive"`
		Metadata
	}
)
