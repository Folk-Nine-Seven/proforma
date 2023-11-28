package database

import (
	"context"
)

var instance *Neo4j

func Instance() (*Neo4j, error) {
	if instance == nil {
		context := context.Background()
		instance = New(context)
		instance.Initialize(InitializeInput{
			Uri:      "neo4j+s://1adaebe1.databases.neo4j.io:7687",
			User:     "neo4j",
			Password: "Zutnassi1qY7Ip1GTENC5l7TXTgbNYHJuayZydriLUU",
		})
	}
	return instance, nil
}
