package exception

const CodeRetSuccess = 200

const (
	CodeInternalError = iota + 1
	CodeQueryFailed
	CodeUnableConnect
	CodeForbidden
	CodeUnauthorized
	CodeNoPermission
)

const (
	CodeInvalidParams = iota + 101
	CodeConvertFailed
	CodeDataNotExist
	CodeDataAlreadyExist
	CodeDataCantSet
	CodeCallGrpcFailed
	CodeCallRestFailed
	CodeNoRoute
	CodeOperateTooFast
	CodeCallDubboFailed
	CodeFileUploadFailed
	CodeFileReadFailed
)

const (
	CodeTokenInvalid = iota + 201
	CodeIsRefreshToken
	CodeTokenConvertFail
	CodeTokenCovered
)

var Desces = map[int]string{
	CodeRetSuccess:    "success",
	CodeInternalError: "server internal error",
	CodeQueryFailed:   "data query failed",
	CodeUnableConnect: "unable to connect to server",
	CodeForbidden:     "access denied",
	CodeUnauthorized:  "unauthorized",
	CodeNoPermission:  "no permission",

	CodeInvalidParams:    "invalid parameter",
	CodeConvertFailed:    "convert data failed",
	CodeDataNotExist:     "data not exist",
	CodeDataAlreadyExist: "data already exist",
	CodeDataCantSet:      "set data failed",
	CodeCallGrpcFailed:   "call GRPC service failed",
	CodeCallRestFailed:   "call REST API service failed",
	CodeNoRoute:          "user did't register a available route",
	CodeOperateTooFast:   "operate too fast, please try again later",
	CodeCallDubboFailed:  "call DUBBO service failed",
	CodeFileUploadFailed: "upload file failed",
	CodeFileReadFailed:   "read file failed",

	CodeTokenInvalid:     "token is invalid",
	CodeIsRefreshToken:   "token is refresh token, can use as normal token",
	CodeTokenConvertFail: "token convert fail",
	CodeTokenCovered:     "token had been covered",
}
