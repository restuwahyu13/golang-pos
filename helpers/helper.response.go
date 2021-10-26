package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/restuwahyu13/golang-pos/schemas"
)

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {

	jsonResponse := schemas.SchemaResponses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ErrorResponse(ctx *gin.Context, Error interface{}) {
	err := schemas.SchemaErrorResponse{
		StatusCode: http.StatusBadRequest,
		Error:      Error,
	}

	ctx.AbortWithStatusJSON(err.StatusCode, err)
}
