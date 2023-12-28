package tx_queries

import (
	"context"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateRelationshipQuery(
	dbTx neo4j.ExplicitTransaction,
	parent *common.Hash,
	childTransaction *common.Transaction,
) {
	ctx := context.Background()

	params := map[string]interface{}{
		"parent": parent.ToString(),
		"child":  childTransaction.Hash.ToString(),
	}
	query := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD]->(t2)"

	if _, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	}
}

func CreateBranchRelationshipQuery(
	dbTx neo4j.ExplicitTransaction,
	parentHash *common.Hash,
	childTransaction *common.Transaction,
) {
	ctx := context.Background()

	params := map[string]interface{}{
		"parent": parentHash.ToString(),
		"child":  childTransaction.Hash.ToString(),
		"from":   childTransaction.From.ToString(),
		"to":     childTransaction.To.ToString(),
	}

	query := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD {from: $from, to: $to}]->(t2)"

	if _, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	}

}
