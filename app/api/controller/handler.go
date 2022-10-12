package controller

import (
	"github.com/gin-gonic/gin"
	"yifan/app/api/service"
	"yifan/app/cache"
	"yifan/app/db"
)

type Handler interface {
	XYZ() gin.HandlerFunc
	UpLoadIPs() gin.HandlerFunc
	SearchIP() gin.HandlerFunc
	AddIP() gin.HandlerFunc
	DeleteIP() gin.HandlerFunc
	QueryIP() gin.HandlerFunc
	ModifyIP() gin.HandlerFunc

	UpLoadSeries() gin.HandlerFunc
	SearchSeries() gin.HandlerFunc
	AddSeries() gin.HandlerFunc
	DeleteSeries() gin.HandlerFunc
	QuerySeries() gin.HandlerFunc
	ModifySeries() gin.HandlerFunc

	UpLoadGoods() gin.HandlerFunc
	SearchGoods() gin.HandlerFunc
	AddGoods() gin.HandlerFunc
	DeleteGoods() gin.HandlerFunc
	QueryGoods() gin.HandlerFunc
	ModifyGoods() gin.HandlerFunc

	AddBox() gin.HandlerFunc
	PageOfPosition() gin.HandlerFunc
	PageOfPositionCondition() gin.HandlerFunc
	SetNormalPrizePosition() gin.HandlerFunc
	DeleteBox() gin.HandlerFunc
	ModifyBox() gin.HandlerFunc
	ModifyBoxStatus() gin.HandlerFunc

	//
	QueryGoodsForBox() gin.HandlerFunc
	GoodsToBePrize() gin.HandlerFunc
	ModifyBoxGoods() gin.HandlerFunc
	DeleteBoxGoods() gin.HandlerFunc

	//

	AddFan() gin.HandlerFunc
	ModifyFanStatus() gin.HandlerFunc
	QueryFan() gin.HandlerFunc
	ModifyFan() gin.HandlerFunc
	ModifySaveFan() gin.HandlerFunc
	QueryPrizePostion() gin.HandlerFunc
	ModifyGoodsPosition() gin.HandlerFunc
}
type handler struct {
	userService  service.UserService
	ipService    service.IpService
	seriService  service.SeriService
	goodsService service.GoodsService
	fanService   service.FanService
	boxService   service.BoxService
}

func New(db db.Repo, cache *cache.CacheRepo) Handler {
	return &handler{
		userService:  service.NewUserService(db),
		ipService:    service.NewIpService(db, cache),
		seriService:  service.NewSeriService(db, cache),
		goodsService: service.NewgoodsService(db, cache),
		fanService:   service.NewFanService(db, cache),
		boxService:   service.NewBoxService(db, cache),
	}
}
