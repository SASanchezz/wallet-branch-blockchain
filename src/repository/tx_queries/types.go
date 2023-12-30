package tx_queries

import (
	"math/big"
)

type NodeData struct {
	Hash                 *string  `json:"hash"`
	Gas                  *uint64  `json:"gas"`
	GasPrice             *big.Int `json:"gasPrice"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
	Value                *big.Int `json:"value"`
	Nonce                *uint64  `json:"nonce"`
}

type Branch []*NodeData

type Branches []*Branch
