package main

import (
	"math/big"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/bootstrap"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/transaction"
)

func generateTransaction() {
	ts := transaction.New()
	defer ts.Close()

	if ts.GetTransaction(src.GenesisTxHash) == nil {
		bootstrap.CreateGenesisBlock()
	}

	gas := uint64(1)
	nonce := uint64(7)
	transactionArgs := common.Transaction{
		From:                 &common.Address{0x02},
		To:                   &common.Address{0x03},
		Gas:                  &gas,
		GasPrice:             big.NewInt(3),
		MaxFeePerGas:         big.NewInt(4),
		MaxPriorityFeePerGas: big.NewInt(5),
		Value:                big.NewInt(6),
		Nonce:                &nonce,
	}
	newTransaction := ts.GenerateTransaction(&transactionArgs)
	ts.SaveTransaction(newTransaction)
}
