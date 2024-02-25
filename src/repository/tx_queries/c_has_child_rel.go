package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CHasChildRelQuery(
	dbTx neo4j.ExplicitTransaction,
	parent *common.Hash,
	childTransaction *common.Transaction,
) {
	params := map[string]interface{}{
		"parent": parent.ToString(),
		"child":  childTransaction.Hash.ToString(),
	}
	template := "MATCH (t2:Transaction {hash: $child}), (t1) " +
		"WHERE ANY(x IN ['BaseTransaction', 'Transaction'] WHERE x IN labels(t1)) " +
		"AND t1.hash = $parent " +
		"CREATE (t1)-[:HAS_CHILD]->(t2)"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/create_relationship.txt").
		Build()

	query.Run()
}
