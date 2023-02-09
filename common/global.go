package common

import (
	"confuse/common/config"
	"confuse/lib/gorm"
	"confuse/lib/http"
	"confuse/lib/logger"
	"confuse/lib/redis"
	oRedis "github.com/go-redis/redis/v9"
	oGorm "gorm.io/gorm"
)

var cachePool *redis.Pool
var cacheClusterPool *redis.ClusterPool
var dbPool *gorm.Pool
var Logger *logger.Logger
var HttpServer *http.Server

func InitLogger() (err error) {
	conf := config.GetConfig().Log
	Logger, err = logger.NewLogger(conf)
	return err
}

func InitDB() (err error) {
	confs := config.GetConfig().DB
	dbPool = gorm.NewPool(Logger)

	for k, v := range confs {
		if err := dbPool.Add(k, v); err != nil {
			return err
		}
	}

	return nil
}

func InitCache() {
	confs := config.GetConfig().Cache
	cachePool = redis.NewPool()

	for k, v := range confs {
		cachePool.Add(k, v)
	}
}

func GetDB(name string) (*oGorm.DB, error) {
	return dbPool.Get(name)
}

func GetCache(name string) (*oRedis.Client, error) {
	return cachePool.Get(name)
}

func InitCacheCluster() {
	confs := config.GetConfig().CacheCluster
	cacheClusterPool = redis.NewClusterPool()

	for k, v := range confs {
		cacheClusterPool.Add(k, v)
	}
}

func GetCacheCluster(name string) (*oRedis.ClusterClient, error) {
	return cacheClusterPool.Get(name)
}

func InitHttpServer(router http.Router) {
	c := config.GetConfig().Server.Http
	HttpServer = http.NewServer(c, router, Logger)
	HttpServer.Start()
}
