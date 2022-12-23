package config

import (
	"confuse/lib/gorm"
	"confuse/lib/logger"
)

type Config struct {
	DB  map[string]*gorm.Config `toml:"db" json:"db"`
	Log *logger.Config          `toml:"log" json:"log"`
}

var conf *Config

func (c *Config) Init() (err error) {
	conf = c
	return
}

func GetConfig() *Config {
	return conf
}
