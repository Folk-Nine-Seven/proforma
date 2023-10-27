package location

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	Address struct {
		Number string `json:"number" binding:"required"`
		Street string `json:"street" binding:"required"`
		Unit   string `json:"unit"`
		City   string `json:"city" binding:"required"`
		State  string `json:"state" binding:"required"`
		Zip    string `json:"zip" binding:"required"`
	}

	Location struct {
		Id      string
		Name    string  `json:"name" binding:"required"`
		Address Address `json:"address" binding:"required"`
	}

	Site struct {
		Location
	}
)

func New(name string, address Address) *Location {
	return &Location{
		Id:      uuid.New().String(),
		Name:    name,
		Address: address,
	}
}

func (a *Address) ToString() string {
	return fmt.Sprintf("%s %s, %s\n%s, %s %s", a.Number, a.Street, a.Unit, a.City, a.State, a.Zip)
}

func (l *Location) ToString() string {
	return fmt.Sprintf("%s\n%s", l.Name, l.Address.ToString())
}
