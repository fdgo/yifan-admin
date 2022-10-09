package main

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"yifan/app/router"
	"yifan/configs"
	_ "yifan/docs"
	"yifan/pkg/logger"
	"yifan/pkg/shutdown"
	_ "yifan/pkg/swagger"
)

// @title wcd API
// @version 1.0.0
// @description  wcd api. developer Comply with this document.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host 192.168.37.104:8899
// @Schemes http
// @BasePath /
func main() {
	err := configs.LoadDefaultConfig()
	if err != nil {
		panic(err)
	}
	logger.Init(configs.GetConfig().LogConfig.Name,
		configs.GetConfig().LogConfig.FileSize,
		configs.GetConfig().LogConfig.MaxBackups,
		configs.GetConfig().LogConfig.MaxAge)
	defer logger.Close()
	logger.SetLevel(configs.GetConfig().LogConfig.Level)

	s, err := router.NewHTTPServer()
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(int(configs.GetConfig().Server.Port)),
		Handler: s.Mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("http server startup err,", err)
		}
	}()
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				logger.Error("server shutdown err,", err)
			}
		},
		func() {
			if s.Db != nil {
				s.Db.Close()
			}
		},
		func() {
			if s.Cache != nil {
				s.Cache.Close()
			}
		},
		func() {
			if s.CronServer != nil {
				s.CronServer.Stop()
			}
		},
	)
}

//Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTEyMDA3NDcsImp0aSI6IjE0MzIzNzYxNDgzNjEyMTYifQ._mg2MoTwj8Lnp8mjaZ-ea54A2tctX_DcnG0Epn6vDJk
//1. IP  2.系列 3款式  4商品（带包装状态）   5.箱子  6.
