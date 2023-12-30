package random

import (
	"math/big"
	"math/rand"
	"strings"
	"wallet-branch-blockchain/src/common"
)

var addresses = []*common.Address{
	common.StringToAddress(strings.Repeat("0", 20)),
	common.StringToAddress(strings.Repeat("1", 20)),
	common.StringToAddress(strings.Repeat("2", 20)),
	common.StringToAddress(strings.Repeat("3", 20)),
	common.StringToAddress(strings.Repeat("4", 20)),
	common.StringToAddress(strings.Repeat("5", 20)),
}

func GetRandomAddress() *common.Address {
	return addresses[rand.Intn(len(addresses))]
}

func GetRandomUint64() *uint64 {
	value := rand.Uint64()
	return &value
}

func GetRandomBigInt() *big.Int {
	n := int64(rand.Int())

	return big.NewInt(n)
}

func GetRandomBytes(length int) []byte {
	byteArray := make([]byte, length)
	_, err := rand.Read(byteArray)

	if err != nil {
		panic(err)
	}

	return byteArray
}
