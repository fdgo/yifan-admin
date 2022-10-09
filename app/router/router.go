package router

import (
	"yifan/app/cache"
	"yifan/app/core"
	"yifan/app/cron"
	"yifan/app/db"
	"yifan/app/router/middleware"
	"yifan/pkg/logger"
)

type resource struct {
	mux     core.Mux
	db      db.Repo
	cache   *cache.CacheRepo
	middles middleware.Middleware
}
type Server struct {
	Mux        core.Mux
	Db         db.Repo
	Cache      *cache.CacheRepo
	CronServer cron.CronService
}

func NewHTTPServer() (*Server, error) {
	r := new(resource)

	dbRepo, err := db.New()
	if err != nil {
		logger.Fatal("new db err", err)
	}
	cacheRepo, err := cache.New()
	if err != nil {
		logger.Fatal("new cache err", err)
	}

	r.db = dbRepo
	r.cache = cacheRepo
	cronServer := cron.NewCronService(dbRepo, cacheRepo)
	cronServer.Start()

	mux, err := core.New()
	if err != nil {
		panic(err)
	}
	r.mux = mux
	r.middles = middleware.New(r.db)
	setApiRouter(r)

	s := new(Server)
	s.Db = r.db
	s.Cache = r.cache
	s.Mux = mux
	s.CronServer = cronServer

	return s, nil
}
