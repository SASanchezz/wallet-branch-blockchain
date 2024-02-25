package tx_queries

import (
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GBranch(dbTx neo4j.ManagedTransaction, inputParams *GBranchParams) []*neo4j.Record {
	params := map[string]interface{}{
		"from":   inputParams.From.ToString(),
		"to":     inputParams.To.ToString(),
		"after":  inputParams.After,
		"before": inputParams.Before,
		"limit":  inputParams.Limit,
	}
	template := "MATCH (b:BaseTransaction {from: $from, to: $to}) " +

		"OPTIONAL MATCH (b)-[:HAS_CHILD*]->(t:Transaction) " +
		"WHERE (t.timestamp >= $after AND t.timestamp <= $before) " +
		"WITH b, t " +

		"ORDER BY t.timestamp ASC " +
		"WITH COLLECT(DISTINCT b) + COLLECT(DISTINCT t) AS allNodes " +
		"RETURN allNodes[0..$limit] as allNodes"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/get_branch.txt").
		Build()

	return query.Run()
}
