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

	tx_queries.CTransactionN(dbTransaction, transactionData)

	if withRelationship {
		if *transactionData.ParentHash != *src.GenesisTxHash {
			tx_queries.CHasChildRelQueryQuery(dbTransaction, transactionData.ParentHash, transactionData)
		} else {
			tx_queries.CBaseHasChildRelQuery(dbTransaction, src.GenesisTxHash, transactionData)
		}
	}

	if err := dbTransaction.Commit(r.Ctx); err != nil {
		panic(err)
	}
}

func (r *Repository) GBranch(params *tx_queries.GBranchParams) *tx_queries.Branch {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx_queries.GBranch(tx, params), nil
	})
	if err != nil {
		panic(err)
	}
	records := result.([]*neo4j.Record)

	parsedTransactions := *parseTransactions(records, "allNodes")

	return &parsedTransactions
}

func (r *Repository) GInterrelatedAddresses(address *common.Address) tx_queries.InterrelatedAddresses {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx_queries.GInterrelatedAddresses(tx, address), nil
	})
	if err != nil {
		panic(err)
	}
	record := result.([]*neo4j.Record)

	if len(record) == 0 {
		return tx_queries.InterrelatedAddresses{
			FromAddresses: []string{},
			ToAddresses:   []string{},
		}
	}

	toAddressesRecord, _ := record[0].Get("toAddresses")
	fromAddressesRecord, _ := record[0].Get("fromAddresses")

	interrelatedAddresses := tx_queries.InterrelatedAddresses{
		FromAddresses: mapAddresses(fromAddressesRecord.([]interface{})),
		ToAddresses:   mapAddresses(toAddressesRecord.([]interface{})),
	}

	return interrelatedAddresses
}

func (r *Repository) GTransaction(hash *common.Hash) *tx_queries.NodeData {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		if record := tx_queries.GTransaction(tx, hash); record == nil {
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

func (r *Repository) GLastTransaction(from *common.Address, to *common.Address) (*tx_queries.NodeData, *tx_queries.RelationshipData) {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx_queries.GLastTransaction(tx, from, to), nil
	})
	if err != nil {
		panic(err)
	}

	record := result.(*neo4j.Record)
	lastNode, _ := record.Get("lastNode")
	rel, _ := record.Get("rel")

	if rel == nil {
		return mapTransaction(lastNode.(dbtype.Node).Props), nil
	}

	return mapTransaction(lastNode.(dbtype.Node).Props), mapRelationship(rel.(dbtype.Relationship).Props)
}

func (r *Repository) GAddresses() []string {
	result, err := r.Session.ExecuteRead(r.Ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx_queries.GAddresses(tx), nil
	})
	if err != nil {
		panic(err)
	}
	record := result.(*neo4j.Record)
	uniqueAddresses, _ := record.Get("uniqueAddresses")

	return mapAddresses(uniqueAddresses.([]interface{}))
}
