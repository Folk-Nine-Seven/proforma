package organizations

import (
	"fmt"
	"folk/proforma/core/actions/organization"
	"folk/proforma/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Describe(c *gin.Context) {
	db, err := database.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	id := c.Param("orgId")
	result, err := neo4j.ExecuteQuery(c, db.Driver,
		fmt.Sprintf("MATCH (o:Organization WHERE o.id = '%s') RETURN o", id),
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result.Records)
}

func List(c *gin.Context) {
	db, err := database.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	result, err := neo4j.ExecuteQuery(c, db.Driver,
		"MATCH (o:Organization) RETURN o",
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result.Records)
}

func Create(c *gin.Context) {
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
