package config

import (
	"confuse/lib/gorm"
	"confuse/lib/logger"
	"confuse/lib/redis"
)

type Config struct {
	Cache        map[string]*redis.Config        `toml:"cache" json:"cache"`
	CacheCluster map[string]*redis.ClusterConfig `toml:"cache_cluster" json:"cache_cluster"`
	DB           map[string]*gorm.Config         `toml:"db" json:"db"`
	Log          *logger.Config                  `toml:"log" json:"log"`
}

var conf *Config

func (c *Config) Init() (err error) {
	conf = c
	return
}

func GetConfig() *Config {
	return conf
}
