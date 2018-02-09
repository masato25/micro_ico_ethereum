package token_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/masato25/micro_ico_ethereum/services/ethereumrpc"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "hello",
	})
	return
}

func CreateErc20Token(c *gin.Context) {
	// ethereumrpc.DepolyContactToBlockChain()
}
