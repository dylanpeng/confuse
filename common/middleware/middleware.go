package middleware

import (
	"confuse/common"
	"confuse/common/consts"
	"confuse/lib/coder"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func JsonCoder(ctx *gin.Context) {
	common.SetCtxCoder(ctx, coder.EncodingJson)
	ctx.Next()
}

func CheckEncoding(ctx *gin.Context) {
	common.SetCtxCoder(ctx, ctx.GetHeader(coder.EncodingHeader))
	ctx.Next()
}

func CrossDomain(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, UserId")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusOK)
	}

	ctx.Next()
}

func Trace(ctx *gin.Context) {
	traceId := ctx.GetHeader(consts.HeaderKeyTraceId)

	if traceId != "" {
		ctx.Set(consts.CtxValueTraceId, traceId)
		return
	}

	trace, exist := ctx.Get(consts.CtxValueTraceId)

	if exist {
		ctx.Set(consts.CtxValueTraceId, trace.(string))
		return
	}

	ctx.Set(consts.CtxValueTraceId, uuid.New().String())

	ctx.Next()
}
