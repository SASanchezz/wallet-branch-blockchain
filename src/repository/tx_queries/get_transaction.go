package tx_queries

import (
	"context"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetTransaction(dbTx neo4j.ManagedTransaction, txHash *common.Hash) *neo4j.Record {
	ctx := context.Background()
	params := map[string]interface{}{
		"hash": txHash.ToString(),
	}
	query := "MATCH (t:Transaction {hash: toString($hash)}) RETURN t"

	result, err := dbTx.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}
	record, err := result.Single(ctx)

	return record
}
