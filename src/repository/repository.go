package repository

import (
	"context"
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
	dbTransaction, err := r.Session.BeginTransaction(r.Ctx, func(*neo4j.TransactionConfig) {})
	if err != nil {
		panic(err)
	}
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

func (r *Repository) GetBranch(params *tx_queries.GetBranchParams) *tx_queries.Branch {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return *tx_queries.GetBranch(tx, params), nil
	})
	if err != nil {
		panic(err)
	}
	records := result.([]*neo4j.Record)

	parsedTransactions := *parseTransactions(records, "allNodes", int8(*params.Limit))

	return &parsedTransactions
}

func (r *Repository) GetToAddresses(from *common.Address) *common.Addresses {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return *tx_queries.GetToAddresses(tx, from), nil
	})
	if err != nil {
		panic(err)
	}
	records := result.([]*neo4j.Record)
	addresses := make(common.Addresses, len(records))
	for i, record := range records {
		rel, _ := record.Get("rels")
		addresses[i] = mapAddress(rel.(dbtype.Relationship).Props, "to")
	}

	return &addresses
}

func (r *Repository) GetFromAddresses(to *common.Address) *common.Addresses {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return *tx_queries.GetFromAddresses(tx, to), nil
	})
	if err != nil {
		panic(err)
	}
	records := result.([]*neo4j.Record)
	addresses := make(common.Addresses, len(records))
	for i, record := range records {
		rel, _ := record.Get("rels")
		addresses[i] = mapAddress(rel.(dbtype.Relationship).Props, "from")
	}

	return &addresses
}

func (r *Repository) GetTransaction(hash *common.Hash) *tx_queries.NodeData {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		if record := tx_queries.GetTransaction(tx, hash); record == nil {
			return nil, nil
		} else {
			return *record, nil
		}
	})
	if err != nil {
		panic(err)
	}
	if result == nil {
		return nil
	}

	record := result.(neo4j.Record)
	t, _ := record.Get("t")

	return mapTransaction(t.(dbtype.Node).Props)
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

	return mapTransaction(lastNode.(dbtype.Node).Props)
}
