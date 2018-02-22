package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/controller/common_helper"
	"github.com/masato25/micro_ico_ethereum/app/model/users"
	log "github.com/sirupsen/logrus"
)

func Index(c *gin.Context) {
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "common_index", gin.H{"login": lg})
	return
}

func Register(c *gin.Context) {
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "user_register", gin.H{"login": lg})
	return
}

func LogIn(c *gin.Context) {
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "user_login", gin.H{"login": lg})
	return
}

func LogOut(c *gin.Context) {
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "user_logout", gin.H{"login": lg})
	return
}

type UserRegFeilds struct {
	UserName      string `form:"username" json:"username" binding:"required"`
	Passwd        string `form:"password" json:"password" binding:"required"`
	EthAddress    string `form:"eth_address" json:"eth_address" binding:"required"`
	EthPrivateKey string `form:"eth_private_key" json:"eth_private_key" binding:"required"`
}

func RegisterJSON(c *gin.Context) {
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

func LogInJSON(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"data": map[string]string{"username": user.UserName, "session": session.SessionKey, "eth_address": user.EthereumAddress}})
	return
}

func SessionCheck(c *gin.Context) {
	user, err := common_helper.GetUserBySession(c)
	if err != nil || user.ID == 0 {
		if err != nil {
			log.Error(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": user, "status": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user, "status": true})
	return
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
