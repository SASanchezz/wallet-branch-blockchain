package bootstrap

import (
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/transaction"
)

func CreateGenesisBlock() {
	ts := transaction.New()
	defer ts.Close()
	if ts.GetTransaction(src.GenesisTxHash) == nil {
		ts.CreateGenesisBlock()
	}
}
