package tx_queries

import (
	"context"
	"time"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/logger"

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
		"OPTIONAL MATCH (r)-[rel:HAS_CHILD]->(t1:Transaction) " +
		"WHERE (rel.from = toString($from) AND rel.to = toString($to)) " +
		"OR (rel.to = toString($from) AND rel.from = toString($to)) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WITH COLLECT(DISTINCT r) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT t2) AS allNodes " +
		"WITH last(allNodes) AS lastNode " +
		"RETURN lastNode"

	start := time.Now()

	if result, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	} else if record, err = result.Single(ctx); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	logger := logger.Logger{
		Path: "../logs/get_last_transaction.txt",
	}

	logger.Log(elapsed.String())

	return record
}
