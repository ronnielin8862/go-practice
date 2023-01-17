package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type (
	GlobalConfig struct {
		Psql *struct {
			Host     string `toml:"host"`
			Port     int64  `toml:"port"`
			User     string `toml:"user"`
			Password string `toml:"password"`
			DBName   string `toml:"db_name"`
		} `toml:"psql"`
	}
)

var Config = GlobalConfig{}

func InitConfig() {
	fileName := "/Users/ronnie/Library/Mobile Documents/com~apple~CloudDocs/Documents/coding/code/go-practice/config.toml"
	_, err := toml.DecodeFile(fileName, &Config)
	if err != nil {
		fmt.Printf("load config file '%s' failed, %s", fileName, err)
	}
}
