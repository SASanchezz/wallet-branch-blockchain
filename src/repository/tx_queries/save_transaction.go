package tx_queries

import (
	"context"
	"wallet-branch-blockchain/src/common"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SaveTransactionQuery(dbTx neo4j.ExplicitTransaction, transactionData *common.Transaction) {
	ctx := context.Background()

	params := map[string]interface{}{
		"hash":                 transactionData.Hash.ToString(),
		"parentHash":           transactionData.ParentHash.ToString(),
		"gas":                  int64(*transactionData.Gas),
		"gasPrice":             transactionData.GasPrice.String(),
		"maxFeePerGas":         transactionData.MaxFeePerGas.String(),
		"maxPriorityFeePerGas": transactionData.MaxPriorityFeePerGas.String(),
		"value":                transactionData.Value.String(),
		"nonce":                int64(*transactionData.Nonce),
	}
	query := "MERGE (t:Transaction { " +
		"hash: $hash, " +
		"parentHash: $parentHash, " +
		"gas: $gas, " +
		"gasPrice: $gasPrice, " +
		"maxFeePerGas: $maxFeePerGas, " +
		"maxPriorityFeePerGas: $maxPriorityFeePerGas, " +
		"value: $value, " +
		"nonce: $nonce}) " +
		"RETURN t"

	if _, err := dbTx.Run(ctx, query, params); err != nil {
		panic(err)
	}
}
