package routes

import (
	"folk/proforma/handlers/projects"

	"github.com/gin-gonic/gin"
)

var Projects map[string]map[string]func(c *gin.Context) = map[string]map[string]func(c *gin.Context){
	"/organizations/:orgId/projects": {
		"GET":  projects.List,
		"POST": projects.Create,
	},
	"/organizations/:orgId/projects/:projId": {
		"GET":    projects.Describe,
		"PATCH":  nil,
		"DELETE": projects.Delete,
	},
}
