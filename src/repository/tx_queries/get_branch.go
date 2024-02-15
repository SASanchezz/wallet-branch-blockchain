package tx_queries

import (
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetBranch(dbTx neo4j.ManagedTransaction, inputParams *GetBranchParams) []*neo4j.Record {
	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"from":     inputParams.From.ToString(),
		"to":       inputParams.To.ToString(),
		"after":    inputParams.After,
		"before":   inputParams.Before,
		"limit":    inputParams.Limit,
	}
	template := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"OPTIONAL MATCH (r)-[rel:HAS_CHILD]->(t1:Transaction) " +
		"WHERE (rel.from = toString($from) AND rel.to = toString($to)) " +
		"OR (rel.to = toString($from) AND rel.from = toString($to)) " +
		"OPTIONAL MATCH (t1)-[:HAS_CHILD*]->(t2:Transaction) " +
		"WHERE (t2.timestamp >= $after AND t2.timestamp <= $before) " +
		"WITH r, t1, t2 " +
		"ORDER BY t2.timestamp ASC " +
		"WITH COLLECT(DISTINCT t2) + COLLECT(DISTINCT t1) + COLLECT(DISTINCT r) AS allNodes " +
		"RETURN allNodes[0..$limit] as allNodes"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/get_branch.txt").
		Build()

	return query.Run()
}
