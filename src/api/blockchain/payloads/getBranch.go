package payloads

import "wallet-branch-blockchain/src/common"

type GetBranch struct {
	From   string `form:"from" binding:"required"`
	To     string `form:"to" binding:"required"`
	Before string `form:"before"`
	After  string `form:"after"`
}

func (payload GetBranch) Validate() (bool, string) {
	if len(payload.From) != common.AddressLength {
		return false, "'from' address is invalid"
	}

	if len(payload.To) != common.AddressLength {
		return false, "'to' address is invalid"
	}

	return true, ""
}
