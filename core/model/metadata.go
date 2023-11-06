package model

import "time"

type (
	Metadata struct {
		Created   time.Time `json:"created"`
		Updated   time.Time `json:"updated"`
		CreatedBy string    `json:"ceatedBy"`
		UpdatedBy string    `json:"updatedBy"`
	}
)
