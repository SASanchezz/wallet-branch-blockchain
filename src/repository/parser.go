package repository

import (
	"wallet-branch-blockchain/src/repository/tx_queries"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func parseTransactions(records []*neo4j.Record, key string, limit int8) *tx_queries.Branch {
	transactions := tx_queries.Branch{}
	for _, record := range records {
		allNodes, _ := record.Get(key)
		data := allNodes.([]interface{})
		branch := *mapTransactions(data)
		transactions = append(transactions, branch...)
	}
	return &transactions
}
