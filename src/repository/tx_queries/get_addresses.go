package tx_queries

import (
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetAddresses(dbTx neo4j.ManagedTransaction) *neo4j.Record {
	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
	}
	template := "MATCH (r:Transaction {hash: toString($rootHash)})-[rel:HAS_CHILD]->(t1:Transaction)" +
		"WHERE (rel.from IS NOT NULL) AND (rel.to IS NOT NULL) " +
		"WITH collect(DISTINCT rel.from) + collect(DISTINCT rel.to) AS allAddresses " +
		"RETURN apoc.coll.toSet(allAddresses) AS uniqueAddresses"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/get_address.txt").
		Build()

	return query.Run()[0]
}
