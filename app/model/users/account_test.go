package users

import (
	"testing"

	"github.com/masato25/micro_ico_ethereum/config"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAccount(t *testing.T) {
	config.ReadConf("../../../config")
	log.SetLevel(log.DebugLevel)
	Convey("test password generator", t, func() {
		log.Info(HashPassword("9999"))
	})
}
