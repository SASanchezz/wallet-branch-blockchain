package repository

import (
	"context"
	"fmt"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/database"
	"wallet-branch-blockchain/src/repository/tx_queries"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type Repository struct {
	Ctx     context.Context
	Session neo4j.SessionWithContext
}

func New() *Repository {
	return &Repository{
		Ctx:     context.Background(),
		Session: database.Connect().NewSession(context.Background(), neo4j.SessionConfig{}),
	}
}

func (r *Repository) Close() {
	r.Session.Close(r.Ctx)
}

func (r *Repository) SaveTransaction(transactionData *common.Transaction, withRelationship bool) {
	if dbTransaction, err := r.Session.BeginTransaction(r.Ctx, func(*neo4j.TransactionConfig) {}); err != nil {
		panic(err)
	} else {
		defer dbTransaction.Close(r.Ctx)
		tx_queries.SaveTransactionQuery(dbTransaction, transactionData)

		if withRelationship {
			if *transactionData.ParentHash != *src.GenesisTxHash {
				tx_queries.CreateRelationshipQuery(dbTransaction, transactionData.ParentHash, transactionData)
			} else {
				tx_queries.CreateBranchRelationshipQuery(dbTransaction, src.GenesisTxHash, transactionData)
			}
		}

		if err := dbTransaction.Commit(r.Ctx); err != nil {
			panic(err)
		}
	}
}

func (r *Repository) SaveTransactions(transactionsData *[]*common.Transaction, withRelationship bool) {
	dbTransaction, err := r.Session.BeginTransaction(r.Ctx, func(*neo4j.TransactionConfig) {})
	defer dbTransaction.Close(r.Ctx)
	if err != nil {
		panic(err)
	}

	for _, transactionData := range *transactionsData {
		tx_queries.SaveTransactionQuery(dbTransaction, transactionData)
		if withRelationship && *transactionData.ParentHash != *src.GenesisTxHash {
			tx_queries.CreateRelationshipQuery(dbTransaction, transactionData.ParentHash, transactionData)
		} else if withRelationship {
			tx_queries.CreateBranchRelationshipQuery(
				dbTransaction,
				src.GenesisTxHash,
				transactionData,
			)
		}
	}

	err = dbTransaction.Commit(r.Ctx)
	if err != nil {
		panic(err)
	}
}

func (r *Repository) GetBranch(from *common.Address, to *common.Address) *[]tx_queries.NodeData {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return *tx_queries.GetBranch(tx, from, to), nil
	})
	if err != nil {
		panic(err)
	}
	records := result.([]*neo4j.Record)

	allNodes, _ := records[0].Get("allNodes")
	data := allNodes.([]interface{})
	transactions := make([]tx_queries.NodeData, len(data))

	fmt.Println("allNodes", allNodes)

	for i, node := range data {
		properties := node.(dbtype.Node).Props
		transactions[i] = mapTransaction(properties)
	}

	return &transactions
}

func (r *Repository) GetTransaction(hash *common.Hash) *tx_queries.NodeData {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
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

func (r *Repository) GetLastBranchTransaction(from *common.Address, to *common.Address) *tx_queries.NodeData {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return *tx_queries.GetLastTransaction(tx, from, to), nil
	})
	if err != nil {
		panic(err)
	}
	record := result.(neo4j.Record)

	lastNode, _ := record.Get("lastNode")
	node := lastNode.(interface{})

	fmt.Println("node", node)

	properties := node.(dbtype.Node).Props
	lastTransaction := mapTransaction(properties)

	return &lastTransaction
}
