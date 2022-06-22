package main

import (
	initMod "SE/src/init"
	"SE/src/middleware"
	Router "SE/src/router"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// init system
	conf, err, logFile := initMod.Init(ConfPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// init gin server
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.RateLimit(time.Second, int64(conf.Server.Capcity), int64(conf.Server.Quantum)))
	router.Use(middleware.Logger(logFile))

	// start server
	Router.MainRouter(router)

}
