package tx_queries

import (
	"math/big"
	"wallet-branch-blockchain/src/common"
)

type TransactionData struct {
	Hash                 *common.Hash `json:"hash"`
	Gas                  *uint64      `json:"gas"`
	GasPrice             *big.Int     `json:"gasPrice"`
	MaxFeePerGas         *big.Int     `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int     `json:"maxPriorityFeePerGas"`
	Value                *big.Int     `json:"value"`
	Nonce                *uint64      `json:"nonce"`
}

type Transaction struct {
	Hash                 *common.Hash    `json:"hash"`
	ParentHash           *common.Hash    `json:"parentHash"`
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *uint64         `json:"gas"`
	GasPrice             *big.Int        `json:"gasPrice"`
	MaxFeePerGas         *big.Int        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int        `json:"maxPriorityFeePerGas"`
	Value                *big.Int        `json:"value"`
	Nonce                *uint64         `json:"nonce"`
}
