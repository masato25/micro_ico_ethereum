package token_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/model/tokens"
	log "github.com/sirupsen/logrus"
)

func TokenDapp(c *gin.Context) {
	contractID := c.Param("id")
	converId, _ := strconv.Atoi(contractID)
	var tokenItem tokens.Token
	dt := db.Where("id = ?", converId).First(&tokenItem)
	if dt.Error != nil {
		log.Error(dt.Error.Error())
		return
	}
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "token_dapp", gin.H{
		"login":        lg,
		"contractAddr": tokenItem.ContractAddress,
		"contractName": tokenItem.TokenName,
		"contractAbi":  tokenItem.Abi,
		"symbol":       tokenItem.TokenSymbol,
	})
	return
}
