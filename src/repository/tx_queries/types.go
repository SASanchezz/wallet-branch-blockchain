package tx_queries

import (
	"math/big"
	"wallet-branch-blockchain/src/common"
)

type GetBranchParams struct {
	From   *common.Address
	To     *common.Address
	Limit  int64
	Before string
	After  string
}

type NodeData struct {
	Hash                 *string  `json:"hash"`
	Gas                  *uint64  `json:"gas"`
	GasPrice             *big.Int `json:"gasPrice"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
	Value                *big.Int `json:"value"`
	Timestamp            *uint64  `json:"timestamp"`
	Nonce                *uint64  `json:"nonce"`
}

type Branch []*NodeData

type Branches []*Branch
