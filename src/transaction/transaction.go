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
	transaction.ParentHash = common.StringToMyHash(*lastTransaction.Hash)

	return transaction
}

func (ts *TransactionService) SaveTransaction(transactionData *common.Transaction) {
	ts.Repository.SaveTransaction(transactionData, true)
}

func (ts *TransactionService) GetLastTransaction(from *common.Address, to *common.Address) *tx_queries.NodeData {
	return ts.Repository.GetLastBranchTransaction(from, to)
}

func (ts *TransactionService) GetBranch(params *tx_queries.GetBranchParams) *tx_queries.Branch {
	return ts.Repository.GetBranch(params)
}

func (ts *TransactionService) GetToAddresses(from *common.Address) *common.Addresses {
	return ts.Repository.GetToAddresses(from)
}

func (ts *TransactionService) GetFromAddresses(to *common.Address) *common.Addresses {
	return ts.Repository.GetFromAddresses(to)
}

func (ts *TransactionService) GetTransaction(hash *common.Hash) *tx_queries.NodeData {
	return ts.Repository.GetTransaction(hash)
}
