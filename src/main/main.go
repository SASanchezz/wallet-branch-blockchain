package main

import (
	"wallet-branch-blockchain/src/bootstrap"
)

func main() {
	bootstrap.LoadEnv()

	generateTransactions()
}
