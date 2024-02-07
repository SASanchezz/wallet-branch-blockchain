package tx_queries

import (
	"context"
	"time"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/logger"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetTransaction(dbTx neo4j.ManagedTransaction, txHash *common.Hash) *neo4j.Record {
	ctx := context.Background()
	params := map[string]interface{}{
		"hash": txHash.ToString(),
	}
	query := "MATCH (t:Transaction {hash: toString($hash)}) RETURN t"

	start := time.Now()

	result, err := dbTx.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	logger := logger.Logger{
		Path: "../logs/get_transaction.txt",
	}

	logger.Log(elapsed.String())

	record, err := result.Single(ctx)

	return record
}
