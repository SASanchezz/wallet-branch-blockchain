package main

import (
	"wallet-branch-blockchain/src/core"
	"wallet-branch-blockchain/src/random"
	"wallet-branch-blockchain/src/transaction"
)

func generateTransactions() {
	ts := transaction.New()

	transactionArgs := random.GetRandomTransactions(5)
	for _, transactionArg := range *transactionArgs {
		transaction := ts.GenerateTransaction(transactionArg)
		transaction.Hash = core.GetHash(&transaction)
		ts.SaveTransaction(transaction)
	}
}
