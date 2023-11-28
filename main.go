package main

import (
	"fmt"
	"folk/proforma/core/actions/organization"
	"folk/proforma/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var db *database.Neo4j

func main() {
	db, err := database.Instance()
	if err != nil {
		log.Fatal("no database connection")
	}
	defer db.Close()

	r := gin.Default()
	public := r.Group("/api")

	public.GET("/", version)

	public.GET("/organizations/:id", getOrganization)

	public.GET("/organizations/:id/projects", getProjects)

	public.POST("organizations/:id/projects", createProject)

	public.GET("/organizations", getOrganizations)

	public.POST("/organizations", createOrganizations)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func version(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, `v0.1`)
}

func getProjects(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusNotFound, id)
}

func getOrganization(c *gin.Context) {
	id := c.Param("id")
	result, err := neo4j.ExecuteQuery(c, db.Driver,
		fmt.Sprintf("MATCH (o:Organization WHERE o.id = '%s') RETURN o", id),
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result.Records)
}

func createProject(c *gin.Context) {
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

func getOrganizations(c *gin.Context) {
	result, err := neo4j.ExecuteQuery(c, db.Driver,
		"MATCH (o:Organization) RETURN o",
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result.Records)
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

	//orgs = append(orgs, *newOrganization)
	c.IndentedJSON(http.StatusCreated, greeting)
}
