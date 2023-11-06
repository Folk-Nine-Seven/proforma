package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password"`
	Metadata
}
