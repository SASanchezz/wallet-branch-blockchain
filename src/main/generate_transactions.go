package main

import (
	"wallet-branch-blockchain/src/random"
	"wallet-branch-blockchain/src/transaction"
)

func generateTransactions() {
	ts := transaction.New()

	transactionArgs := random.GetRandomTransactions(25)
	for _, transactionArg := range *transactionArgs {
		ts.SaveTransaction(ts.GenerateTransaction(transactionArg))
	}
}
