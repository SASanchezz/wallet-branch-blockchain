package main

import (
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/bootstrap"
	"wallet-branch-blockchain/src/random"
	"wallet-branch-blockchain/src/transaction"
)

func generateTransactions() {
	ts := transaction.New()

	if ts.GetTransaction(src.GenesisTxHash) == nil {
		bootstrap.CreateGenesisBlock()
	}

	transactionArgs := random.GetRandomTransactions(25)
	for _, transactionArg := range *transactionArgs {
		ts.SaveTransaction(ts.GenerateTransaction(transactionArg))
	}
}
