package payloads

import "wallet-branch-blockchain/src/common"

type GetFromAddresses struct {
	To string `form:"to" binding:"required"`
}

func (payload GetFromAddresses) Validate() (bool, string) {
	if len(payload.To) != common.AddressLength {
		return false, "'to' address is invalid"
	}

	return true, ""
}
