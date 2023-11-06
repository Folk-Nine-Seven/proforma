package model

type (
	Address struct {
		Number string `json:"number" binding:"required"`
		Street string `json:"street" binding:"required"`
		Unit   string `json:"unit,omitempty"`
		City   string `json:"city" binding:"required"`
		State  string `json:"state" binding:"required"`
		Zip    string `json:"zip" binding:"required"`
	}
)
