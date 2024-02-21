package blockchain

import (
	"fmt"
	"net/http"
	"time"
	"wallet-branch-blockchain/src/api/blockchain/payloads"
	"wallet-branch-blockchain/src/api/utilities"
	"wallet-branch-blockchain/src/common"
	"wallet-branch-blockchain/src/repository/tx_queries"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *Service
}

func NewController(blService *Service) *Controller {
	return &Controller{blService}
}

func (cont *Controller) GetInterrelatedAddresses(c *gin.Context) {
	var input payloads.GetInterrelatedAddresses
	if parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c); !(parseOk && validOk) {
		return
	}

	start := time.Now()

	addresses := cont.Service.GetInterrelatedAddresses(input.Address)

	elapsed := time.Since(start)
	fmt.Printf("Request took %s\n", elapsed)

	c.JSON(http.StatusOK, addresses)
}

func (cont *Controller) GetByHash(c *gin.Context) {
	var input payloads.GetByHash
	if parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c); !(parseOk && validOk) {
		return
	}

	transaction := cont.Service.GetByHash(input.Hash)

	c.JSON(http.StatusOK, transaction)
}

func (cont *Controller) GetBranch(c *gin.Context) {
	var input payloads.GetBranch //TODO: refactoring place
	if parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c); !(parseOk && validOk) {
		return
	}

	getBranchParams := tx_queries.GetBranchParams{
		From:   common.StringToAddress(input.From),
		To:     common.StringToAddress(input.To),
		Limit:  &input.Limit,
		Before: &input.Before,
		After:  &input.After,
	}

	branch := cont.Service.GetBranch(&getBranchParams)

	c.JSON(http.StatusOK, branch)
}

func (cont *Controller) GetAddresses(c *gin.Context) {
	addresses := cont.Service.GetAddresses()

	c.JSON(http.StatusOK, addresses)
}
