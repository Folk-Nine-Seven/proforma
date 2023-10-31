package model

type (
	Organization struct {
		Id       string             `json:"id"`
		Name     string             `json:"name" binding:"required"`
		Projects map[string]Project `json:"projects" binding:"dive"`
	}
)
