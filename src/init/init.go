package init

import (
	"SE/src/Router"
	"SE/src/database"
	"SE/src/middleware"

	"time"

	"github.com/gin-gonic/gin"
)

func Init(ConfPath string) error {

	// read Config file
	conf, err := initConf(ConfPath)
	if err != nil {
		return err
	}

	// init log
	logFile := initLog(conf.Log.Path)

	// init db
	err = database.InitDB(conf.Database.Type, conf.Database.Path)
	if err != nil {
		return err
	}

	// init gin server
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.RateLimit(time.Second, int64(conf.Server.Capcity), int64(conf.Server.Quantum)))
	router.Use(middleware.Logger(logFile))

	Router.MainRouter(router)

	return nil
}
