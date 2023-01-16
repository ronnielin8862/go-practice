package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type (
	GlobalConfig2 struct {
		Psql2 *struct {
			Host     string `toml:"host"`
			Port     int64  `toml:"port"`
			User     string `toml:"user"`
			Password string `toml:"password"`
			DBName   string `toml:"db_name"`
		} `toml:"psql"`
	}
)

var config2 = GlobalConfig2{}

func LoadGlobalConfig2() *GlobalConfig2 {
	fileName := "/Users/ronnie/Library/Mobile Documents/com~apple~CloudDocs/Documents/coding/code/go-practice/config.toml"
	_, err := toml.DecodeFile(fileName, &config2)
	if err != nil {
		fmt.Printf("load config file '%s' failed, %s", fileName, err)
	}
	//config2.valid = true
	return &config2
}
