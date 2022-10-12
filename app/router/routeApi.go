package router

import (
	"yifan/app/api/controller"
)

func setApiRouter(r *resource) { //Use(cors.New(config))
	handler := controller.New(r.db, r.cache)
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//config.AllowHeaders = append(config.AllowHeaders, "token")
	ip := r.mux.Group("/v1/ip")
	{
		ip.POST("/upload", handler.UpLoadIPs())
		ip.POST("/search", handler.SearchIP())
		ip.POST("/create", handler.AddIP())
		ip.POST("/delete", handler.DeleteIP())
		ip.POST("/query", handler.QueryIP())
		ip.POST("/modify", handler.ModifyIP())
	}
	series := r.mux.Group("/v1/series").
		Use(r.middles.Cors())
	{
		series.POST("/upload", handler.UpLoadSeries())
		series.POST("/search", handler.SearchSeries())
		series.POST("/create", handler.AddSeries())
		series.POST("/delete", handler.DeleteSeries())
		series.POST("/query", handler.QuerySeries())
		series.POST("/modify", handler.ModifySeries())
	}
	goods := r.mux.Group("/v1/goods").
		Use(r.middles.Cors())
	{
		goods.POST("/upload", handler.UpLoadGoods())
		goods.POST("/search", handler.SearchGoods())
		goods.POST("/create", handler.AddGoods())
		goods.POST("/delete", handler.DeleteGoods())
		goods.POST("/query", handler.QueryGoods())
		goods.POST("/modify", handler.ModifyGoods())
	}
	box := r.mux.Group("/v1/box").
		Use(r.middles.Cors())
	{
		box.POST("/create", handler.AddBox())
		box.POST("/pagePosition", handler.PageOfPosition())
		box.POST("/pagePosition/condition", handler.PageOfPositionCondition())
		box.POST("/setnPosition", handler.SetNormalPrizePosition())
		box.POST("/delete", handler.DeleteBox())
		box.POST("/modify/status", handler.ModifyBoxStatus())
		box.POST("/goods/query", handler.QueryGoodsForBox())
		box.POST("/goods/toBePrize", handler.GoodsToBePrize())
		box.POST("/goods/modify", handler.ModifyBoxGoods())
		box.POST("/goods/delete", handler.DeleteBoxGoods())
	}
	fan := r.mux.Group("/v1/fan").
		Use(r.middles.Cors())
	{
		fan.POST("/modify/status", handler.ModifyFanStatus())
		fan.POST("/query", handler.QueryFan())
		fan.POST("/modify", handler.ModifyFan())
		fan.POST("/modify/save", handler.ModifySaveFan())
		fan.POST("/queryPostion", handler.QueryPrizePostion())
		fan.POST("/modifyPosition", handler.ModifyGoodsPosition())
	}
	//memoryStore := persist.NewMemoryStore(1 * time.Minute)
	//power.GET("/123", cache.CacheByRequestURI(memoryStore, 3600*time.Second),
	//	handler.GetUser(),
	//)
}
