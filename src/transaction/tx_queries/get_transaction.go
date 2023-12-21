package tx_queries

import (
	"context"
	"fmt"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetTransaction(dbTransaction neo4j.ManagedTransaction, txHash *common.Hash) *neo4j.Record {
	ctx := context.Background()
	params := map[string]interface{}{
		"hash": string(txHash[:]),
	}
	query := "MATCH (t:Transaction {hash: toString($hash)}) RETURN t"

	result, err := dbTransaction.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}
	record, err := result.Single(ctx)

	fmt.Println("record: ", record)

	return record
}
