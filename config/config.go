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

var config = GlobalConfig{}

func LoadGlobalConfig(filename string) (*GlobalConfig, error) {
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return nil, fmt.Errorf("load config file '%s' failed, %s", filename, err)
	}
	//config.valid = true

	return &config, nil
}
