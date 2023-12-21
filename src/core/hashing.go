package core

import (
	"encoding/hex"
	"sync"
	"wallet-branch-blockchain/src/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

//TODO: take function from libs to own package

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

func rlpHash(x interface{}) *common.Hash {
	h := &common.Hash{}
	tmp := [32]byte{}
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(tmp[:])

	hexString := hex.EncodeToString(tmp[:])
	copy(h[:], hexString)
	return h
}

func GetHash(x interface{}) *common.Hash {
	return rlpHash(x)
}

func GetBranchKey(from *common.Address, to *common.Address) *common.BranchKey {
	branchKey := common.BranchKey{}
	copy(branchKey[:common.AddressLength], from[:])
	branchKey[common.AddressLength] = '-'
	copy(branchKey[common.AddressLength+1:common.BranchKeyLength], to[:])

	return &branchKey
}
