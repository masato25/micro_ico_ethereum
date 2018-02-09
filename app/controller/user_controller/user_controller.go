package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/model/users"
	log "github.com/sirupsen/logrus"
)

type UserRegFeilds struct {
	UserName      string `form:"username" json:"username" binding:"required"`
	Passwd        string `form:"password" json:"password" binding:"required"`
	EthAddress    string `form:"eth_address" json:"eth_address" binding:"required"`
	EthPrivateKey string `form:"eth_private_key" json:"eth_private_key" binding:"required"`
}

func Register(c *gin.Context) {
	inputs := UserRegFeilds{}
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := users.Account{
		UserName:                 inputs.UserName,
		EthereumAddress:          inputs.EthAddress,
		EthereumPrivateKeyEncode: inputs.EthPrivateKey,
	}
	passwd, err := users.HashPassword(inputs.Passwd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unknown problem"})
		log.Error(err)
		return
	}
	user.Password = passwd
	dt := db.Create(&user)
	if dt.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dt.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "user created."})
	return
}

type LoginInput struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"password" json:"password" binding:"required"`
}

func LogIn(c *gin.Context) {
	inputs := LoginInput{}
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user users.Account
	passwd, err := users.HashPassword(inputs.Passwd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unknown problem"})
		log.Error(err)
		return
	}
	dt := db.Preload("Sessions").Where(&users.Account{UserName: inputs.UserName, Password: passwd}).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	} else if dt.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dt.Error.Error()})
		return
	}
	session, err := sessionCreateOrGet(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dt.Error.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"data": map[string]string{"session": session.SessionKey, "eth_address": user.EthereumAddress}})
}

func LogOut(c *gin.Context) {

}

func SessionCheck(c *gin.Context) {

}

func sessionCreateOrGet(user users.Account) (session users.Session, err error) {
	sessions := user.Sessions
	// if sessions not existing, create one
	if len(sessions) == 0 {
		session = users.Session{
			AccountID: user.ID,
		}
		sessionKey := users.SessionGen()
		session.SessionKey = sessionKey
		dt := db.Create(&session)
		if dt.Error != nil {
			err = dt.Error
		}
	} else {
		session = sessions[0]
	}
	return
}

func sessionCheck(sessionKey string) bool {
	return true
}

func deleteSession(sessionKey string) bool {
	return true
}
