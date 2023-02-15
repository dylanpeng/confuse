package model

import (
	"confuse/common"
	"time"
)

var Token = &tokenModel{
	baseCacheModel: createCacheModel(
		"main-slave",
		"main-master",
		"token",
		time.Minute*10,
	),
}

type tokenModel struct {
	*baseCacheModel
}

func (m *tokenModel) GetTokenKey(userId int64) string {
	return common.GetKey(m.prefix, "id", userId)
}

func (m *tokenModel) GetRefreshTokenKey(userId int64) string {
	return common.GetKey(m.prefix, "refresh", "id", userId)
}

func (m *tokenModel) GetTokenByUserId(userId int64) (result string, err error) {
	key := m.GetTokenKey(userId)

	var exist bool
	exist, err = m.Get(key, &result)

	if err != nil {
		return "", err
	}

	if !exist {
		return "", nil
	}

	return
}

func (m *tokenModel) SetTokenByUserId(userId int64, token string, expire time.Duration) (err error) {
	key := m.GetTokenKey(userId)
	err = m.Set(key, token, expire)

	if err != nil {
		return
	}

	return
}

func (m *tokenModel) DelTokenByUserId(userId int64) (err error) {
	key := m.GetTokenKey(userId)
	err = m.Del(key)

	return
}

func (m *tokenModel) GetRefreshTokenByUserId(userId int64) (result string, err error) {
	key := m.GetRefreshTokenKey(userId)

	var exist bool
	exist, err = m.Get(key, &result)

	if err != nil {
		return "", err
	}

	if !exist {
		return "", nil
	}

	return
}

func (m *tokenModel) SetRefreshTokenByUserId(userId int64, token string, expire time.Duration) (err error) {
	key := m.GetRefreshTokenKey(userId)
	err = m.Set(key, token, expire)

	if err != nil {
		return
	}

	return
}

func (m *tokenModel) DelRefreshTokenByUserId(userId int64) (err error) {
	key := m.GetRefreshTokenKey(userId)
	err = m.Del(key)

	return
}
