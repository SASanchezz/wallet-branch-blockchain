package api

import (
	"wallet-branch-blockchain/src/api/blockchain"
	"wallet-branch-blockchain/src/api/middlewares"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	blockchainController := blockchain.NewController(blockchain.NewService())

	router.GET("/", blockchainController.GetByHash)
	router.GET("/branch", blockchainController.GetBranch)
	router.GET("/to", blockchainController.GetToAddresses)
	router.GET("/from", blockchainController.GetFromAddresses)
	router.GET("/addresses", blockchainController.GetAddresses)

	router.Run("127.0.0.1:8080")
}
