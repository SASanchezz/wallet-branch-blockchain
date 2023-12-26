package tx_queries

import (
	"context"
	"fmt"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetBranch(dbTransaction neo4j.ManagedTransaction, from *common.Address, to *common.Address) *[]*neo4j.Record {
	ctx := context.Background()
	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"from":     from.ToString(),
		"to":       to.ToString(),
	}
	query := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"OPTIONAL MATCH (r)-[:HAS_CHILD {from: toString($from), to: toString($to)}]->(t1:Transaction) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WITH COLLECT(DISTINCT r) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT t2) AS allNodes " +
		"RETURN allNodes"

	result, err := dbTransaction.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}

	records, err := result.Collect(ctx)

	fmt.Println("rootHash: " + src.GenesisTxHash.ToString())
	fmt.Println("from: " + from.ToString())
	fmt.Println("to: " + to.ToString())

	return &records
}
