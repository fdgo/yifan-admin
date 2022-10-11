package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"yifan/app/db"
)

type Middleware interface {
	Jwt() gin.HandlerFunc

	TimedHandler(duration time.Duration) gin.HandlerFunc

	Timeout(timeout time.Duration) gin.HandlerFunc
	// // Token 签名验证，对登录用户的验证
	RateLimit() gin.HandlerFunc

	Recover() gin.HandlerFunc

	Cors() gin.HandlerFunc
}
type middleware struct {
	db db.Repo
}

func New(db db.Repo) Middleware {
	return &middleware{
		db: db,
	}
}
