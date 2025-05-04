package api

import (
	"go-blockchain-explorer/internal/eth"
	"go-blockchain-explorer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/blocks/latest", getLatestBlocks)
	r.GET("/block/:number", getBlockByNumber)
	r.GET("/block/:number/transactions", getBlockTransactions)
}

func getLatestBlocks(c *gin.Context) {
	var response models.BlockListResponse

	blocks, err := eth.GetLatestBlocks(5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Items = blocks

	c.JSON(http.StatusOK, response)
}

func getBlockByNumber(c *gin.Context) {
	number := c.Param("number")
	block, err := eth.GetBlockByNumber(number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, block)
}

func getBlockTransactions(c *gin.Context) {
	var response models.TransactionListResponse

	number := c.Param("number")
	txs, err := eth.GetTransactionsByBlockNumber(number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Items = txs

	c.JSON(http.StatusOK, response)
}
