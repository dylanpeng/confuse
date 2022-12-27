package model

import (
	"confuse/common/entity"
	"time"
)

type baseModel struct {
	DB    *baseDBModel
	Cache *baseCacheModel
}

func createModel(dbReadInstance, dbWriteInstance, cacheReadInstance, cacheWriteInstance string, cachePrefix string, cacheExpire time.Duration) *baseModel {
	return &baseModel{
		DB:    createDBModel(dbReadInstance, dbWriteInstance),
		Cache: createCacheModel(cacheReadInstance, cacheWriteInstance, cachePrefix, cacheExpire),
	}
}

func (m *baseModel) Add(e entity.IEntity) error {
	err := m.DB.Add(e)

	if err == nil {
		_ = m.Cache.Set(m.Cache.PrimaryKey(e), e)
	}

	return err
}

func (m *baseModel) Get(e entity.IEntity) (exist bool, err error) {
	if !e.PrimarySeted() {
		err = ErrPrimaryAttrEmpty
		return
	}

	exist, err = m.Cache.Get(m.Cache.PrimaryKey(e), e)

	if exist && err == nil {
		return
	}

	exist, err = m.DB.Get(e)

	if exist {
		_ = m.Cache.Set(m.Cache.PrimaryKey(e), e)
	}

	return
}

func (m *baseModel) Update(e entity.IEntity, props map[string]interface{}) error {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	err := m.DB.Update(e, props)

	if err != nil {
		return err
	}

	return m.Cache.Del(m.Cache.PrimaryKey(e))
}
