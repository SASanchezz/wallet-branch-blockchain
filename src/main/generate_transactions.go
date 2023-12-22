package main

import (
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/bootstrap"
	"wallet-branch-blockchain/src/random"
	"wallet-branch-blockchain/src/transaction"
)

func generateTransactions() {
	if transaction.GetTransaction(src.GenesisTxHash) == nil {
		bootstrap.CreateGenesisBlock()
	}

	transactionArgs := random.GetRandomTransactions(10000)
	newTransactions := transaction.GenerateTransactions(transactionArgs)
	transaction.SaveTransactions(newTransactions)
}
