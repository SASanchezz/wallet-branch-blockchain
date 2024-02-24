package transaction

import (
	"math/big"
	"wallet-branch-blockchain/src"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/tx_queries"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (ts *TransactionService) CreateGenesisBlock() {
	gas := uint64(1)
	nonce := uint64(7)
	timestamp := uint64(0)
	genesisTx := common.Transaction{
		Hash:                 src.GenesisTxHash,
		ParentHash:           nil,
		From:                 nil,
		To:                   nil,
		Gas:                  &gas,
		GasPrice:             big.NewInt(0),
		MaxFeePerGas:         big.NewInt(0),
		MaxPriorityFeePerGas: big.NewInt(0),
		Value:                big.NewInt(0),
		Timestamp:            &timestamp,
		Nonce:                &nonce,
	}

	dbTransaction, _ := ts.Repository.Session.BeginTransaction(ts.Repository.Ctx, func(*neo4j.TransactionConfig) {})
	defer dbTransaction.Close(ts.Repository.Ctx)

	tx_queries.CTransactionN(dbTransaction, &genesisTx)

	err := dbTransaction.Commit(ts.Repository.Ctx)
	if err != nil {
		panic(err)
	}
}
