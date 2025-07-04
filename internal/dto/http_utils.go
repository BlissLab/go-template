package dto

import "github.com/gin-gonic/gin"

var (
	Http HttpUtils
)

type HttpUtils struct{}

func (HttpUtils) Ok(ctx *gin.Context, data any) {
	response := Response{
		Code:    SUCCESS,
		Message: SUCCESS_MSG,
		Data:    data,
	}
	if response.Data == nil {
		response.Data = struct{}{}
	}
	ctx.JSON(200, response)
}

func (HttpUtils) Fail(ctx *gin.Context, data any) {
	response := Response{
		Code:    FAILED,
		Message: FAILED_MSG,
		Data:    data,
	}
	if response.Data == nil {
		response.Data = struct{}{}
	}
	ctx.JSON(200, response)
}

func (HttpUtils) FailWith(ctx *gin.Context, code int, msg string, data any) {
	response := Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	if response.Data == nil {
		response.Data = struct{}{}
	}
	ctx.JSON(200, response)
}
