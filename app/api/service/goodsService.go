package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type GoodsService interface {
	UpLoadGoods() error
	SearchGoods(req param.ReqSearchGoods) (param.RespSearchGoods, error)
	AddGoods(req param.ReqAddGoods) (uint, error)
	DeleteGoods(req param.ReqDeleteGoods) error
	QueryGoods(req param.ReqQueryGoods) (param.RespQueryGoods, error)
	ModifyGoods(req param.ReqModifyGoods) error
}

type GoodsServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewgoodsService(db db.Repo, cache *cache.CacheRepo) GoodsService {
	return &GoodsServiceImpl{
		db:    db,
		cache: cache,
	}
}
