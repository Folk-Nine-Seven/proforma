package neo4j

import (
	"context"
	"fmt"
	"folk/proforma/core/model"
	"folk/proforma/gateways/secrets"

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

var instance *Neo4j

func Instance() (*Neo4j, error) {
	if instance == nil {
		sm := secrets.NewSecretsManager()
		password, err := sm.GetSecret(secrets.Neo4j)
		if err != nil {
			return nil, err
		}

		instance = New(&gin.Context{})
		instance.Initialize(InitializeInput{
			Uri:      "neo4j+s://1adaebe1.databases.neo4j.io:7687",
			User:     "neo4j",
			Password: password, // "Zutnassi1qY7Ip1GTENC5l7TXTgbNYHJuayZydriLUU",
		})
	}
	return instance, nil
}

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

func (db *Neo4j) CreateOrganization(newOrg model.Organization) (*model.Organization, error) {
	ctx := context.Background()
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		cmd := fmt.Sprintf(`CREATE (a:Organization) SET a.id = "%s" SET a.name = "%s" SET a.description = "%s"`, newOrg.Id, newOrg.Name, newOrg.Description)
		result, err := transaction.Run(ctx, cmd, nil)
		if err != nil {
			return nil, fmt.Errorf("this is not good: %s", err)
		}
		return newOrg, result.Err()
	})
	if err != nil {
		return nil, err
	}
	return &newOrg, nil
}

func (db *Neo4j) DeleteOrganization(id string) error {
	result, err := neo4j.ExecuteQuery(
		db.context,
		db.Driver,
		fmt.Sprintf(`MATCH (n:Organization {id: '%s'})
		DELETE n`, id),
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (db *Neo4j) UpdateOrganization(id string, changes *model.Organization) (*model.Organization, error) {
	return nil, nil
}

func (db *Neo4j) DescribeOrganization(id string) (*model.Organization, error) {
	result, err := neo4j.ExecuteQuery(
		db.context, db.Driver,
		NewMatch("Organization", map[string]any{"id": id}).Serialize(),
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

func (db *Neo4j) ListOrganizations() ([]model.Organization, error) {
	result, err := neo4j.ExecuteQuery(db.context, db.Driver,
		NewMatch("Organization", nil).Serialize(),
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

func (db *Neo4j) CreateProject(orgId string, newProj model.Project) (*model.Project, error) {
	ctx := context.Background()
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		cmd := fmt.Sprintf(`MATCH (b:Organization) WHERE b.id = "%s" CREATE (a:Project)-[c:MEMBER_OF]->(b) SET a.id = "%s" SET a.name = "%s" SET a.description = "%s"`, orgId, newProj.Id, newProj.Name, newProj.Description)
		result, err := transaction.Run(ctx, cmd, nil)
		if err != nil {
			return nil, err
		}
		return newProj, result.Err()
	})
	if err != nil {
		return nil, err
	}
	return &newProj, nil
}

func (db *Neo4j) DescribeProject(orgId, projId string) (model.Project, error) {
	return model.Project{}, nil
}

func (db *Neo4j) UpdateProject(orgId, projId string, proj model.Project) (model.Project, error) {
	return model.Project{}, nil
}

func (db *Neo4j) DeleteProject(orgId, projId string) error {
	return nil
}

func (db *Neo4j) ListProjects(orgId string) ([]model.Project, error) {
	return []model.Project{}, nil
}
