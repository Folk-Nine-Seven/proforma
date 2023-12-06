package main

import (
	"fmt"
	"folk/proforma/gateways/neo4j"
	"folk/proforma/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := neo4j.Instance()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	defer db.Close()

	os.Setenv("PORT", "8080")

	r := gin.Default()
	public := r.Group("/api")

	public.GET("/", version)

	initializePaths(public, routes.Organizations)
	initializePaths(public, routes.Projects)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initializePaths(group *gin.RouterGroup, routes map[string]map[string]func(c *gin.Context)) {
	for path, content := range routes {
		for verb, function := range content {
			switch verb {
			case "DELETE":
				group.DELETE(path, function)
			case "GET":
				group.GET(path, function)
			case "PATCH":
				group.PATCH(path, function)
			case "POST":
				group.POST(path, function)
			case "PUT":
				group.PUT(path, function)
			}
		}
	}
}

func version(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, `v0.1`)
}
