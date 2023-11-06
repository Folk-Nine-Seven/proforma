package main

import (
	"fmt"
	"folk/proforma/core/actions/organization"
	"folk/proforma/core/actions/project"
	"folk/proforma/core/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// var projects []model.Project
var orgs []model.Organization

func main() {
	//projects = make([]model.Project, 0)
	orgs = make([]model.Organization, 0)
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
	for _, organization := range orgs {
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
	for ndx, organization := range orgs {
		if id == organization.Id {
			orgs[ndx].Projects[newProject.Id] = *newProject
			c.IndentedJSON(http.StatusCreated, newProject)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, "organization not found")
}

func getOrganizations(c *gin.Context) {
	uri := "neo4j+s://1adaebe1.databases.neo4j.io:7687"
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth("neo4j", "Zutnassi1qY7Ip1GTENC5l7TXTgbNYHJuayZydriLUU", ""))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	defer driver.Close(c)

	session := driver.NewSession(c, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(c)

	// newOrganization := organization.New(organization.NewOrganizationInput{})

	// if err := c.BindJSON(&newOrganization); err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, err)
	// 	return
	// }
	organizations, err := session.ExecuteWrite(c, func(transaction neo4j.ManagedTransaction) (any, error) {
		cmd := `MATCH (o:Organization)
		RETURN o`
		result, err := transaction.Run(c, cmd, nil)
		if err != nil {
			return nil, err
		}

		if result.Next(c) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, organizations)
}

func createOrganizations(c *gin.Context) {
	uri := "neo4j+s://1adaebe1.databases.neo4j.io:7687"
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth("neo4j", "Zutnassi1qY7Ip1GTENC5l7TXTgbNYHJuayZydriLUU", ""))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	defer driver.Close(c)

	session := driver.NewSession(c, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(c)

	newOrganization := organization.New(organization.NewOrganizationInput{})

	if err := c.BindJSON(&newOrganization); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	greeting, err := session.ExecuteWrite(c, func(transaction neo4j.ManagedTransaction) (any, error) {
		cmd := fmt.Sprintf(`CREATE (a:Organization) SET a.id = "%s" SET a.name = "%s" SET a.description = "%s"`, newOrganization.Id, newOrganization.Name, newOrganization.Description)
		result, err := transaction.Run(c, cmd, nil)
		if err != nil {
			return nil, err
		}

		if result.Next(c) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	orgs = append(orgs, *newOrganization)
	c.IndentedJSON(http.StatusCreated, greeting)
}
