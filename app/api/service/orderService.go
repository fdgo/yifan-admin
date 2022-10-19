package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type OrderService interface {
	PageOfOrder(req param.ReqPageOfOrder) (param.RespPageOfOrder, error)
	PageOfOrderCondition(req param.ReqPageOfOrderCondition) (param.RespPageOfOrderCondition, error)
}

type OrderServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewOrderServiceImpl(db db.Repo, cache *cache.CacheRepo) OrderService {
	return &OrderServiceImpl{
		db:    db,
		cache: cache,
	}
}
