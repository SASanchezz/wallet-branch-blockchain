package tx_queries

import (
	"context"
	"time"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/logger"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetAddresses(dbTx neo4j.ManagedTransaction) *neo4j.Record {
	ctx := context.Background()
	var record *neo4j.Record

	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
	}
	query := "MATCH (r:Transaction {hash: toString($rootHash)})-[rel:HAS_CHILD]->(t1:Transaction)" +
		"WHERE (rel.from IS NOT NULL) AND (rel.to IS NOT NULL) " +
		"WITH collect(DISTINCT rel.from) + collect(DISTINCT rel.to) AS allAddresses " +
		"RETURN apoc.coll.toSet(allAddresses) AS uniqueAddresses"

	start := time.Now()

	if result, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	} else if record, err = result.Single(ctx); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	logger := logger.Logger{
		Path: "../logs/get_address.txt",
	}

	logger.LogInt64(int64(elapsed / time.Millisecond))

	return record
}
