package httpex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_STATUS_OK             = 10000 //成功
	ERROR_STATUS_NOTAUTH_ERROR  = 10401 //未授权
	ERROR_STATUS_PWD_ERROR      = 10402 //密码不一致
	ERROR_STATUS_PARAMS_ERROR   = 10403 //参数错误
	ERROR_STATUS_DATA_NOTEXIST  = 10404 //数据不存在
	ERROR_STATUS_DATA_EXIST     = 10405 //数据已存在
	ERROR_STATUS_REQ_REJECT     = 10406 //请求被拒绝
	ERROR_STATUS_CODE_ERROR     = 10407 //验证码错误
	ERROR_STATUS_NAME_PWD_ERROR = 10408 //用户名或密码错误
	ERROR_STATUS_DATA_EXPIRE    = 10409 //数据已过期
	ERROR_STATUS_DATA_ERROR     = 10410 //数据出错
	ERROR_STATUS_SERVER_ERROR   = 10500 //服务器错误
)

type BaseApiResponse struct {
	Status int         `json:"status_code"`
	Msg    string      `json:"status_msg"`
	Result interface{} `json:"result"`
}

func DefaultApiResponse() *BaseApiResponse {
	rsp := &BaseApiResponse{http.StatusOK, "", nil}
	return rsp
}

func errMessage(statusCode int) string {
	switch statusCode {
	case ERROR_STATUS_OK:
		return "成功"
	case ERROR_STATUS_NOTAUTH_ERROR:
		return "用户未授权"
	case ERROR_STATUS_PWD_ERROR:
		return "密码错误"
	case ERROR_STATUS_PARAMS_ERROR:
		return "错误错误"
	case ERROR_STATUS_DATA_NOTEXIST:
		return "数据不存在"
	case ERROR_STATUS_DATA_EXIST:
		return "数据已存在"
	case ERROR_STATUS_REQ_REJECT:
		return "请求被拒绝"
	case ERROR_STATUS_CODE_ERROR:
		return "验证码错误"
	case ERROR_STATUS_SERVER_ERROR:
		return "服务器错误"
	}

	return "未知错误"
}

func Response(ctx *gin.Context, statusCode int, result interface{}) {
	if result == nil {
		result = struct{}{}
	}
	response := DefaultApiResponse()
	response.Status = statusCode
	response.Msg = errMessage(statusCode)
	response.Result = result
	ctx.JSON(statusCode, response)
}

func AbortWithStatus(ctx *gin.Context, statusCode int) {
	Response(ctx, statusCode, nil)
}

func ParseJSONPayload(ctx *gin.Context, payload interface{}) bool {
	if err := ctx.BindJSON(payload); err != nil {
		Response(ctx, ERROR_STATUS_PARAMS_ERROR, nil)
		return false
	}
	return true
}
