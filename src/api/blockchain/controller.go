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

func (cont *Controller) GInterrelatedAddresses(c *gin.Context) {
	var input payloads.GInterrelatedAddresses
	if parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c); !(parseOk && validOk) {
		return
	}

	start := time.Now()

	addresses := cont.Service.GInterrelatedAddresses(input.Address)

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

func (cont *Controller) GBranch(c *gin.Context) {
	var input payloads.GBranch
	parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c)
	if !(parseOk && validOk) {
		return
	}

	getBranchParams := tx_queries.GBranchParams{
		From:   common.StringToAddress(input.From),
		To:     common.StringToAddress(input.To),
		Limit:  &input.Limit,
		Before: &input.Before,
		After:  &input.After,
	}

	branch := cont.Service.GBranch(&getBranchParams)

	c.JSON(http.StatusOK, branch)
}

func (cont *Controller) GAddresses(c *gin.Context) {
	addresses := cont.Service.GAddresses()

	c.JSON(http.StatusOK, addresses)
}
