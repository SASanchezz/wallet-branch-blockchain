package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GInterrelatedAddresses(dbTx neo4j.ManagedTransaction, address *common.Address) []*neo4j.Record {
	params := map[string]interface{}{
		"address": address.ToString(),
	}

	template := "OPTIONAL MATCH (b1:BaseTransaction {from: $address}), (b2:BaseTransaction {to: $address}) " +
		"WITH COLLECT(distinct b1.to) as toAddresses, COLLECT(distinct b2.from) as fromAddresses " +
		"RETURN toAddresses, fromAddresses"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/get_interrelated_addresses.txt").
		Build()

	return query.Run()
}
