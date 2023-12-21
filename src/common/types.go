package common

const (
	HashLength      = 64
	AddressLength   = 20
	BranchKeyLength = AddressLength*2 + 1
)

type TransactionHash [HashLength]byte
type BranchKey [BranchKeyLength]byte
type Address [AddressLength]byte
type Hash [HashLength]byte
