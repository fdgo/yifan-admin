package cron

import (
	"github.com/robfig/cron"
	"yifan/app/cache"
	"yifan/app/db"
)

type CronService interface {
	Start()
	Stop()
	// AddJob(spec string, job cron.Job)
}

type cronServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
	cron  *cron.Cron
}

func NewCronService(db db.Repo, cache *cache.CacheRepo) CronService {
	return &cronServiceImpl{
		db:    db,
		cache: cache,
		cron:  cron.New(),
	}
}
