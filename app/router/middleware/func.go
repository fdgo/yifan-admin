package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
	"yifan/configs"
	"yifan/pkg/errno"
	"yifan/pkg/jwtex"
)

const _MaxBurstSize = 102400

func (h *middleware) Local() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Local.................")
		if context.ClientIP() != "127.0.0.1" {
			context.AbortWithError(http.StatusUnauthorized, errors.New(errno.Text(errno.StatusUnauthorized)))
		}
		context.Next()
	}
}

func (h *middleware) Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Recover.................")
		defer func() {
			if err := recover(); err != nil {
				var Err *errno.Error
				if e, ok := err.(*errno.Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = errno.OtherError(e.Error())
				} else {
					Err = errno.ServerError
				}
				context.JSON(Err.StatusCode, Err)
				return
			}
		}()
		context.Next()
	}
}

func (h *middleware) RateLimit() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("RateLimit.................")
		limiter := rate.NewLimiter(rate.Every(time.Second*1), _MaxBurstSize)
		if !limiter.Allow() {
			context.AbortWithError(http.StatusTooManyRequests, errors.New(errno.Text(errno.TooManyRequests)))
			return
		}
		context.Next()
	}
}
func (h *middleware) Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Timeout.................")
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
func (h *middleware) Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := jwtex.JwtToken{
			SigningKey: []byte(configs.GetConfig().Jwt.JwtSecret),
		}
		sub, err := jwt.Decode(c.Request, c.Writer)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		userId, _ := strconv.Atoi(sub.Id)
		c.Set("user_id", userId)
		c.Next()
	}
}

func (h *middleware) TimedHandler(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("TimedHandler.................")
		ctx := c.Request.Context()
		type responseData struct {
			status int
			body   map[string]interface{}
		}
		doneChan := make(chan responseData)
		go func() {
			time.Sleep(duration)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()
		select {
		case <-ctx.Done():
			return
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}
func (h *middleware) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization,Authorization, Token, browserVersion")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("content-type", "application/json")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
