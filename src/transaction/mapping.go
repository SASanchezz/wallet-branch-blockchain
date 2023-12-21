package transaction

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/transaction/tx_queries"
)

func mapTransaction(properties map[string]any) tx_queries.TransactionData {
	hash := properties["hash"].(string)
	gas := uint64(properties["gas"].(int64))
	gasPrice := properties["gasPrice"].(string)
	maxFeePerGas := properties["maxFeePerGas"].(string)
	maxPriorityFeePerGas := properties["maxPriorityFeePerGas"].(string)
	value := properties["value"].(string)
	nonce := uint64(properties["nonce"].(int64))

	return tx_queries.TransactionData{
		Hash:                 common.StringToMyHash(hash),
		Gas:                  &gas,
		GasPrice:             common.StringToBigInt(gasPrice),
		MaxFeePerGas:         common.StringToBigInt(maxFeePerGas),
		MaxPriorityFeePerGas: common.StringToBigInt(maxPriorityFeePerGas),
		Value:                common.StringToBigInt(value),
		Nonce:                &nonce,
	}
}
