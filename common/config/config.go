package config

import (
	"confuse/lib/coder"
	"confuse/lib/gorm"
	"confuse/lib/http"
	"confuse/lib/logger"
	"confuse/lib/net2"
	"confuse/lib/redis"
	"fmt"
)

var conf *Config

type Config struct {
	App          *App                            `toml:"app" json:"app"`
	Cache        map[string]*redis.Config        `toml:"cache" json:"cache"`
	CacheCluster map[string]*redis.ClusterConfig `toml:"cache_cluster" json:"cache_cluster"`
	DB           map[string]*gorm.Config         `toml:"db" json:"db"`
	Log          *logger.Config                  `toml:"log" json:"log"`
	Server       *Server                         `toml:"server" json:"server"`
}

func (c *Config) Init() (err error) {
	// set coder
	if c.App.HttpCode == coder.EncodingProtobuf {
		c.App.HttpCoder = coder.ProtoCoder
	} else {
		c.App.HttpCoder = coder.JsonCoder
	}

	if c.App.TcpCode == coder.EncodingProtobuf {
		c.App.TcpCoder = coder.ProtoCoder
	} else {
		c.App.TcpCoder = coder.JsonCoder
	}

	if err = c.SetHost(); err != nil {
		return
	}

	conf = c
	return
}

func (c *Config) SetHost() (err error) {
	if c.Server != nil && c.Server.BindInterface != "" {
		interfaceIp, err := net2.GetInterfaceIp(c.Server.BindInterface)

		if err != nil {
			return err
		}

		ipStr := interfaceIp.String()

		if ipStr != "" {
			if c.Server.Http != nil {
				c.Server.Http.Host = ipStr
			}

			c.App.Host = ipStr
		}
	}

	if c.App.Host == "" {
		localIp, err := net2.GetLocalIp()

		if err != nil {
			return err
		}

		localIpStr := localIp.String()

		if localIpStr != "" {
			c.App.Host = localIpStr
		}
	}

	return
}

type App struct {
	Name      string       `toml:"name" json:"name"`
	Project   string       `toml:"project" json:"project"`
	Env       string       `toml:"env" json:"env"`
	Debug     bool         `toml:"debug" json:"debug"`
	Host      string       `toml:"host" json:"host"`
	HttpCode  string       `toml:"http_code" json:"http_code"`
	TcpCode   string       `toml:"tcp_code" json:"tcp_code"`
	HttpCoder coder.ICoder `toml:"-" json:"-"`
	TcpCoder  coder.ICoder `toml:"-" json:"-"`
}

func (c *App) GetKeyPrefix() string {
	return fmt.Sprintf("%s:%s:", c.Project, c.Env)
}

type Server struct {
	BindInterface string       `toml:"bind_interface" json:"bind_interface"`
	Http          *http.Config `toml:"http" json:"http"`
}

func Default() *Config {
	return &Config{
		Server: &Server{
			Http: http.DefaultConfig(),
		},
		Log: logger.DefaultConfig(),
	}
}

func GetConfig() *Config {
	return conf
}
