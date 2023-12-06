package secrets

import (
	"context"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type (
	SecretsManager struct {
		secrets map[string]string
	}
)

const (
	Neo4j = "neo4j"
)

var keys map[string]string = map[string]string{
	Neo4j: "projects/25497090578/secrets/neo4j/versions/1",
}

func NewSecretsManager() *SecretsManager {
	return &SecretsManager{
		secrets: make(map[string]string),
	}
}

func (sm *SecretsManager) GetSecret(key string) (string, error) {
	if val, ok := sm.secrets[key]; ok {
		return val, nil
	}
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	secret, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: keys[Neo4j],
	})
	if err != nil {
		return "", err
	}
	sm.secrets[key] = string(secret.GetPayload().Data)
	return sm.secrets[key], nil
}
