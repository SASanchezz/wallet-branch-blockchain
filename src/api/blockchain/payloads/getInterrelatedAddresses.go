package payloads

import "wallet-branch-blockchain/src/common"

type GInterrelatedAddresses struct {
	Address string `form:"address" binding:"required"`
}

func (payload GInterrelatedAddresses) Validate() (bool, string) {
	if len(payload.Address) != common.AddressLength {
		return false, "address is invalid"
	}

	return true, ""
}
