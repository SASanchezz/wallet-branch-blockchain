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

func (s *Service) GetToAddresses(from string) []string {
	addresses := transaction.New().GetToAddresses(common.StringToAddress(from))
	return addresses.ToString()
}

func (s *Service) GetFromAddresses(to string) []string {
	addresses := transaction.New().GetFromAddresses(common.StringToAddress(to))
	return addresses.ToString()
}

func (s *Service) GetByHash(hash string) *tx_queries.NodeData {
	return transaction.New().GetTransaction(common.StringToMyHash(hash))
}

func (s *Service) GetBranch(from string, to string) *tx_queries.Branch {
	branch := transaction.New().GetBranch(common.StringToAddress(from), common.StringToAddress(to))
	return branch
}
