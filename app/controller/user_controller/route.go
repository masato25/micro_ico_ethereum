package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/masato25/micro_ico_ethereum/app/setup"
)

var db *gorm.DB

func Route(r *gin.Engine) {
	db = setup.GetConn()
	// html views
	r.GET("/", Index)
	r.GET("/register", Register)
	r.GET("/login", LogIn)

	rapi := r.Group("/api/v1/users")
	rapi.POST("/resigter", RegisterJSON)
	rapi.POST("/login", LogInJSON)
	rapi.GET("/session", SessionCheck)
	rapi.GET("/logout", LogOut)
}
