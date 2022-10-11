package response

import (
	"net/http"
	"yifan/pkg/errno"

	"github.com/gin-gonic/gin"
)

type responseSucess struct {
	Code int         ` json:"code" example:"200" `
	Msg  string      ` json:"msg" example:"success" `
	Data interface{} ` json:"data" `
}
type responseFailure struct {
	Code int         ` json:"code" example:"400" `
	Msg  string      ` json:"msg" example:"failure" `
	Data interface{} ` json:"data" `
}

func ResposeSuccess(data interface{}, context *gin.Context) {
	context.JSON(http.StatusOK, responseSucess{
		errno.StatusOK,
		errno.Text(errno.StatusOK),
		data,
	})
}

func AbortWithBadRequestWithError(err error, context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusOK, responseFailure{
		errno.PostParamError,
		err.Error(),
		nil,
	})
}

func AbortWithBadRequestWithData(err error, msg interface{}, context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusOK, responseFailure{
		errno.PostParamError,
		err.Error(),
		msg,
	})
}
