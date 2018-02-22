package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/controller/common_helper"
	"github.com/masato25/micro_ico_ethereum/app/controller/token_controller"
	"github.com/masato25/micro_ico_ethereum/app/controller/user_controller"
	"github.com/masato25/micro_ico_ethereum/app/controller/wallet_controller"
)

func Routes(route *gin.Engine) {
	route.Use(common_helper.SessionAuth())
	token_controller.Route(route)
	user_controller.Route(route)
	wallet_controller.Route(route)
}
