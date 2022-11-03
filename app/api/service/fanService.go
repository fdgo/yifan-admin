package service

import (
	"github.com/gin-gonic/gin"
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type FanService interface {
	ModifyFanStatus(req param.ReqModifyFanStatus) error
	QueryFanStatus(req param.ReqQueryFanStatus) (param.RespQueryFanStatus, error)
	QueryFanStatusCondition(req param.ReqQueryFanStatusCondition) (param.RespQueryFanStatusCondition, error)
	QueryFan(req param.ReqQueryFan) (param.RespQueryFan, error)
	ModifyFan(req param.ReqModifyFan) (param.RespModifyFan, error)
	ModifySaveFan(req param.ReqModifySaveFan) (param.RespModifySaveFan, error)

	QueryPrizePostion(req param.ReqQueryPrizePostion) (param.RespQueryPrizePostion, error)
	ModifyGoodsPosition(req param.ReqModifyGoodsPosition) (param.RespModifyGoodsPosition, error)

	GoodsDownLoad(c *gin.Context)
	GoodsDownLoadEmpty(c *gin.Context)
	PrizesDownLoad(c *gin.Context)
	OrderDownLoad(c *gin.Context)
	LuggageDownLoad(c *gin.Context)
	DeliverDownLoad(c *gin.Context)
	DeliverDetailDownLoad(c *gin.Context)
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
