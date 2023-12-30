package payloads

import "wallet-branch-blockchain/src/common"

type GetToAddresses struct {
	From string `form:"from" binding:"required"`
}

func (payload GetToAddresses) Validate() (bool, string) {
	if len(payload.From) != common.AddressLength {
		return false, "'from' address is invalid"
	}

	return true, ""
}
