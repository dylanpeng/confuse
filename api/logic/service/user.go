package service

import (
	"confuse/api/config"
	"confuse/api/util"
	"confuse/common"
	"confuse/common/consts"
	"confuse/common/entity"
	"confuse/common/exception"
	"confuse/common/model"
	"confuse/lib/proto/confuse_api"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

var User = &userSrv{}

type userSrv struct{}

func (s *userSrv) Login(account, password string) (tokenData *confuse_api.TokenData, err *exception.Exception) {
	/*
		check pass word
	*/
	expireDay, refreshExpireDay := config.GetConfig().Jwt.ExpireTime, config.GetConfig().Jwt.RefreshExpireTime

	if expireDay <= 0 {
		expireDay = 3
	}

	if refreshExpireDay <= 0 {
		refreshExpireDay = 7
	}

	claim := &entity.UserClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "confuse-api",
			Subject:   "confuse-api-token",
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, expireDay)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
		UserId:    "123456",
		Name:      "user_name",
		IsRefresh: false,
	}

	refreshClaim := &entity.UserClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "confuse-api",
			Subject:   "confuse-api-token",
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, refreshExpireDay)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
		UserId:    claim.UserId,
		Name:      claim.Name,
		IsRefresh: true,
	}

	tokenString, e := util.JwtClient.GenerateToken(claim)

	if e != nil {
		common.Logger.Warningf("userSrv.Login Generate token fail. | claim: %s | err: %s", claim, e)
		err = exception.New(exception.CodeInternalError)
		return
	}

	refreshTokenString, e := util.JwtClient.GenerateToken(refreshClaim)

	if e != nil {
		common.Logger.Warningf("userSrv.Login Generate refresh token fail. | claim: %s | err: %s", claim, e)
		err = exception.New(exception.CodeInternalError)
		return
	}

	tokenData = &confuse_api.TokenData{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		User: &confuse_api.UserData{
			Id:   claim.UserId,
			Name: claim.Name,
		},
	}

	uId, _ := strconv.ParseInt(claim.UserId, 10, 64)
	_ = model.Token.SetTokenByUserId(uId, tokenString, time.Duration(expireDay)*time.Hour*24)
	_ = model.Token.SetRefreshTokenByUserId(uId, refreshTokenString, time.Duration(refreshExpireDay)*time.Hour*24)

	common.Logger.Infof("user_id: %d | token: %s | refresh_token: %s", uId, tokenString, refreshTokenString)

	return
}

func (s *userSrv) GetContextUser(ctx *gin.Context) (result *confuse_api.UserData) {
	value, exist := ctx.Get(consts.CtxValueAuth)

	if !exist {
		return nil
	}

	var ok bool
	result, ok = value.(*confuse_api.UserData)

	if !ok {
		return nil
	}

	return
}

func (s *userSrv) Logout(userId string) {
	id, _ := strconv.ParseInt(userId, 10, 64)
	_ = model.Token.DelTokenByUserId(id)
	_ = model.Token.DelRefreshTokenByUserId(id)
}

func (s *userSrv) RefreshToken(refreshToken string) (tokenData *confuse_api.TokenData, err *exception.Exception) {
	tokenString := strings.TrimPrefix(strings.TrimSpace(refreshToken), consts.HeaderKeyTokenPrefix)

	if tokenString == "" {
		common.Logger.Infof("token is null.")
		err = exception.New(exception.CodeTokenInvalid)
		return
	}

	token, e := util.JwtClient.ParseToken(tokenString, &entity.UserClaims{})

	if e != nil {
		common.Logger.Infof("Auth ParseToken fail. | token: %s | err: %s", tokenString, e)
		err = exception.New(exception.CodeTokenInvalid)
		return
	}

	userClaim, ok := token.Claims.(*entity.UserClaims)
	if !ok {
		common.Logger.Infof("Auth token convert fail. | claims: %s", token.Claims)
		err = exception.New(exception.CodeTokenConvertFail)
		return
	}

	if !userClaim.IsRefresh {
		common.Logger.Infof("access token can't use as refresh token. | token: %s | claims: %s", tokenString, userClaim)
		err = exception.New(exception.CodeIsRefreshToken)
		return
	}

	uId, _ := strconv.ParseInt(userClaim.UserId, 10, 64)
	cacheRefresh, e := model.Token.GetRefreshTokenByUserId(uId)

	if e != nil {
		common.Logger.Infof("get redis token fail. | err: %s", e)
		err = exception.New(exception.CodeInternalError)
		return
	}

	if cacheRefresh != refreshToken {
		common.Logger.Infof("refresh not equal. cache: %s | refreshToken: %s", cacheRefresh, refreshToken)
		err = exception.New(exception.CodeTokenCovered)
		return
	}

	tokenData, err = s.Login("", "")

	return
}
