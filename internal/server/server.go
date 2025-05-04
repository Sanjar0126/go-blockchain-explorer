package server

import (
	"go-blockchain-explorer/api"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	api.RegisterRoutes(r)
	r.Run(":8080")
}
