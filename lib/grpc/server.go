package grpc

import (
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type Config struct {
	Host string `toml:"host" json:"host"`
	Port int    `toml:"port" json:"port"`
}

func (c *Config) GetAddr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

type Router interface {
	RegGrpcService(server *grpc.Server)
}
