package tx_queries

import (
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetInterrelatedAddresses(dbTx neo4j.ManagedTransaction, address *common.Address) []*neo4j.Record {
	params := map[string]interface{}{
		"rootHash": src.GenesisTxHash.ToString(),
		"address":  address.ToString(),
	}

	template := "MATCH (r:Transaction {hash: toString($rootHash)}) " +
		"WITH r " +
		"OPTIONAL MATCH (r)-[fromRels:HAS_CHILD {from: toString($address)}]->(:Transaction) " +
		"WITH r, COLLECT(distinct fromRels.to) as toAddresses " +
		"OPTIONAL MATCH (r)-[toRels:HAS_CHILD {to: toString($address)}]->(:Transaction) " +
		"WITH toAddresses, COLLECT(distinct toRels.from) as fromAddresses " +
		"RETURN toAddresses, fromAddresses"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/get_interrelated_addresses.txt").
		Build()

	return query.Run()
}
