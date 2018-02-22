package token_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/masato25/micro_ico_ethereum/app/setup"
)

var db *gorm.DB

func Route(r *gin.Engine) {
	db = setup.GetConn()

	//html views
	r.GET("tokens", TokenList)
	r.GET("tokens/new", TokenCreate)
	r.GET("token_dapp/:id", TokenDapp)
	rapi := r.Group("/api/v1/tokens")
	rapi.GET("/", TokenListJSON)
	rapi.POST("/new", CreateErc20Token)
}
