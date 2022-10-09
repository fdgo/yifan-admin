package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type SeriService interface {
	UpLoadSeries(req param.ReqUpLoadSeries) (param.RespUpLoadSeries, error)
	SearchSeries(req param.ReqSearchSeries) (param.RespSearchSeries, error)
	AddSeries(req param.ReqAddSeries) (uint, error)
	DeleteSeries(req param.ReqDeleteSeries) error
	QuerySeries(req param.ReqQuerySeries) (param.RespQuerySeries, error)
	ModifySeries(req param.ReqModifySeries) error
}

type SeriServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewSeriService(db db.Repo, cache *cache.CacheRepo) SeriService {
	return &SeriServiceImpl{
		db:    db,
		cache: cache,
	}
}
