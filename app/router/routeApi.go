package router

import (
	"yifan/app/api/controller"
)

func setApiRouter(r *resource) { //Use(cors.New(config))
	handler := controller.New(r.db, r.cache)
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//config.AllowHeaders = append(config.AllowHeaders, "token")
	config := r.mux.Group("/v1/admin/config")
	{
		config.POST("", handler.SetGlobalConfig())
	}
	user := r.mux.Group("/v1/admin/user")
	{
		user.POST("/list", handler.UserList())
		user.POST("/list/condition", handler.UserListCondition())
	}
	ip := r.mux.Group("/v1/admin/ip")
	{
		ip.POST("/upload", handler.UpLoadIPs())
		ip.POST("/search", handler.SearchIP())
		ip.POST("/create", handler.AddIP())
		ip.POST("/delete", handler.DeleteIP())
		ip.POST("/query", handler.QueryIP())
		ip.POST("/modify", handler.ModifyIP())
	}
	series := r.mux.Group("/v1/admin/series").
		Use(r.middles.Cors())
	{
		series.POST("/upload", handler.UpLoadSeries())
		series.POST("/search", handler.SearchSeries())
		series.POST("/create", handler.AddSeries())
		series.POST("/delete", handler.DeleteSeries())
		series.POST("/query", handler.QuerySeries())
		series.POST("/modify", handler.ModifySeries())
	}
	goods := r.mux.Group("/v1/admin/goods").
		Use(r.middles.Cors())
	{
		goods.POST("/upload", handler.UpLoadGoods())
		goods.POST("/search", handler.SearchGoods())
		goods.POST("/create", handler.AddGoods())
		goods.POST("/delete", handler.DeleteGoods())
		goods.POST("/query", handler.QueryGoods())
		goods.POST("/modify", handler.ModifyGoods())
	}
	box := r.mux.Group("/v1/admin/box").
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
	fan := r.mux.Group("/v1/admin/fan").
		Use(r.middles.Cors())
	{
		fan.POST("/modify/status", handler.ModifyFanStatus())
		fan.POST("/query/status", handler.QueryFanStatus())
		fan.POST("/query/status/condition", handler.QueryFanStatusCondition())
		fan.POST("/query", handler.QueryFan())
		fan.POST("/modify", handler.ModifyFan())
		fan.POST("/modify/save", handler.ModifySaveFan())
		fan.POST("/queryPostion", handler.QueryPrizePostion())
		fan.POST("/modifyPosition", handler.ModifyGoodsPosition())
	}
	order := r.mux.Group("/v1/admin/order").Use(r.middles.Cors())
	{
		order.POST("/addRemark", handler.AddRemark())
		order.POST("/pageOrder", handler.PageOfOrder())
		order.POST("/pageOrder/condition", handler.PageOfOrderCondition())
		order.POST("/pageOrder/detail", handler.PageOfOrderDetail())

	}
	adver := r.mux.Group("/v1/admin/adver").Use(r.middles.Cors())
	{
		adver.POST("/banner/query", handler.GetBannerPic())
		adver.POST("/banner/create", handler.SetBannerPic())
		adver.POST("/banner/isShow", handler.ShowOrHideBanner())
		adver.POST("/banner/delete", handler.DelBannerPic())
		adver.POST("/activeByMan", handler.ActiveByMan())
		adver.POST("/singleClick", handler.SingleClick())

		adver.POST("/tab/second/create", handler.AddSecondTab())
		adver.POST("/tab/second/son/create", handler.AddSecondTabSon())

		adver.POST("/tab/second/query", handler.QuerySecondTab())
		adver.POST("/tab/second/son/query", handler.QuerySecondSonTab())

		adver.POST("/tab/second/isShow", handler.ShowOrHideSecondTab())
		adver.POST("/tab/second/son/isShow", handler.ShowOrHideSecondTabSon())
		adver.POST("/son/delete", handler.DeleteTabSon())
		adver.POST("/modify", handler.ModifyAndSaveSecondTab())
	}
	file := r.mux.Group("/v1/admin/file").Use(r.middles.Cors())
	{
		file.GET("/download/goods", handler.GoodsDownLoad())
		file.GET("/download/goods-mould", handler.GoodsDownLoadEmpty())
		file.GET("/download/prize", handler.PrizesDownLoad())
		file.GET("/download/order", handler.OrderDownLoad())
		file.GET("/download/luggage", handler.LuggageDownLoad())

		file.GET("/download/deliver", handler.DeliverDownLoad())
		file.GET("/download/deliver/detail", handler.DeliverDetailDownLoad())
	}
	luggage := r.mux.Group("/v1/admin/luggage").Use(r.middles.Cors())
	{
		luggage.POST("/delever", handler.Delever())
		luggage.POST("/delever/condition", handler.DeleverCondition())
		luggage.POST("/delever/detail", handler.DeleverDetail())
		luggage.POST("/delever/setDelId", handler.SetDelId())
	}

	//memoryStore := persist.NewMemoryStore(1 * time.Minute)
	//power.GET("/123", cache.CacheByRequestURI(memoryStore, 3600*time.Second),
	//	handler.GetUser(),
	//)
}
