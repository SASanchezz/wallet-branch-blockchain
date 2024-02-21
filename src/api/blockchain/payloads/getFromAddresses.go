package payloads

import "wallet-branch-blockchain/src/common"

type GetInterrelatedAddresses struct {
	Address string `form:"address" binding:"required"`
}

func (payload GetInterrelatedAddresses) Validate() (bool, string) {
	if len(payload.Address) != common.AddressLength {
		return false, "address is invalid"
	}

	return true, ""
}
