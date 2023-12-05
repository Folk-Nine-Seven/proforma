package database

import (
	"context"
	"fmt"
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

func (db *Neo4j) Create(newOrg model.Organization) (*model.Organization, error) {
	ctx := context.Background()
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		cmd := fmt.Sprintf(`CREATE (a:Organization) SET a.id = "%s" SET a.name = "%s" SET a.description = "%s"`, newOrg.Id, newOrg.Name, newOrg.Description)
		result, err := transaction.Run(ctx, cmd, nil)
		if err != nil {
			return nil, err
		}
		return newOrg, result.Err()
	})
	if err != nil {
		return nil, err
	}
	return &newOrg, nil
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

	if len(orgs) > 0 {
		return &orgs[0], nil
	}

	return nil, nil
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
