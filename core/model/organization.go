package model

type (
	Organization struct {
		Id          string             `json:"id"`
		Name        string             `json:"name" binding:"required"`
		Description string             `json:"description"`
		Projects    map[string]Project `json:"projects" binding:"dive"`
	}
)
