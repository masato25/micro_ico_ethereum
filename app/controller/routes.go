package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/controller/token_controller"
	"github.com/masato25/micro_ico_ethereum/app/controller/user_controller"
)

func Routes(route *gin.Engine) {
	token_controller.Route(route)
	user_controller.Route(route)
}
