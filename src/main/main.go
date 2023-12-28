package main

import (
	"wallet-branch-blockchain/src/bootstrap"
	"wallet-branch-blockchain/src/listener"
)

func main() {
	bootstrap.LoadEnv()
	bootstrap.CreateGenesisBlock()

	// generateTransactions()

	listener.Listen()
}
