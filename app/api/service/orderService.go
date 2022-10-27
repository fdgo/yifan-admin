package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type OrderService interface {
	AddRemark(req param.ReqAddRemark) error
	PageOfOrder(req param.ReqPageOfOrder) (param.RespPageOfOrder, error)
	PageOfOrderCondition(req param.ReqPageOfOrderCondition) (param.RespPageOfOrderCondition, error)
	PageOfOrderDetail(req param.ReqPageOfOrderDetail) (param.RespPageOfOrderDetail, error)
}

type OrderServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewOrderService(db db.Repo, cache *cache.CacheRepo) OrderService {
	return &OrderServiceImpl{
		db:    db,
		cache: cache,
	}
}
