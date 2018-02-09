package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/controller"
	"github.com/masato25/micro_ico_ethereum/app/setup"
	"github.com/masato25/micro_ico_ethereum/app/views"
	"github.com/masato25/micro_ico_ethereum/config"
	"github.com/masato25/micro_ico_ethereum/services/ethereumrpc"
	log "github.com/sirupsen/logrus"
)

var myconf config.ViperConfig

func startService() {
	if myconf.Ether.Enable {
		log.Info("ethereum service will start")
		ethereumrpc.Conn()
		ethereumrpc.SetTimeOut()
	} else {
		log.Warn("ethereum service not enable")
	}
	setup.ConnDB()
}

func main() {
	log.Info("[[[[[[[[[[[[[[[[[Micron ICO Services]]]]]]]]]]]]]]]]]")
	err := config.ReadConf()
	if err != nil {
		log.Error(err)
		os.Exit(3)
	}
	myconf = config.MyConfig()
	startService()
	r := gin.Default()
	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/assets", "./app/assets")
	r.HTMLRender = views.GetMultiRender()
	controller.Routes(r)
	r.Run(fmt.Sprintf(":%d", myconf.Web.Port)) // listen and serve on 0.0.0.0:8080
}
