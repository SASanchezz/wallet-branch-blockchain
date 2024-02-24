package blockchain

import (
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/tx_queries"
	"wallet-branch-blockchain/src/transaction"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GInterrelatedAddresses(address string) tx_queries.InterrelatedAddresses {
	addresses := transaction.New().GInterrelatedAddresses(common.StringToAddress(address))
	return addresses
}

func (s *Service) GetByHash(hash string) *tx_queries.NodeData {
	return transaction.New().GTransaction(common.StringToMyHash(hash))
}

func (s *Service) GBranch(params *tx_queries.GBranchParams) *tx_queries.Branch {
	if *params.After == 0 {
		*params.After = 0
	}
	if *params.Before == 0 {
		*params.Before = 9223372036854775807 // max int64
	}
	if *params.Limit == 0 {
		*params.Limit = 100
	}

	branch := transaction.New().GBranch(params)
	return branch
}

func (s *Service) GAddresses() []string {
	addresses := transaction.New().GAddresses()
	return addresses
}
