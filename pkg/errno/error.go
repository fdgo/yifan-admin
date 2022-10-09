package errno

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	StatusOK           = 200
	PostParamError     = 400
	StatusUnauthorized = 401
	InternalError      = 500
	TooManyRequests    = 1000
)

func Text(code int) string {
	return enUSText[code]
}

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	Success     = NewError(http.StatusOK, 200, "success")
	ServerError = NewError(http.StatusInternalServerError, 500, "系统异常，请稍后重试!")
	NotFound    = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}
