package database

import (
	"context"
	"fmt"
	"folk/proforma/core/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type (
	Neo4j struct {
		context *context.Context
		Driver  neo4j.DriverWithContext
	}

	InitializeInput struct {
		Uri      string
		User     string
		Password string
	}
)

func New(ctx context.Context) *Neo4j {
	return &Neo4j{
		context: &ctx,
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
	return db.Driver.Close(*db.context)
}

func (db *Neo4j) VerifyConnectivity() error {
	return db.Driver.VerifyConnectivity(*db.context)
}

func (db *Neo4j) Create(name, description string) (*model.Organization, error) {
	return nil, nil
}

func (db *Neo4j) Delete(id string) error {
	return nil
}

func (db *Neo4j) Update(id string, changes *model.Organization) (*model.Organization, error) {
	return nil, nil
}

func (db *Neo4j) Describe(id string) (*model.Organization, error) {
	result, err := neo4j.ExecuteQuery(
		*db.context, db.Driver,
		fmt.Sprintf("MATCH (o:Organization WHERE o.id = '%s') RETURN o", id),
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println(result.Records[0].Values...)

	return &model.Organization{}, nil
}

func (db *Neo4j) List() ([]model.Organization, error) {
	return nil, nil
}
