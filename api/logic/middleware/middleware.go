package middleware

import (
	"confuse/api/util"
	"confuse/common"
	"confuse/common/consts"
	"confuse/common/control"
	"confuse/common/entity"
	"confuse/common/exception"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(ctx *gin.Context) {
	headToken := ctx.GetHeader(consts.HeaderKeyAuth)
	tokenString := strings.TrimPrefix(strings.TrimSpace(headToken), consts.HeaderKeyTokenPrefix)

	if tokenString == "" {
		common.Logger.Infof("token is null.")
		control.Exception(ctx, exception.New(exception.CodeTokenInvalid))
		ctx.Abort()
		return
	}

	token, e := util.JwtClient.ParseToken(tokenString, &entity.UserClaims{})

	if e != nil {
		common.Logger.Infof("Auth ParseToken fail. | token: %s | err: %s", tokenString, e)
		control.Exception(ctx, exception.New(exception.CodeTokenInvalid))
		ctx.Abort()
		return
	}

	userClaim, ok := token.Claims.(*entity.UserClaims)
	if !ok {
		common.Logger.Infof("Auth token convert fail. | claims: %s", token.Claims)
		control.Exception(ctx, exception.New(exception.CodeTokenConvertFail))
		ctx.Abort()
		return
	}

	if userClaim.IsRefresh {
		common.Logger.Infof("refresh token can't use as access token. | token: %s | claims: %s", tokenString, userClaim)
		control.Exception(ctx, exception.New(exception.CodeIsRefreshToken))
		ctx.Abort()
		return
	}

	ctx.Set("user", userClaim)

	return
}
