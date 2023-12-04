package organizations

import (
	"fmt"
	"folk/proforma/core/actions/organizations"
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

	fmt.Println(org)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
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
	// db, err := database.Instance()
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// //orgs = append(orgs, *newOrganization)
	// c.IndentedJSON(http.StatusCreated, greeting)
}
