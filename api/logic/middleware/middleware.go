package middleware

import (
	"confuse/api/util"
	"confuse/common"
	"confuse/common/consts"
	"confuse/common/control"
	"confuse/common/entity"
	"confuse/common/exception"
	"confuse/common/model"
	"confuse/lib/proto/confuse_api"
	"github.com/gin-gonic/gin"
	"strconv"
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

	uId, _ := strconv.ParseInt(userClaim.UserId, 10, 64)
	currentToken, e := model.Token.GetTokenByUserId(uId)

	if e != nil {
		common.Logger.Infof("get redis token fail. | err: %s", e)
		control.Exception(ctx, exception.New(exception.CodeInternalError))
		ctx.Abort()
		return
	}

	if currentToken != tokenString {
		common.Logger.Infof("token not equal. | err: %s", e)
		control.Exception(ctx, exception.New(exception.CodeTokenCovered))
		ctx.Abort()
		return
	}

	ctx.Set(consts.CtxValueAuth, &confuse_api.UserData{
		Id:   userClaim.UserId,
		Name: userClaim.Name,
	})

	ctx.Next()
}
