package routing

import "github.com/gin-gonic/gin"

var projects map[string]map[string]func(c *gin.Context) = map[string]map[string]func(c *gin.Context){
	"/organizations/:orgId/projects": {
		"GET":  listProjects,
		"POST": createProject,
	},
	"/organizations/:orgId/projects/:id": {
		"GET":    describeProject,
		"PATCH":  nil,
		"DELETE": nil,
	},
}

func listProjects(c *gin.Context) {
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

func describeProject(c *gin.Context) {
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
