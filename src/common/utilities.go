package common

import (
	"encoding/hex"
	"math/big"
)

func StringToMyHash(s string) *Hash {
	var result Hash
	copy(result[:], s)
	return &result
}

func StringToHexMyHash(s string) *Hash {
	if len(s) != 32 {
		panic("StringToMyHash: s is not 32 bytes long")
	}
	var result Hash
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	copy(result[:], data)
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
