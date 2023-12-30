package api

import (
	"wallet-branch-blockchain/src/api/blockchain"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	blockchainController := blockchain.NewController(blockchain.NewService())

	router.GET("/", blockchainController.GetByHash)
	router.GET("/branch", blockchainController.GetBranch)
	router.GET("/to", blockchainController.GetToAddresses)
	router.GET("/from", blockchainController.GetFromAddresses)

	router.Run(":8080")
}
