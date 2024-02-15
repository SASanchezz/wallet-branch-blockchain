package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateBaseRelationshipQuery(
	dbTx neo4j.ExplicitTransaction,
	parentHash *common.Hash,
	childTransaction *common.Transaction,
) {
	params := map[string]interface{}{
		"parent": parentHash.ToString(),
		"child":  childTransaction.Hash.ToString(),
		"from":   childTransaction.From.ToString(),
		"to":     childTransaction.To.ToString(),
	}

	template := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD {from: $from, to: $to}]->(t2)"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/create_base_relationship.txt").
		Build()

	query.Run()
}
