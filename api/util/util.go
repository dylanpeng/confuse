package util

import (
	"confuse/api/config"
	"confuse/lib/jwt"
)

var JwtClient *jwt.JwtClient

func InitJwtClient() (err error) {
	conf := config.GetJwt()
	JwtClient, err = jwt.NewJwtClient(conf)
	return
}
