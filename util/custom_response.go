package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSingleResponse(ctx *gin.Context, message string, data any, code int) {
	ctx.JSON(http.StatusOK, SingleResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	})
}

func SendPagingResponse(ctx *gin.Context, message string, data []any, code int, paging Paging) {
	ctx.JSON(http.StatusOK, PagingResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Data:   data,
		Paging: paging,
	})
}
