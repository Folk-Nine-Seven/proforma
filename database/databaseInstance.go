package database

import (
	"context"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/gin-gonic/gin"
)

var instance *Neo4j

func Instance() (*Neo4j, error) {
	if instance == nil {

		context := context.Background()
		client, err := secretmanager.NewClient(context)
		if err != nil {
			log.Fatalf("failed to setup client: %v", err)
		}
		defer client.Close()

		secret, err := client.AccessSecretVersion(context, &secretmanagerpb.AccessSecretVersionRequest{
			Name: "projects/25497090578/secrets/neo4j/versions/1",
		})

		if err != nil {
			return nil, err
		}

		password := string(secret.GetPayload().Data)

		instance = New(&gin.Context{})
		instance.Initialize(InitializeInput{
			Uri:      "neo4j+s://1adaebe1.databases.neo4j.io:7687",
			User:     "neo4j",
			Password: password, // "Zutnassi1qY7Ip1GTENC5l7TXTgbNYHJuayZydriLUU",
		})
	}
	return instance, nil
}
