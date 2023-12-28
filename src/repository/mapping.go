package repository

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/tx_queries"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

func mapTransactions(properties []interface{}) *tx_queries.Branch {
	transactions := make(tx_queries.Branch, len(properties))

	for i, node := range properties {
		transactions[i] = mapTransaction(node.(dbtype.Node).Props)
	}
	return &transactions
}

func mapTransaction(properties map[string]any) *tx_queries.NodeData {
	hash := properties["hash"].(string)
	gas := uint64(properties["gas"].(int64))
	gasPrice := properties["gasPrice"].(string)
	maxFeePerGas := properties["maxFeePerGas"].(string)
	maxPriorityFeePerGas := properties["maxPriorityFeePerGas"].(string)
	value := properties["value"].(string)
	nonce := uint64(properties["nonce"].(int64))

	return &tx_queries.NodeData{
		Hash:                 common.StringToMyHash(hash),
		Gas:                  &gas,
		GasPrice:             common.StringToBigInt(gasPrice),
		MaxFeePerGas:         common.StringToBigInt(maxFeePerGas),
		MaxPriorityFeePerGas: common.StringToBigInt(maxPriorityFeePerGas),
		Value:                common.StringToBigInt(value),
		Nonce:                &nonce,
	}
}
