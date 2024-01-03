package random

import (
	"wallet-branch-blockchain/src/common"
)

func GetRandomTransaction() *common.Transaction {
	return &common.Transaction{
		From:                 GetRandomAddress(),
		To:                   GetRandomAddress(),
		Gas:                  GetRandomUint64(),
		GasPrice:             GetRandomBigInt(),
		MaxFeePerGas:         GetRandomBigInt(),
		MaxPriorityFeePerGas: GetRandomBigInt(),
		Value:                GetRandomBigInt(),
		Timestamp:            GetRandomTimestamp(),
		Nonce:                GetRandomUint64(),
	}
}

func GetRandomTransactions(amount int) *[]*common.Transaction {
	transactions := make([]*common.Transaction, amount)

	for i := 0; i < amount; i++ {
		transactions[i] = GetRandomTransaction()
	}

	return &transactions
}
