package tx_queries

import (
	"context"
	"time"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/logger"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetBranch(dbTx neo4j.ManagedTransaction, inputParams *GetBranchParams) *[]*neo4j.Record {
	ctx := context.Background()
	var records []*neo4j.Record
	if *inputParams.After == 0 {
		*inputParams.After = 0
	}
	if *inputParams.Before == 0 {
		*inputParams.Before = 9223372036854775807 // max int64
	}

	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"from":     inputParams.From.ToString(),
		"to":       inputParams.To.ToString(),
		"after":    inputParams.After,
		"before":   inputParams.Before,
		"limit":    inputParams.Limit,
	}
	query := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"OPTIONAL MATCH (r)-[:HAS_CHILD {from: toString($from), to: toString($to)}]->(t1:Transaction) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WHERE (t2.timestamp >= $after AND t2.timestamp <= $before) " +
		"WITH r, t1, t2 " +
		"ORDER BY t2.timestamp ASC " +
		"WITH COLLECT(DISTINCT t2) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT r) AS allNodes " +
		"RETURN allNodes[0..$limit] as allNodes"

	start := time.Now()

	if result, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	} else if records, err = result.Collect(ctx); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	logger := logger.Logger{
		Path: "../logs/get_branch.txt",
	}

	logger.Log(elapsed.String())

	return &records
}
