package core

import (
	"encoding/hex"
	"encoding/json"
	"sync"
	"wallet-branch-blockchain/src/common"

	"golang.org/x/crypto/sha3"
)

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

func GetHash(x interface{}) *common.Hash {
	hash := sha3.New256()
	var buf []byte
	b, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	hash.Write(b)
	buf = hash.Sum(nil)

	return common.StringToMyHash(hex.EncodeToString(buf))
}

func GBranchKey(from *common.Address, to *common.Address) *common.BranchKey {
	branchKey := common.BranchKey{}
	copy(branchKey[:common.AddressLength], from[:])
	branchKey[common.AddressLength] = '-'
	copy(branchKey[common.AddressLength+1:common.BranchKeyLength], to[:])

	return &branchKey
}
