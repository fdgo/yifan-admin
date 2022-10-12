package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type BoxService interface {
	AddBox(req param.ReqAddBox) (param.RespAddBox, error)
	PageOfPosition(req param.ReqPageOfPosition) (param.RespPageOfPosition, error)
	PageOfPositionCondition(req param.ReqPageOfPositionCondition) (param.RespPageOfPositionCondition, error)
	SetNormalPrizePosition(req param.ReqSetNormalPrizePosition) error
	DeleteBox(req param.ReqDeleteBox) error
	ModifyBoxStatus(req param.ReqModifyBoxStatus) error

	QueryGoodsForBox(req param.ReqQueryGoodsForBox) (param.RespQueryGoodsForBox, error)
	GoodsToBePrize(req param.ReqGoodsToBePrize) error
	ModifyBoxGoods(req param.ReqModifyBoxGoods) error
	DeleteBoxGoods(req param.ReqDeleteBoxGoods) error
}

type BoxServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewBoxService(db db.Repo, cache *cache.CacheRepo) BoxService {
	return &BoxServiceImpl{
		db:    db,
		cache: cache,
	}
}
