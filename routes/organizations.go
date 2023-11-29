package routes

import (
	"folk/proforma/handlers/organizations"

	"github.com/gin-gonic/gin"
)

var Organizations map[string]map[string]func(c *gin.Context) = map[string]map[string]func(c *gin.Context){
	"/organizations": {
		"GET":  organizations.List,
		"POST": organizations.Create,
	},
	"/organizations/:orgId": {
		"GET":    organizations.Describe,
		"DELETE": nil,
		"PUT":    nil,
	},
}
