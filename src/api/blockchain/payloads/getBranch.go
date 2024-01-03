package payloads

import "wallet-branch-blockchain/src/common"

type GetBranch struct {
	From   string `form:"from" binding:"required"`
	To     string `form:"to" binding:"required"`
	Before int64  `form:"before"`
	After  int64  `form:"after"`
	Limit  int64  `form:"limit"`
}

func (payload GetBranch) Validate() (bool, string) {
	if len(payload.From) != common.AddressLength {
		return false, "'from' address is invalid"
	}

	if len(payload.To) != common.AddressLength {
		return false, "'to' address is invalid"
	}

	if (payload.Limit) < 0 {
		return false, "'limit' must be positive"
	}
	if (payload.Limit) > 100 {
		return false, "'limit' must be less than 100"
	}

	return true, ""
}
