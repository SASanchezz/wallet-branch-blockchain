package tx_queries

import (
	"context"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetFromAddresses(dbTx neo4j.ManagedTransaction, to *common.Address) *[]*neo4j.Record {
	ctx := context.Background()
	var records []*neo4j.Record

	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"to":       to.ToString(),
	}

	query := "MATCH (t:Transaction {hash: toString($rootHash)})- " +
		"[rels:HAS_CHILD {to: toString($to)}]->(:Transaction) " +
		"RETURN rels"

	if result, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	} else if records, err = result.Collect(ctx); err != nil {
		panic(err)
	}

	return &records
}
