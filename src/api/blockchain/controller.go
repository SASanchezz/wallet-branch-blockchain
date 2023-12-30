package blockchain

import (
	"net/http"
	"wallet-branch-blockchain/src/api/blockchain/payloads"
	"wallet-branch-blockchain/src/api/utilities"

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

	addresses := cont.Service.GetBranch(input.From, input.To)
	c.JSON(http.StatusOK, addresses)
}
