package control

import (
	"confuse/common"
	"confuse/common/config"
	"confuse/common/consts"
	"confuse/common/exception"
	"confuse/lib/coder"
	opCommon "confuse/lib/proto/common"
	"github.com/gin-gonic/gin"
	"time"
)

func ErrorProto(errCode int, args ...interface{}) *opCommon.Error {
	ex := exception.New(errCode, args...)
	return &opCommon.Error{
		Code:    int32(ex.GetCode()),
		Message: ex.GetMessage(),
	}
}

func ExceptionProto(ex *exception.Exception) *opCommon.Response {
	return &opCommon.Response{Code: int32(ex.GetCode()), Message: ex.GetMessage()}
}

func Error(ctx *gin.Context, errCode int, args ...interface{}) {
	ex := exception.New(errCode, args...)
	SendRsp(ctx, ExceptionProto(ex))
	ctx.Abort()
}

func Exception(ctx *gin.Context, ex *exception.Exception) {
	SendRsp(ctx, ExceptionProto(ex))
	ctx.Abort()
}

func ParamAssert(ctx *gin.Context, req interface{}, condition bool) (ok bool) {
	if condition {
		common.Logger.Warningf("invalid parameter | req: { %s}", req)
		Error(ctx, exception.CodeInvalidParams)
		return false
	}

	return true
}

func DecodeReq(ctx *gin.Context, req interface{}) bool {
	if err := common.GetCtxCoder(ctx).DecodeRequest(ctx, req); err != nil {
		body, _ := coder.GetRequestBody(ctx)
		common.Logger.Warningf("invalid parameter | req: %s | error: %s", body, err)
		Error(ctx, exception.CodeInvalidParams)
		return false
	}

	ctx.Set(consts.CtxValueRequest, req)
	return true
}

func DecodeQuery(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBindQuery(req); err != nil {
		values := ctx.Request.URL.Query()
		common.Logger.Warningf("invalid parameter | query: %+v | error: %s", values, err)
		Error(ctx, exception.CodeInvalidParams)
		return false
	}

	ctx.Set(consts.CtxValueRequest, req)
	return true
}

func SendRsp(ctx *gin.Context, rsp interface{}) {
	if config.GetConfig().App.Debug {
		if ex, ok := rsp.(*opCommon.Response); ok && ex.Code != consts.RespCodeSuccess {
			common.Logger.Warningf("exception request | %s %s | %+v | %+v",
				ctx.Request.Method, ctx.Request.URL.RequestURI(), ctx.GetString(consts.CtxValueRequest), rsp)
		}
	}

	if err := common.GetCtxCoder(ctx).SendResponse(ctx, rsp); err != nil {
		common.Logger.Warningf("can't send http response | error: %s", err)
	}
}

func Health(ctx *gin.Context) {
	SendRsp(ctx, &opCommon.HealthRsp{
		Code:      consts.RespCodeSuccess,
		Message:   consts.RespMsgSuccess,
		Timestamp: time.Now().Unix(),
	})
}

func CommonRsp(ctx *gin.Context) {
	SendRsp(ctx, &opCommon.Response{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
	})
}
