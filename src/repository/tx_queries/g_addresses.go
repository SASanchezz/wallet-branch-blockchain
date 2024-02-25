package tx_queries

import (
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GAddresses(dbTx neo4j.ManagedTransaction) *neo4j.Record {
	template := "MATCH (b:BaseTransaction)" +
		"WITH collect(DISTINCT b.from) + collect(DISTINCT b.to) AS allAddresses " +
		"RETURN apoc.coll.toSet(allAddresses) AS uniqueAddresses"

	query := core.NewQueryBuilder(dbTx).
		WithTemplate(template).
		WithLogPath("../logs/get_addresses.txt").
		Build()

	return query.Run()[0]
}
