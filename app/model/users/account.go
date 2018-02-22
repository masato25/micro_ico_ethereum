package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"github.com/masato25/micro_ico_ethereum/config"
)

type Account struct {
	gorm.Model
	UserName                 string `gorm:"type:varchar(100);not null;unique_index`
	Password                 string `json:"-"`
	EthereumAddress          string
	EthereumPrivateKeyEncode string    `json:"-"`
	EthereumJsonKeyStore     string    `json:"-"`
	Sessions                 []Session `json:"-"`
}

func (self Account) AuthPassword(passstring string) bool {
	return self.CheckPasswordHash(passstring)
}

func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password should not be empty!")
	}
	passSalted := fmt.Sprintf("%s@%s", password, config.MyConfig().Web.SaltPass)
	h := md5.New()
	io.WriteString(h, passSalted)
	passwordHash := fmt.Sprintf("%x", h.Sum(nil))
	return passwordHash, nil
}

func (self Account) CheckPasswordHash(password string) bool {
	passhash, err := HashPassword(password)
	if err != nil {
		log.Error(err)
		return false
	}
	return self.Password == passhash
}
