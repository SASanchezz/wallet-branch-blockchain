package transaction

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/core"
	"wallet-branch-blockchain/src/repository"
	"wallet-branch-blockchain/src/repository/tx_queries"
)

type TransactionService struct {
	Repository *repository.Repository
}

func New() *TransactionService {
	return &TransactionService{
		Repository: repository.New(),
	}
}

func (ts *TransactionService) Close() {
	ts.Repository.Close()
}

func (ts *TransactionService) GenerateTransaction(transaction *common.Transaction) *common.Transaction {
	lastTransaction := ts.Repository.GetLastBranchTransaction(transaction.From, transaction.To)

	transaction.Hash = core.GetHash(&transaction)
	transaction.ParentHash = lastTransaction.Hash

	return transaction
}

func (ts *TransactionService) SaveTransaction(transactionData *common.Transaction) {
	ts.Repository.SaveTransaction(transactionData, true)
}

func (ts *TransactionService) GetLastTransaction(from *common.Address, to *common.Address) *tx_queries.NodeData {
	return ts.Repository.GetLastBranchTransaction(from, to)
}

func (ts *TransactionService) GetBranch(from *common.Address, to *common.Address) *[]tx_queries.NodeData {
	return ts.Repository.GetBranch(from, to)
}

func (ts *TransactionService) GetTransaction(hash *common.Hash) *tx_queries.NodeData {
	return ts.Repository.GetTransaction(hash)
}
