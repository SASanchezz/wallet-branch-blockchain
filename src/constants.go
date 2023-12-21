package src

import (
	"strings"
	"wallet-branch-blockchain/src/common"
)

var (
	GenesisTxHash = common.StringToMyHash(strings.Repeat("0", 64))
)
