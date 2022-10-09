package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"yifan/app/db"
)

type Middleware interface {
	// // Jwt 中间件
	// Jwt(ctx core.Context) (userId int64, userName string, err errno.Error)

	// // Resubmit 中间件
	// Resubmit() core.HandlerFunc

	// // DisableLog 不记录日志
	// DisableLog() core.HandlerFunc
	Jwt() gin.HandlerFunc

	TimedHandler(duration time.Duration) gin.HandlerFunc

	Timeout(timeout time.Duration) gin.HandlerFunc
	// // Token 签名验证，对登录用户的验证
	RateLimit() gin.HandlerFunc

	Recover() gin.HandlerFunc

	Cors() gin.HandlerFunc
	// // RBAC 权限验证
	// RBAC() core.HandlerFunc
	Local() gin.HandlerFunc
}
type middleware struct {
	db db.Repo
}

func New(db db.Repo) Middleware {
	return &middleware{
		db: db,
	}
}
