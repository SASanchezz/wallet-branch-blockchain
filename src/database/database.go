package database

import (
	"context"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Singleton
var driver neo4j.DriverWithContext

func Connect() neo4j.DriverWithContext {
	if driver != nil {
		return driver
	}
	ctx := context.Background()
	username, _ := os.LookupEnv("NEO4J_USERNAME")
	password, _ := os.LookupEnv("NEO4J_PASSWORD")
	dbUri, _ := os.LookupEnv("NEO4J_URI")

	dbUser := username
	dbPassword := password
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	//defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	return driver
}

func ExecuteQuery(driver neo4j.DriverWithContext, query string, params map[string]interface{}) (*neo4j.EagerResult, error) {
	ctx := context.Background()
	result, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Close(driver neo4j.DriverWithContext) {
	driver.Close(context.Background())
}
