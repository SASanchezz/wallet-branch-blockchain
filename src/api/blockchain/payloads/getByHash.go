package payloads

import "wallet-branch-blockchain/src/common"

type GetByHash struct {
	Hash string `form:"hash" binding:"required"`
}

func (payload GetByHash) Validate() (bool, string) {
	if len(payload.Hash) != common.HashLength {
		return false, "'hash' is invalid"

	}

	return true, ""
}
