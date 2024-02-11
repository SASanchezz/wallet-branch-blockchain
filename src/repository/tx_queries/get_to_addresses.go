package tx_queries

import (
	"context"
	"fmt"
	"time"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/logger"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetToAddresses(dbTx neo4j.ManagedTransaction, from *common.Address) *[]*neo4j.Record {
	ctx := context.Background()
	var records []*neo4j.Record

	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"from":     from.ToString(),
	}

	query := "MATCH (t:Transaction {hash: toString($rootHash)})- " +
		"[rels:HAS_CHILD {from: toString($from)}]->(:Transaction) " +
		"RETURN rels"

	start := time.Now()

	if result, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	} else if records, err = result.Collect(ctx); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	logger := logger.Logger{
		Path: "../logs/get_to_addresses.txt",
	}

	fmt.Println("GetTransaction elapsed time: ", elapsed)
	logger.LogInt64(int64(elapsed / time.Microsecond))

	return &records
}
