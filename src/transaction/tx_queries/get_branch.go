package tx_queries

import (
	"context"
	"fmt"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetBranch(dbTransaction neo4j.ManagedTransaction, branchKey *common.BranchKey) *[]*neo4j.Record {
	ctx := context.Background()
	params := map[string]interface{}{
		"rootHash":  string(src.GenesisTxHash[:]),
		"branchKey": string(branchKey[:]),
	}
	query := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"OPTIONAL MATCH (r)-[:HAS_CHILD {hash: toString($branchKey)}]->(t1:Transaction) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WITH COLLECT(DISTINCT r) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT t2) AS allNodes " +
		"RETURN allNodes"

	result, err := dbTransaction.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}

	records, err := result.Collect(ctx)

	fmt.Println("rootHash: " + string(src.GenesisTxHash[:]))
	fmt.Println("branchKey: " + string(branchKey[:]))

	return &records
}
