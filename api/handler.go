package api

import (
	"net/http"
	"go-blockchain-explorer/internal/eth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/blocks/latest", getLatestBlocks)
	r.GET("/block/:number", getBlockByNumber)
}

func getLatestBlocks(c *gin.Context) {
	blocks, err := eth.GetLatestBlocks(5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blocks)
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