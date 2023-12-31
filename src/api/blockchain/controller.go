package blockchain

import (
	"net/http"
	"strconv"
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
	utilities.ParseInput(&input, c)
	utilities.ValidateInput(input, c)

	addresses := cont.Service.GetToAddresses(input.From)
	c.JSON(http.StatusOK, addresses)
}

func (cont *Controller) GetFromAddresses(c *gin.Context) {
	var input payloads.GetFromAddresses
	utilities.ParseInput(&input, c)
	utilities.ValidateInput(input, c)

	addresses := cont.Service.GetFromAddresses(input.To)
	c.JSON(http.StatusOK, addresses)
}

func (cont *Controller) GetByHash(c *gin.Context) {
	var input payloads.GetByHash
	utilities.ParseInput(&input, c)
	utilities.ValidateInput(input, c)

	addresses := cont.Service.GetByHash(input.Hash)
	c.JSON(http.StatusOK, addresses)
}

func (cont *Controller) GetBranch(c *gin.Context) {
	var input payloads.GetBranch
	utilities.ParseInput(&input, c)
	utilities.ValidateInput(input, c)
	limitStr := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be an integer"})
		return
	} else if limit < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be greater than 0"})
		return
	} else if limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be less than 100"})
		return
	}

	getBranchParams := tx_queries.GetBranchParams{
		From:   common.StringToAddress(input.From),
		To:     common.StringToAddress(input.To),
		Limit:  int64(limit),
		Before: input.Before,
		After:  input.After,
	}

	addresses := cont.Service.GetBranch(&getBranchParams)
	c.JSON(http.StatusOK, addresses)
}
