package common_helper

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/model/users"
	"github.com/masato25/micro_ico_ethereum/app/setup"
	log "github.com/sirupsen/logrus"
)

type SessionObj struct {
	SessionKey string `json:"session_key"`
	UserName   string `json:"user_name"`
}

func SessionCheck(c *gin.Context) (input *SessionObj, err error) {
	sessionInfo := c.GetHeader("session")
	if sessionInfo == "" {
		getCookie, _ := c.Cookie("auth_session")
		sessionInfo = getCookie
	}
	log.Debugf("[SessionCheck] sessionInfo: %s", sessionInfo)
	err = json.Unmarshal([]byte(sessionInfo), &input)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func GetUser(mysession SessionObj) (user users.Account, err error) {
	db := setup.GetConn()
	user = users.Account{UserName: mysession.UserName}
	var userTmp []users.Account
	db.Where(user).Find(&userTmp)
	if len(userTmp) == 0 {
		err = errors.New("user not found")
		return
	}
	user = userTmp[0]
	return
}

func GetUserBySession(c *gin.Context) (user users.Account, err error) {
	var input *SessionObj
	input, err = SessionCheck(c)
	if err != nil {
		log.Error(err)
		return
	}
	user, err = GetUser(*input)
	return
}

func GetUserByJsonSessionString(sessionHeader string) (user users.Account, err error) {
	var input *SessionObj
	err = json.Unmarshal([]byte(sessionHeader), &input)
	if err != nil {
		log.Error(err)
		return
	}
	user, err = GetUser(*input)
	return
}
