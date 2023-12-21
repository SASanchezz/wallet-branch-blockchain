package main

import (
	"fmt"
	"wallet-branch-blockchain/src/bootstrap"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/transaction"
)

func main() {
	bootstrap.LoadEnv()

	hash := common.StringToMyHash("f963525f16b863cc488d963e39062b6cb4f63cdd0ea97fb0c7bf585444efe676")
	transaction := transaction.GetTransaction(hash)

	fmt.Println(transaction)

	// generateTransaction()
}
