package organizations

import (
	"folk/proforma/core/actions/organizations"
	"folk/proforma/gateways/neo4j"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Describe(c *gin.Context) {
	db, err := neo4j.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	id := c.Param("orgId")
	org, err := organizations.GetOrganization(id, db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if org == nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, org)
}

func List(c *gin.Context) {
	db, err := neo4j.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	orgs, err := organizations.GetOrganizations(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func Create(c *gin.Context) {
	db, err := neo4j.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	newOrganization := organizations.New(organizations.NewOrganizationInput{})
	if err := c.BindJSON(&newOrganization); err != nil {
		return
	}

	orgs, err := organizations.CreateOrganization(*newOrganization, db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, *orgs)
}

func Delete(c *gin.Context) {
	db, err := neo4j.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	id := c.Param("orgId")

	organizations.DeleteOrganization(id, db)
}
