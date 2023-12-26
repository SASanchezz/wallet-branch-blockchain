package tx_queries

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SaveTransactionQuery(dbTransaction neo4j.ExplicitTransaction, transactionData *Transaction) {
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

	result, err := dbTransaction.Run(ctx, query, params)
	if err != nil {
		panic(err)
	}

	resultSummary, err := result.Consume(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created %v nodes in %+v.\n",
		resultSummary.Counters().NodesCreated(),
		resultSummary.ResultAvailableAfter())
}
