package init

import (
	"SE/src/database"

	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title    string            `toml:"title"`
	Database database.Database `toml:"database"`
	Log      Log               `toml:"log"`
	Server   Server            `toml:"server"`
}

type Log struct {
	Path string `toml:"path"`
}

type Server struct {
	Capcity int    `toml:"capcity"`
	Quantum int    `toml:"quantum"`
	Key     string `toml:"key"`
}

func initConf(ConfPath string) (*Config, error) {
	var config Config
	fmt.Println("reading config file: ", ConfPath)
	if _, err := toml.DecodeFile(ConfPath, &config); err != nil {
		return nil, err
	}
	fmt.Println("config file read succed")
	return &config, nil
}
