package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CHasChildRelQueryQuery(
	dbTx neo4j.ExplicitTransaction,
	parent *common.Hash,
	childTransaction *common.Transaction,
) {
	params := map[string]interface{}{
		"parent": parent.ToString(),
		"child":  childTransaction.Hash.ToString(),
	}
	template := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD]->(t2)"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/create_relationship.txt").
		Build()

	query.Run()
}
