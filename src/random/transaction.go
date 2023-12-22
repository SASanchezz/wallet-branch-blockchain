package random

import "wallet-branch-blockchain/src/transaction"

func GetRandomTransaction() *transaction.TransactionArgs {
	return &transaction.TransactionArgs{
		From:                 GetRandomAddress(),
		To:                   GetRandomAddress(),
		Gas:                  GetRandomUint64(),
		GasPrice:             GetRandomBigInt(),
		MaxFeePerGas:         GetRandomBigInt(),
		MaxPriorityFeePerGas: GetRandomBigInt(),
		Value:                GetRandomBigInt(),
		Nonce:                GetRandomUint64(),
	}
}

func GetRandomTransactions(amount int) *[]*transaction.TransactionArgs {
	transactions := make([]*transaction.TransactionArgs, amount)

	for i := 0; i < amount; i++ {
		transactions[i] = GetRandomTransaction()
	}

	return &transactions
}
