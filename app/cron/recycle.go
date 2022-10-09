package cron

import (
	"fmt"
	"time"
	"yifan/app/cache"
	"yifan/app/db"
	"yifan/pkg/logger"
)

const (
	poweringFactor = 10
	precision      = 5
)

func NewRecycleService(db db.Repo, cache *cache.CacheRepo) recycle {
	r := make(map[int8]float64)
	return recycle{
		db:           db,
		cacheRepo:    cache,
		maxRandom:    poweringFactor*2 + 1,
		coefficients: r,
	}
}

type recycle struct {
	db           db.Repo
	cacheRepo    *cache.CacheRepo
	maxRandom    int64
	coefficients map[int8]float64
}

func (r recycle) Run() {
	now := time.Now().UTC()
	tTimeStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0)

	end, _ := time.Parse("2006-01-02 15:04:05", tTimeStr)
	start := end.Add(-1 * time.Hour)

	go r.doTask(start, end)
}

func (r recycle) doTask(start, end time.Time) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("doTask error:", err)
		}
	}()
}
