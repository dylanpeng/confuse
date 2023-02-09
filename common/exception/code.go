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
	CodeReceiveFail = iota + 201
	CodeUserAlreadyRegistered
	CodeUserAlreadyInvited
	CodeOtpSendFailed
	CodeOtpVerifyFailed
	CodeActivityExpired
)

const (
	CodeSubCategoryHaveOnlinePushCannotDeleted = iota + 301
	CodePushConfigCannotDeleted
	CodePushConfigCannotOnline
	CodePushConfigExpired
)

const (
	CodePromotionNameExist = iota + 1001
	CodePromotionCodeExist
	CodePromotionNotExist
	CodePromotionReportExpired
	CodePromotionReported
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

	CodeReceiveFail:           "failed to claim bonus",
	CodeUserAlreadyRegistered: "This user is already registered and cannot claim the reward",
	CodeUserAlreadyInvited:    "This user has been invited. You can download opay directly to receive rewards",
	CodeOtpSendFailed:         "SMS OTP sending failed",
	CodeOtpVerifyFailed:       "SMS OTP verify failed",
	CodeActivityExpired:       "The activity has expired",

	CodePromotionNameExist:     "promotion merchant name has been used",
	CodePromotionCodeExist:     "promotion code name has been used",
	CodePromotionNotExist:      "promotion merchant does not exist",
	CodePromotionReportExpired: "promotion report information has expired",
	CodePromotionReported:      "promotion information has been reported",
}
