package init

import (
	"SE/src/database"
)

func Init(ConfPath string) error {

	// read Config file
	conf, err := initConf(ConfPath)
	if err != nil {
		return err
	}

	// init log
	initLog(conf.Log.Path)

	// init db
	err = database.InitDB(conf.Database.Type, conf.Database.Path)
	if err != nil {
		return err
	}

	return nil
}
