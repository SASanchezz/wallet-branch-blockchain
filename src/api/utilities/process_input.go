package utilities

import (
	"net/http"
	"wallet-branch-blockchain/src/api/blockchain/payloads"

	"github.com/gin-gonic/gin"
)

func ParseInput(input interface{}, c *gin.Context) bool {
	if err := c.ShouldBindQuery(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

func ValidateInput(input payloads.Payload, c *gin.Context) bool {
	if ok, msg := input.Validate(); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return false
	}
	return true
}
