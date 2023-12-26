package tx_queries

import (
	"context"
	"fmt"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateRelationshipQuery(
	dbTransaction neo4j.ExplicitTransaction,
	parent *common.Hash,
	childTransaction *Transaction,
) {
	ctx := context.Background()

	params := map[string]interface{}{
		"parent": parent.ToString(),
		"child":  childTransaction.Hash.ToString(),
	}
	query := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD]->(t2)"

	result, err := dbTransaction.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}

	resultSummary, _ := result.Consume(ctx)

	fmt.Printf("Created %v relationships in %+v.\n",
		resultSummary.Counters().RelationshipsCreated(),
		resultSummary.ResultAvailableAfter())
}

func CreateBranchRelationshipQuery(
	dbTransaction neo4j.ExplicitTransaction,
	parentHash *common.Hash,
	childTransaction *Transaction,
) {
	ctx := context.Background()

	query := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD {from: $from, to: $to}]->(t2)"

	params := map[string]interface{}{
		"parent": parentHash.ToString(),
		"child":  childTransaction.Hash.ToString(),
		"from":   childTransaction.From.ToString(),
		"to":     childTransaction.To.ToString(),
	}
	result, err := dbTransaction.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}

	resultSummary, _ := result.Consume(ctx)

	fmt.Printf("Created %v relationships in %+v.\n",
		resultSummary.Counters().RelationshipsCreated(),
		resultSummary.ResultAvailableAfter())
}
