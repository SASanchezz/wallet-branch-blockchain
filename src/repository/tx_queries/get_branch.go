package tx_queries

import (
	"context"
	"wallet-branch-blockchain/src"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetBranch(dbTx neo4j.ManagedTransaction, inputParams *GetBranchParams) *neo4j.Record {
	ctx := context.Background()
	var record *neo4j.Record
	if inputParams.After == "" {
		inputParams.After = "0"
	}
	if inputParams.Before == "" {
		inputParams.Before = "9223372036854775807" // max int64
	}

	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"from":     inputParams.From.ToString(),
		"to":       inputParams.To.ToString(),
		"after":    inputParams.After,
		"before":   inputParams.Before,
	}

	query := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"OPTIONAL MATCH (r)-[:HAS_CHILD {from: toString($from), to: toString($to)}]->(t1:Transaction) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WITH r, t1, t2 " +
		"ORDER BY t2.timestamp ASC " +
		"WHERE (r.timestamp >= tofloat($after) AND r.timestamp <= tofloat($before)) " +
		"OR (t1.timestamp >= tofloat($after) AND t1.timestamp <= tofloat($before)) " +
		"OR (t2.timestamp >= tofloat($after) AND t2.timestamp <= tofloat($before))	" +
		"WITH COLLECT(DISTINCT t2) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT r) AS allNodes " +
		"RETURN allNodes"

	if result, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	} else if record, err = result.Single(ctx); err != nil {
		panic(err)
	}

	return record
}
