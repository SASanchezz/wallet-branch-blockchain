package main

import (
	"wallet-branch-blockchain/src/api"
	"wallet-branch-blockchain/src/bootstrap"
)

func main() {
	bootstrap.LoadEnv()
	bootstrap.CreateGenesisBlock()
	generateTransactions()

	api.Run()

	// listener.Listen()
}
