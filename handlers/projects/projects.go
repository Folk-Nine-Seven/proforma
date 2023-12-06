package projects

import (
	"folk/proforma/core/actions/projects"
	actions "folk/proforma/core/actions/projects"
	"folk/proforma/gateways/neo4j"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	// newProject := project.New(project.NewProjectInput{})

	// if err := c.BindJSON(&newProject); err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, err)
	// 	return
	// }
	// id := c.Param("id")
	// for ndx, organization := range orgs {
	// 	if id == organization.Id {
	// 		orgs[ndx].Projects[newProject.Id] = *newProject
	// 		c.IndentedJSON(http.StatusCreated, newProject)
	// 		return
	// 	}

	// c.IndentedJSON(http.StatusNotFound, "organization not found")
}

func Describe(c *gin.Context) {
	// newProject := project.New(project.NewProjectInput{})

	// if err := c.BindJSON(&newProject); err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, err)
	// 	return
	// }
	// id := c.Param("id")
	// for ndx, organization := range orgs {
	// 	if id == organization.Id {
	// 		orgs[ndx].Projects[newProject.Id] = *newProject
	// 		c.IndentedJSON(http.StatusCreated, newProject)
	// 		return
	// 	}

	// c.IndentedJSON(http.StatusNotFound, "organization not found")
}

func Create(c *gin.Context) {
	db, err := neo4j.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	newProject := projects.New(projects.NewProjectInput{})

	if err := c.BindJSON(&newProject); err != nil {
		return
	}
	orgId := c.Param("orgId")

	proj, err := actions.Create(orgId, *newProject, db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, proj)
}

func Delete(c *gin.Context) {

}
