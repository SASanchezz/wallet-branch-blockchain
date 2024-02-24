package transaction

import (
	"wallet-branch-blockchain/src/common"
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
	lastTransaction, _ := ts.GLastTransaction(transaction.From, transaction.To)
	transaction.ParentHash = common.StringToMyHash(*lastTransaction.Hash)

	return transaction
}

func (ts *TransactionService) SaveTransaction(transactionData *common.Transaction) {
	ts.Repository.SaveTransaction(transactionData, true)
}

func (ts *TransactionService) GLastTransaction(from *common.Address, to *common.Address) (*tx_queries.NodeData, *tx_queries.RelationshipData) {
	return ts.Repository.GLastTransaction(from, to)
}

func (ts *TransactionService) GBranch(params *tx_queries.GBranchParams) *tx_queries.Branch {
	return ts.Repository.GBranch(params)
}

func (ts *TransactionService) GInterrelatedAddresses(address *common.Address) tx_queries.InterrelatedAddresses {
	return ts.Repository.GInterrelatedAddresses(address)
}

func (ts *TransactionService) GTransaction(hash *common.Hash) *tx_queries.NodeData {
	return ts.Repository.GTransaction(hash)
}

func (ts *TransactionService) GAddresses() []string {
	return ts.Repository.GAddresses()
}
