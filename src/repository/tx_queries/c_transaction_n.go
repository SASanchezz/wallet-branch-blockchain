package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CTransactionN(dbTx neo4j.ExplicitTransaction, transactionData *common.Transaction) {
	params := map[string]interface{}{
		"hash":                 transactionData.Hash.ToString(),
		"parentHash":           transactionData.ParentHash.ToString(),
		"gas":                  int64(*transactionData.Gas),
		"gasPrice":             transactionData.GasPrice.String(),
		"maxFeePerGas":         transactionData.MaxFeePerGas.String(),
		"maxPriorityFeePerGas": transactionData.MaxPriorityFeePerGas.String(),
		"value":                transactionData.Value.String(),
		"timestamp":            int64(*transactionData.Timestamp),
		"nonce":                int64(*transactionData.Nonce),
	}
	template := "MERGE (t:Transaction { " +
		"hash: $hash, " +
		"parentHash: $parentHash, " +
		"gas: $gas, " +
		"gasPrice: $gasPrice, " +
		"maxFeePerGas: $maxFeePerGas, " +
		"maxPriorityFeePerGas: $maxPriorityFeePerGas, " +
		"value: $value, " +
		"timestamp: $timestamp, " +
		"nonce: $nonce}) " +
		"RETURN t"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/save_transaction.txt").
		Build()

	query.Run()
}
