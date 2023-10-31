package main

import (
	"folk/proforma/core/actions/organization"
	"folk/proforma/core/actions/project"
	"folk/proforma/core/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var projects []model.Project
var organizations []model.Organization

func main() {
	projects = make([]model.Project, 0)
	organizations = make([]model.Organization, 0)
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

	public.GET("/organizations/:id/projects", getProjects)

	public.POST("organizations/:id/projects", createProject)

	public.GET("/organizations", getOrganizations)

	public.POST("/organizations", createOrganizations)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getProjects(c *gin.Context) {
	id := c.Param("id")
	for _, organization := range organizations {
		if id == organization.Id {
			projs := make([]model.Project, 0)
			for _, p := range organization.Projects {
				projs = append(projs, p)
			}
			c.JSON(http.StatusOK, projs)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, id)
}

func createProject(c *gin.Context) {
	newProject := project.New(project.NewProjectInput{})

	if err := c.BindJSON(&newProject); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	id := c.Param("id")
	for ndx, organization := range organizations {
		if id == organization.Id {
			organizations[ndx].Projects[newProject.Id] = *newProject
			c.IndentedJSON(http.StatusCreated, newProject)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, "organization not found")
}

func getOrganizations(c *gin.Context) {
	c.JSON(http.StatusOK, organizations)
}

func createOrganizations(c *gin.Context) {
	newOrganization := organization.New(organization.NewOrganizationInput{})

	if err := c.BindJSON(&newOrganization); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	organizations = append(organizations, *newOrganization)
	c.IndentedJSON(http.StatusCreated, newOrganization)
}
