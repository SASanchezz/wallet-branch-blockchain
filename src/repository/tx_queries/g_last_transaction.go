package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GLastTransaction(dbTx neo4j.ManagedTransaction, from *common.Address, to *common.Address) *neo4j.Record {
	params := map[string]interface{}{
		"from": from.ToString(),
		"to":   to.ToString(),
	}
	template := "OPTIONAL MATCH (b:BaseTransaction {from: $from, to: $to}) " +
		"OPTIONAL MATCH path = (b)-[:HAS_CHILD*]->(t:Transaction) " +
		"WITH COLLECT(DISTINCT b) + COLLECT(DISTINCT t) AS allNodes " +
		"WITH last(allNodes) AS lastNode " +
		"RETURN lastNode"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/get_last_transaction.txt").
		Build()

	if result := query.Run(); len(result) == 0 {
		return nil
	} else {
		return result[0]
	}
}
