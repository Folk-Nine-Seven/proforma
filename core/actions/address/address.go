package address

import (
	"fmt"
	"folk/proforma/core/model"
)

func String(address *model.Address) string {
	return fmt.Sprintf("%s %s, %s\n%s, %s %s", address.Number, address.Street, address.Unit, address.City, address.State, address.Zip)
}
