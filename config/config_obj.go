package config

import (
	"strconv"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	Database Database
	Web      Web
	Ether    Ether
}

func (self ViperConfig) SetAll() ViperConfig {
	self = self.SetDatabase()
	self = self.SetWeb()
	self = self.SetEther()
	return self
}

func (self ViperConfig) SetDatabase() ViperConfig {
	dbc := viper.GetStringMap("database")
	self.Database = Database{
		Debug:    dbc["debug"].(bool),
		Host:     dbc["host"].(string),
		Port:     dbc["port"].(int),
		User:     dbc["user"].(string),
		Password: dbc["password"].(string),
		DBName:   dbc["dbname"].(string),
	}
	return self
}

func (self ViperConfig) SetWeb() ViperConfig {
	webc := viper.GetStringMap("web")
	self.Web = Web{
		Port:     webc["port"].(int),
		SaltPass: webc["saltpass"].(string),
	}
	return self
}

func (self ViperConfig) SetEther() ViperConfig {
	ethc := viper.GetStringMapString("ether")
	gasfee, _ := strconv.Atoi(ethc["gasfee"])
	timeOutMs, _ := strconv.Atoi(ethc["time_out_ms"])
	self.Ether = Ether{
		Enable:         ethc["enable"] == "true",
		Address:        ethc["address"],
		KeyPath:        ethc["keyjsonpath"],
		DisplayMessage: ethc["displaymessage"],
		KeyPassword:    ethc["keypassword"],
		RPCHost:        ethc["ethereum_rpc_host"],
		TimeOutMS:      int64(timeOutMs),
		GasFee:         int64(gasfee),
	}
	return self
}
