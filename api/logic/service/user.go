package service

import (
	"confuse/api/util"
	"confuse/common"
	"confuse/common/entity"
	"confuse/common/exception"
	"confuse/lib/proto/confuse_api"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var User = &userSrv{}

type userSrv struct{}

func (s *userSrv) Login(account, password string) (tokenData *confuse_api.TokenData, err *exception.Exception) {
	/*
		check pass word
	*/
	claim := &entity.UserClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "confuse-api",
			Subject:   "confuse-api-token",
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
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
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
		UserId:    "123456",
		Name:      "user_name",
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
			Id:   "123456",
			Name: "user_name",
		},
	}

	return
}
