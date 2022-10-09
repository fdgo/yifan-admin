package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type FanService interface {
	AddFan(req param.ReqAddFan) (uint, error)
	ModifyFanStatus(req param.ReqModifyFanStatus) error
	QueryFan(req param.ReqQueryFan) (param.RespQueryFan, error)
	ModifyFan(req param.ReqModifyFan) (param.RespModifyFan, error)
	ModifySaveFan(req param.ReqModifySaveFan) (param.RespModifySaveFan, error)

	QueryPrizePostion(req param.ReqQueryPrizePostion) (param.RespQueryPrizePostion, error)
	ModifyGoodsPosition(req param.ReqModifyGoodsPosition) (param.RespModifyGoodsPosition, error)

	Buy(req param.ReqBuy) (param.RespBuy, error)
	BuySure(req param.ReqBuySure) (param.RespBuySures, error)
	BuyQuery(req param.ReqBuyQuery) (param.RespBuyQuerys, error)
}

type FanServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewFanService(db db.Repo, cache *cache.CacheRepo) FanService {
	return &FanServiceImpl{
		db:    db,
		cache: cache,
	}
}
