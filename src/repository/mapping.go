package repository

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/tx_queries"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

func mapTransactions(properties []interface{}, limit int64) *tx_queries.Branch {
	transactions := make(tx_queries.Branch, limit)

	for i, node := range properties {
		if int64(i) == limit {
			break
		}

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
	timestamp := uint64(properties["timestamp"].(int64))
	nonce := uint64(properties["nonce"].(int64))

	return &tx_queries.NodeData{
		Hash:                 &hash,
		Gas:                  &gas,
		GasPrice:             common.StringToBigInt(gasPrice),
		MaxFeePerGas:         common.StringToBigInt(maxFeePerGas),
		MaxPriorityFeePerGas: common.StringToBigInt(maxPriorityFeePerGas),
		Value:                common.StringToBigInt(value),
		Timestamp:            &timestamp,
		Nonce:                &nonce,
	}
}

func mapAddress(properties map[string]any, fieldName string) *common.Address {
	return common.StringToAddress(properties[fieldName].(string))
}
