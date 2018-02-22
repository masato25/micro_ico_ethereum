package common_helper

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("SessionAuth")
		user, err := GetUserBySession(c)
		c.Set("login", "true")
		if user.ID == 0 {
			c.Set("login", "false")
		} else if err != nil {
			log.Error(err)
			c.Set("login", "false")
		}
		c.Next()
	}
}
