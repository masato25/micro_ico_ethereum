package wallet_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "wallet_index", gin.H{"login": lg})
	return
}
