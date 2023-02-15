package control

import (
	"confuse/api/logic/service"
	"confuse/common"
	"confuse/common/consts"
	ctrl "confuse/common/control"
	"confuse/common/exception"
	"confuse/lib/proto/confuse_api"
	"github.com/gin-gonic/gin"
)

var User = &userCtrl{}

type userCtrl struct{}

func (c *userCtrl) Login(ctx *gin.Context) {
	req := &confuse_api.LoginReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Account == "" || req.Password == "") {
		return
	}

	token, err := service.User.Login(req.Account, req.Password)

	if err != nil {
		ctrl.Error(ctx, err.GetCode())
		return
	}

	rsp := &confuse_api.LoginRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data:    token,
	}

	ctrl.SendRsp(ctx, rsp)
}

func (c *userCtrl) GetInfo(ctx *gin.Context) {
	authUser := service.User.GetContextUser(ctx)

	if authUser == nil {
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	common.Logger.Debugf("get user info. | trace_id: %s | user_info: %+v", common.GetTraceId(ctx), authUser)

	ctrl.SendRsp(ctx, authUser)
}

func (c *userCtrl) Logout(ctx *gin.Context) {
	authUser := service.User.GetContextUser(ctx)

	if authUser == nil {
		ctrl.Exception(ctx, exception.New(exception.CodeInternalError))
		return
	}

	service.User.Logout(authUser.Id)

	ctrl.SendRsp(ctx, authUser)
}

func (c *userCtrl) RefreshToken(ctx *gin.Context) {
	req := &confuse_api.RefreshTokenReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.RefreshToken == "") {
		return
	}

	token, err := service.User.RefreshToken(req.RefreshToken)

	if err != nil {
		ctrl.Error(ctx, err.GetCode())
		return
	}

	rsp := &confuse_api.RefreshTokenRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data:    token,
	}

	ctrl.SendRsp(ctx, rsp)
}
