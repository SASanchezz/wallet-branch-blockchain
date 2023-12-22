package bootstrap

import (
	"context"
	"math/big"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/database"
	"wallet-branch-blockchain/src/transaction/tx_queries"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateGenesisBlock() {
	gas := uint64(1)
	nonce := uint64(7)
	genesisTx := tx_queries.Transaction{
		Hash:                 src.GenesisTxHash,
		ParentHash:           nil,
		From:                 nil,
		To:                   nil,
		Gas:                  &gas,
		GasPrice:             big.NewInt(0),
		MaxFeePerGas:         big.NewInt(0),
		MaxPriorityFeePerGas: big.NewInt(0),
		Value:                big.NewInt(0),
		Nonce:                &nonce,
	}
	ctx := context.Background()
	driver := database.Connect()
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	dbTransaction, _ := session.BeginTransaction(ctx, func(*neo4j.TransactionConfig) {})
	defer dbTransaction.Close(ctx)

	tx_queries.SaveTransactionQuery(dbTransaction, &genesisTx)

	err := dbTransaction.Commit(ctx)
	if err != nil {
		panic(err)
	}
}
