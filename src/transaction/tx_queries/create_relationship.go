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
	child *common.Hash,
) {
	ctx := context.Background()

	query := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD]->(t2)"

	params := map[string]interface{}{
		"parent": parent.ToString(),
		"child":  child.ToString(),
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

func CreateBranchRelationshipQuery(
	dbTransaction neo4j.ExplicitTransaction,
	parent *common.Hash,
	child *common.Hash,
	branchKey *common.BranchKey,
) {
	ctx := context.Background()

	query := "MATCH (t1:Transaction {hash: $parent}) " +
		"MATCH (t2:Transaction {hash: $child})" +
		"CREATE (t1)-[:HAS_CHILD {hash: $branchKey}]->(t2)"

	params := map[string]interface{}{
		"parent":    parent.ToString(),
		"child":     child.ToString(),
		"branchKey": branchKey.ToString(),
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
