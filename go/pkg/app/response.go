package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, SuccessResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func (r *Response) ToErrorResponse(err interface{}) {
	r.Ctx.JSON(http.StatusOK, ErrorResponse{
		Code:    1,
		Message: "error",
		Data:    err,
	})
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list":      list,
		"totalRows": totalRows,
	})
}
