package init

import (
	"SE/src/database"
	"SE/src/middleware"
	"os"
)

func Init(ConfPath string) (*Config, error, *os.File) {

	// read Config file
	conf, err := initConf(ConfPath)
	if err != nil {
		return nil, err, nil
	}

	// init log
	logFile := initLog(conf.Log.Path)

	// init db
	err = database.InitDB(conf.Database)
	if err != nil {
		return nil, err, nil
	}

	// init token
	middleware.InitToken(conf.Server.Key, conf.Server.Period)

	return conf, nil, logFile
}
