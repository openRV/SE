// @Title init.go
// @Description 配置、数据库、log初始化相关的函数
// @Author 杜沛然 ${DATE} ${TIME}

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
