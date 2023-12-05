package organizations

import (
	"folk/proforma/core/actions/organizations"
	"folk/proforma/core/model"
	"folk/proforma/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Describe(c *gin.Context) {
	db, err := database.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	organizations.SetDataStore(db)
	id := c.Param("orgId")
	org, err := organizations.GetOrganization(id)
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
	db, err := database.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	organizations.SetDataStore(db)
	orgs, err := organizations.GetOrganizations()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func Create(c *gin.Context) {
	db, err := database.Instance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	var org model.Organization
	err = c.BindJSON(org)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	organizations.SetDataStore(db)
	o, err := organizations.CreateOrganization(org.Name, org.Description)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	//orgs = append(orgs, *newOrganization)
	c.IndentedJSON(http.StatusCreated, *o)
}
