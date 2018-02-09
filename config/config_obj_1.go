package config

import (
	"bytes"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Database struct {
	Debug    bool
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Web struct {
	Port     int
	SaltPass string
}

type Ether struct {
	Address        string
	KeyPath        string
	DisplayMessage string
	KeyPassword    string
	RPCHost        string
	GasFee         int64
	Enable         bool
	TimeOutMS      int64
}

func (self Ether) GetKeyString() (content string, err error) {
	log.Debugf("GetKeyString - path: %s", self.KeyPath)
	var file *os.File
	file, err = os.Open(self.KeyPath)
	defer file.Close()
	if err != nil {
		return
	}
	bufc := bytes.NewBuffer(nil)
	_, err = io.Copy(bufc, file)
	if err != nil {
		return
	}
	content = string(bufc.Bytes())
	return
}
