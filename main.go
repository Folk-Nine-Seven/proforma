package main

import (
	"folk/proforma/core/entities/project"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var projects []project.Project

func main() {
	projects = make([]project.Project, 0)
	// p := project.New("temp strip club")
	// p.AddLocation(&location.Location{
	// 	Address: location.Address{
	// 		Number: "16557",
	// 		Street: "SW Sidney Lane",
	// 		City:   "Sherwood",
	// 		State:  "OR",
	// 	},
	// })

	// projects = append(projects, *p)

	r := gin.Default()
	public := r.Group("/api")

	public.GET("/projects", getProjects)

	public.POST("/projects", createProject)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getProjects(c *gin.Context) {
	c.JSON(http.StatusOK, projects)
}

func createProject(c *gin.Context) {
	var newProject project.Project

	if err := c.BindJSON(&newProject); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	newProject.Id = uuid.New().String()

	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)

}
