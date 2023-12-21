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

func (hash *Hash) ToString() string {
	if hash == nil {
		return ""
	}
	returnVal := string(hash[:])
	return returnVal
}

func (branchKey *BranchKey) ToString() string {
	if branchKey == nil {
		return ""
	}
	return string(branchKey[:])
}
