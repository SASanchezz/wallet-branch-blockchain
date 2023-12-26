package transaction

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/core"
	"wallet-branch-blockchain/src/transaction/tx_queries"
)

func GenerateTransaction(transactionArgs *TransactionArgs) *tx_queries.Transaction {
	lastTransaction, err := GetLastTransaction(transactionArgs.From, transactionArgs.To)
	if err != nil {
		panic(err)
	}

	if lastTransaction.Hash == nil {
		panic(err)
	}

	transaction := tx_queries.Transaction{
		Hash:                 &common.Hash{},
		ParentHash:           lastTransaction.Hash,
		From:                 transactionArgs.From,
		To:                   transactionArgs.To,
		Gas:                  transactionArgs.Gas,
		GasPrice:             transactionArgs.GasPrice,
		MaxFeePerGas:         transactionArgs.MaxFeePerGas,
		MaxPriorityFeePerGas: transactionArgs.MaxPriorityFeePerGas,
		Value:                transactionArgs.Value,
		Nonce:                transactionArgs.Nonce,
	}

	transaction.Hash = core.GetHash(&transaction)

	return &transaction
}

func SaveTransaction(transactionData *tx_queries.Transaction) {
	saveTransaction(transactionData, true)
}

func GetLastTransaction(from *common.Address, to *common.Address) (*tx_queries.TransactionData, error) {
	branchTransactions := GetBranch(from, to)

	if branchTransactions == nil {
		return &tx_queries.TransactionData{}, nil
	}
	derefedBranchTransactions := *branchTransactions

	return &derefedBranchTransactions[len(derefedBranchTransactions)-1], nil
}

func GetBranch(from *common.Address, to *common.Address) *[]tx_queries.TransactionData {
	return getBranch(from, to)
}

func GetTransaction(hash *common.Hash) *tx_queries.TransactionData {
	return getTransaction(hash)
}
