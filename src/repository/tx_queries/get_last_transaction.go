package tx_queries

import (
	"context"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetLastTransaction(dbTx neo4j.ManagedTransaction, from *common.Address, to *common.Address) *neo4j.Record {
	ctx := context.Background()
	var record *neo4j.Record

	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"from":     from.ToString(),
		"to":       to.ToString(),
	}
	query := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"OPTIONAL MATCH (r)-[:HAS_CHILD {from: toString($from), to: toString($to)}]->(t1:Transaction) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WITH COLLECT(DISTINCT r) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT t2) AS allNodes " +
		"WITH last(allNodes) AS lastNode " +
		"RETURN lastNode"

	if result, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	} else if record, err = result.Single(ctx); err != nil {
		panic(err)
	}

	return record
}
