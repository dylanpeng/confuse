package control

import (
	"confuse/common"
	"confuse/common/consts"
	ctrl "confuse/common/control"
	"confuse/lib/proto/confuse_api"
	"github.com/gin-gonic/gin"
)

var Home = &homeCtrl{}

type homeCtrl struct{}

func (c *homeCtrl) HomePage(ctx *gin.Context) {
	req := &confuse_api.HomeReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.UserId == 0 || req.UserName == "") {
		return
	}

	rsp := &confuse_api.HomeRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &confuse_api.HomeData{
			UserId:   req.UserId,
			UserName: req.UserName,
		},
	}

	common.Logger.Debugf("traceId: %s", common.GetTraceId(ctx))

	ctrl.SendRsp(ctx, rsp)
}
