package wallet_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/masato25/micro_ico_ethereum/app/setup"
)

var db *gorm.DB

func Route(r *gin.Engine) {
	db = setup.GetConn()
	// html views
	r.GET("/wallet", Index)
}
