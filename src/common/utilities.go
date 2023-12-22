package common

import (
	"math/big"
)

func StringToMyHash(s string) *Hash {
	var result Hash
	copy(result[:], s)
	return &result
}

func StringToBigInt(s string) *big.Int {
	var result big.Int
	result.SetString(s, 10)
	return &result
}

func BytesToAddress(b []byte) *Address {
	var result Address
	copy(result[:], b)
	return &result
}
