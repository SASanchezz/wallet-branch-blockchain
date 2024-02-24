package tx_queries

import (
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GLastTransaction(dbTx neo4j.ManagedTransaction, from *common.Address, to *common.Address) *neo4j.Record {
	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"from":     from.ToString(),
		"to":       to.ToString(),
	}
	template := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"OPTIONAL MATCH (r)-[rel:HAS_CHILD]->(t1:Transaction) " +
		"WHERE (rel.from = toString($from) AND rel.to = toString($to)) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WITH rel, COLLECT(DISTINCT r) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT t2) AS allNodes " +
		"WITH rel, last(allNodes) AS lastNode " +
		"RETURN rel, lastNode"

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
