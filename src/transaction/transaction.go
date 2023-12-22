package transaction

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/core"
	"wallet-branch-blockchain/src/transaction/tx_queries"
)

func GenerateTransaction(transactionArgs *TransactionArgs) *tx_queries.Transaction {
	branchKey := core.GetBranchKey(transactionArgs.From, transactionArgs.To)
	lastTransaction, err := GetLastTransaction(branchKey)
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

func GenerateTransactions(transactionArgs *[]*TransactionArgs) *[]*tx_queries.Transaction {
	resultTransactions := make([]*tx_queries.Transaction, len((*transactionArgs)[:]))

	for i, transactionArg := range *transactionArgs {
		resultTransactions[i] = GenerateTransaction(transactionArg)
	}

	return &resultTransactions
}

func SaveTransaction(transactionData *tx_queries.Transaction) {
	saveTransaction(transactionData, true)
}

func SaveTransactions(transactionsDatas *[]*tx_queries.Transaction) {
	saveTransactions(transactionsDatas, true)
}

func GetLastTransaction(branchKey *common.BranchKey) (*tx_queries.TransactionData, error) {
	branchTransactions := GetBranch(branchKey)

	if branchTransactions == nil {
		return &tx_queries.TransactionData{}, nil
	}
	derefedBranchTransactions := *branchTransactions

	return &derefedBranchTransactions[len(derefedBranchTransactions)-1], nil
}

func GetBranch(branchKey *common.BranchKey) *[]tx_queries.TransactionData {
	return getBranch(branchKey)
}

func GetTransaction(hash *common.Hash) *tx_queries.TransactionData {
	return getTransaction(hash)
}
