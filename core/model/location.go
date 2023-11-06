package model

type (
	Location struct {
		Metadata
		Id      string
		Name    string  `json:"name" binding:"required"`
		Address Address `json:"address" binding:"required"`
	}

	Site struct {
		Location
	}
)
