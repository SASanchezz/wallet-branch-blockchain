package transaction

import (
	"context"
	"fmt"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/core"
	"wallet-branch-blockchain/src/database"
	"wallet-branch-blockchain/src/transaction/tx_queries"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

func saveTransaction(transactionData tx_queries.Transaction, withRelationship bool) {
	ctx := context.Background()
	driver := database.Connect()
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	dbTransaction, err := session.BeginTransaction(ctx, func(*neo4j.TransactionConfig) {})
	defer dbTransaction.Close(ctx)

	if err != nil {
		panic(err)
	}
	tx_queries.SaveTransactionQuery(dbTransaction, transactionData)
	if withRelationship && *transactionData.ParentHash != *src.GenesisTxHash {
		tx_queries.CreateRelationshipQuery(dbTransaction, transactionData.ParentHash, transactionData.Hash)
	} else if withRelationship {
		tx_queries.CreateBranchRelationshipQuery(
			dbTransaction,
			src.GenesisTxHash,
			transactionData.Hash,
			core.GetBranchKey(transactionData.From, transactionData.To))
	}

	err = dbTransaction.Commit(ctx)
	if err != nil {
		panic(err)
	}
}

func getBranch(branchKey *common.BranchKey) *[]tx_queries.TransactionData {
	ctx := context.Background()
	driver := database.Connect()
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return *tx_queries.GetBranch(tx, branchKey), nil
	})
	if err != nil {
		panic(err)
	}
	records := result.([]*neo4j.Record)

	allNodes, _ := records[0].Get("allNodes")
	data := allNodes.([]interface{})
	transactions := make([]tx_queries.TransactionData, len(data))

	fmt.Println("allNodes", allNodes)

	for i, node := range data {
		properties := node.(dbtype.Node).Props
		transactions[i] = mapTransaction(properties)
	}

	return &transactions
}

func getTransaction(hash *common.Hash) *tx_queries.TransactionData {
	ctx := context.Background()
	driver := database.Connect()
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		record := tx_queries.GetTransaction(tx, hash)
		if record == nil {
			return nil, nil
		}
		return *record, nil
	})
	if err != nil {
		panic(err)
	}
	if result == nil {
		return nil
	}
	record := result.(neo4j.Record)

	t, _ := record.Get("t")
	node := t.(dbtype.Node)

	properties := node.Props
	tx := mapTransaction(properties)

	return &tx
}
