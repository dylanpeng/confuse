package common

import (
	"confuse/common/config"
	"confuse/lib/gorm"
	"confuse/lib/logger"
	oGorm "gorm.io/gorm"
)

var dbPool *gorm.Pool
var Logger *logger.Logger

func InitLogger() (err error) {
	conf := config.GetConfig().Log
	Logger, err = logger.NewLogger(conf)
	return err
}

func InitDB() (err error) {
	confs := config.GetConfig().DB
	dbPool = gorm.NewPool()

	for k, v := range confs {
		if err := dbPool.Add(k, v); err != nil {
			return err
		}
	}

	return nil
}

func GetDB(name string) (*oGorm.DB, error) {
	return dbPool.Get(name)
}
