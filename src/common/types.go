package common

import (
	"math/big"
)

const (
	HashLength      = 64
	AddressLength   = 20
	BranchKeyLength = AddressLength*2 + 1
)

type TransactionHash [HashLength]byte
type BranchKey [BranchKeyLength]byte
type Address [AddressLength]byte
type Hash [HashLength]byte

type Addresses []*Address

type Transaction struct {
	Hash                 *Hash    `json:"hash"`
	ParentHash           *Hash    `json:"parentHash"`
	From                 *Address `json:"from"`
	To                   *Address `json:"to"`
	Gas                  *uint64  `json:"gas"`
	GasPrice             *big.Int `json:"gasPrice"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
	Value                *big.Int `json:"value"`
	Nonce                *uint64  `json:"nonce"`
}
