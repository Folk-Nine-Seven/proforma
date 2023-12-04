package database

import (
	"fmt"
	"folk/proforma/core/actions/organizations"
	"folk/proforma/core/model"
	"folk/proforma/database/queries"

	"github.com/gin-gonic/gin"
	"github.com/isaacp/alchem"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type (
	Neo4j struct {
		context *gin.Context
		Driver  neo4j.DriverWithContext
	}

	InitializeInput struct {
		Uri      string
		User     string
		Password string
	}
)

func New(ctx *gin.Context) *Neo4j {
	return &Neo4j{
		context: ctx,
	}
}

func (db *Neo4j) Initialize(input InitializeInput) {
	driver, err := neo4j.NewDriverWithContext(
		input.Uri,
		neo4j.BasicAuth(input.User, input.Password, ""),
	)
	if err != nil {
		panic(err)
	}
	db.Driver = driver

}

func (db *Neo4j) Close() error {
	return db.Driver.Close(db.context)
}

func (db *Neo4j) VerifyConnectivity() error {
	return db.Driver.VerifyConnectivity(db.context)
}

func (db *Neo4j) Create(name, description string) (*model.Organization, error) {
	session := db.Driver.NewSession(db.context, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(db.context)

	newOrganization := organizations.New(organizations.NewOrganizationInput{})
	c := (*gin.Context)(db.context)
	if err := c.BindJSON(&newOrganization); err != nil {
		return nil, err
	}
	org, err := session.ExecuteWrite(c, func(transaction neo4j.ManagedTransaction) (any, error) {
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
		return nil, err
	}

	retval = org.(model.Organization)
	return retval, nil
}

func (db *Neo4j) Delete(id string) error {
	return nil
}

func (db *Neo4j) Update(id string, changes *model.Organization) (*model.Organization, error) {
	return nil, nil
}

func (db *Neo4j) Describe(id string) (*model.Organization, error) {
	result, err := neo4j.ExecuteQuery(
		db.context, db.Driver,
		queries.NewMatch("Organization", map[string]any{"id": id}).Serialize(),
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, err
	}

	orgs, err := alchem.TransformObject[[]model.Organization](result.Records, "map(.Values.[].Props)")
	if err != nil {
		return nil, err
	}

	return &orgs[0], nil
}

func (db *Neo4j) List() ([]model.Organization, error) {
	result, err := neo4j.ExecuteQuery(db.context, db.Driver,
		queries.NewMatch("Organization", nil).Serialize(),
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		return nil, err
	}

	orgs, err := alchem.TransformObject[[]model.Organization](result.Records, "map(.Values.[].Props)")
	if err != nil {
		return nil, err
	}

	return orgs, nil
}
