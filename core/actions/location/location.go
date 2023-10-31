package location

import (
	"fmt"
	"folk/proforma/core/actions/address"
	"folk/proforma/core/model"

	"github.com/google/uuid"
)

func New(name string, address model.Address) *model.Location {
	return &model.Location{
		Id:      uuid.New().String(),
		Name:    name,
		Address: address,
	}
}

func ToString(location *model.Location) string {
	return fmt.Sprintf("%s\n%s", location.Name, address.String(&location.Address))
}
