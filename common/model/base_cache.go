package model

import (
	"confuse/common"
	"confuse/common/entity"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"time"
)

type baseCacheModel struct {
	readInstance  string
	writeInstance string
	prefix        string
	expire        time.Duration
}

func createCacheModel(readInstance, writeInstance string, prefix string, expire time.Duration) *baseCacheModel {
	return &baseCacheModel{
		readInstance:  readInstance,
		writeInstance: writeInstance,
		prefix:        prefix,
		expire:        expire,
	}
}

func (m *baseCacheModel) GetKey(items ...interface{}) string {
	return common.GetKey(m.prefix, items...)
}

func (m *baseCacheModel) PrimaryKey(e entity.IEntity) string {
	return m.GetKey(e.PrimaryPairs()...)
}

func (m *baseCacheModel) getCache(write bool) (*redis.Client, error) {
	if !write {
		return common.GetCache(m.readInstance)
	}

	return common.GetCache(m.writeInstance)
}

func (m *baseCacheModel) Get(key string, data interface{}) (bool, error) {
	cache, err := m.getCache(false)

	if err != nil {
		return false, err
	}

	result, err := cache.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(result), data)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *baseCacheModel) Set(key string, data interface{}, exp ...time.Duration) error {
	cache, err := m.getCache(true)

	if err != nil {
		return err
	}

	result, err := json.Marshal(data)

	if err != nil {
		return err
	}

	expire := m.expire

	if len(exp) > 0 {
		expire = exp[0]
	}

	return cache.Set(context.Background(), key, result, expire).Err()
}

func (m *baseCacheModel) Del(key string) error {
	cache, err := m.getCache(true)

	if err != nil {
		return err
	}

	return cache.Del(context.Background(), key).Err()
}
