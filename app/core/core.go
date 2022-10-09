package core

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"yifan/pkg/swagger"
)

type Mux interface {
	http.Handler
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return m.engine.Group(relativePath, handlers...)
}

func New() (Mux, error) {
	gin.SetMode(gin.ReleaseMode)
	mux := &mux{
		engine: gin.New(),
	}
	mux.engine.Use(cors.Default())
	mux.engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	//mux.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if swagger.SwagHandler != nil {
		mux.engine.GET("/swagger/*any", swagger.SwagHandler)
	}
	return mux, nil
}
