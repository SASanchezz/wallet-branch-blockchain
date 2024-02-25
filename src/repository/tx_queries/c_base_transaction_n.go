package tx_queries

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/core"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CBaseTransactionN(dbTx neo4j.ExplicitTransaction, transactionData *common.Transaction) {
	params := map[string]interface{}{
		"from":                 transactionData.From.ToString(),
		"to":                   transactionData.To.ToString(),
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
	template := "MERGE (b:BaseTransaction { " +
		"from: $from, " +
		"to: $to, " +
		"hash: $hash, " +
		"parentHash: $parentHash, " +
		"gas: $gas, " +
		"gasPrice: $gasPrice, " +
		"maxFeePerGas: $maxFeePerGas, " +
		"maxPriorityFeePerGas: $maxPriorityFeePerGas, " +
		"value: $value, " +
		"timestamp: $timestamp, " +
		"nonce: $nonce}) " +
		"RETURN b"

	query := core.NewQueryBuilder(dbTx).
		WithParams(params).
		WithTemplate(template).
		WithLogPath("../logs/save_transaction.txt").
		Build()

	query.Run()
}
