package random

import (
	"math/big"
	"math/rand"
	"wallet-branch-blockchain/src/common"
)

var addresses = []*common.Address{
	{0x02},
	{0x03},
	{0x04},
	{0x05},
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
