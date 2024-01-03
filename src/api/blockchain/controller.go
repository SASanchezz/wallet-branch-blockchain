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

func (cont *Controller) GetToAddresses(c *gin.Context) {
	var input payloads.GetToAddresses
	if parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c); !(parseOk && validOk) {
		return
	}

	start := time.Now()

	addresses := cont.Service.GetToAddresses(input.From)

	elapsed := time.Since(start)
	fmt.Printf("Request took %s\n", elapsed)

	c.JSON(http.StatusOK, addresses)
}

func (cont *Controller) GetFromAddresses(c *gin.Context) {
	var input payloads.GetFromAddresses
	if parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c); !(parseOk && validOk) {
		return
	}

	start := time.Now()

	addresses := cont.Service.GetFromAddresses(input.To)

	elapsed := time.Since(start)
	fmt.Printf("Request took %s\n", elapsed)

	c.JSON(http.StatusOK, addresses)
}

func (cont *Controller) GetByHash(c *gin.Context) {
	var input payloads.GetByHash
	if parseOk, validOk := utilities.ParseInput(&input, c), utilities.ValidateInput(input, c); !(parseOk && validOk) {
		return
	}

	start := time.Now()

	transaction := cont.Service.GetByHash(input.Hash)

	elapsed := time.Since(start)
	fmt.Printf("Request took %s\n", elapsed)

	c.JSON(http.StatusOK, transaction)
}

func (cont *Controller) GetBranch(c *gin.Context) {
	var input payloads.GetBranch
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

	start := time.Now()

	branch := cont.Service.GetBranch(&getBranchParams)

	elapsed := time.Since(start)
	fmt.Printf("Request took %s\n", elapsed)

	c.JSON(http.StatusOK, branch)
}
