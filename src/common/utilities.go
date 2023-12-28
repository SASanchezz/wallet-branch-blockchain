package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func StringToMyHash(s string) *Hash {
	var result Hash
	copy(result[:], s)
	return &result
}

func StringToMyAddress(s string) *Address {
	var result Address
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

func GetFromAddress(tx *types.Transaction) *Address {
	// from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	// if err != nil {
	// 	from, err = types.Sender(types.HomesteadSigner{}, tx)
	// }

	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)

	if err != nil {
		panic(err)
	}

	return BytesToAddress([]byte(from.String()))

}
