package database

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type (
	Neo4j struct {
		context *context.Context
		Driver  neo4j.DriverWithContext
		session neo4j.SessionWithContext
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
