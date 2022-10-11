package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yifan/app/api/param"
	"yifan/pkg/response"
)

// @Description UpLoadIPs
// @Tags UpLoadIPs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqUpLoadIPs body param.ReqUpLoadIPs true "批量上传"
// @Success 200 {object} response.responseSucess{data=param.RespUpLoadIPs} "desc"
// @Failure 200 {object} response.responseFailure
// @Router /v1/ip/upload [post]
func (h *handler) UpLoadIPs() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqUpLoadIPs
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.ipService.UpLoadIPs(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description SearchIP
// @Tags SearchIP
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqSearchIP body param.ReqSearchIP true "搜索条件"
// @Success 200 {object} response.responseSucess{data=param.RespSearchIp} "desc"
// @Failure 200 {object} response.responseFailure
// @Router /v1/ip/search [post]
func (h *handler) SearchIP() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqSearchIP
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.ipService.SearchIP(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description AddIP
// @Tags AddIP
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqAddIP body param.ReqAddIP true "name:名字; englishName:英文名; introduce:介绍; popularity:知名度; pic:图片; status:1.上架 2.下架"
// @Success 200 {object} response.responseSucess{data=int} "desc"
// @Failure 200 {object} response.responseFailure
// @Router /v1/ip/create [post]
func (h *handler) AddIP() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqAddIP
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.ipService.AddIP(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description DeleteIP
// @Tags DeleteIP
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqDeleteIP body param.ReqDeleteIP true "id"
// @Success 200 {object} response.responseSucess{data=int} "desc"
// @Failure 400 {object} response.responseFailure
// @Router /v1/ip/delete [post]
func (h *handler) DeleteIP() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqDeleteIP
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.ipService.DeleteIP(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}
func (h *handler) XYZ() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("1111111111111")
	}
}

// @Description QueryIP
// @Tags QueryIP
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqQueryIP body param.ReqQueryIP true "tag为1是所有ip,tag为2是某一个ip,此时id必填"
// @Success 200 {object} response.responseSucess{data=int} "desc"
// @Failure 400 {object} response.responseFailure
// @Router /v1/ip/query [post]
func (h *handler) QueryIP() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqQueryIP
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.ipService.QueryIP(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description ModifyIP
// @Tags ModifyIP
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyIP body param.ReqModifyIP true "id"
// @Success 200 {object} response.responseSucess{data=int} "desc"
// @Failure 400 {object} response.responseFailure
// @Router /v1/ip/modify [post]
func (h *handler) ModifyIP() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyIP
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.ipService.ModifyIP(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description UpLoadSeries
// @Tags UpLoadSeries
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqUpLoadSeries body param.ReqUpLoadSeries true "批量上传"
// @Success 200 {object} response.responseSucess{data=param.RespUpLoadSeries}
// @Failure 400 {object} response.responseFailure
// @Router /v1/series/upload [post]
func (h *handler) UpLoadSeries() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqUpLoadSeries
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.seriService.UpLoadSeries(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description SearchSeries
// @Tags SearchSeries
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqSearchSeries body param.ReqSearchSeries true "666"
// @Success 200 {object} response.responseSucess{data=param.RespSearchSeries}
// @Failure 400 {object} response.responseFailure
// @Router /v1/series/search [post]
func (h *handler) SearchSeries() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqSearchSeries
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.seriService.SearchSeries(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description AddSeries
// @Tags AddSeries
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqAddSeries body param.ReqAddSeries true "666"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/series/create [post]
func (h *handler) AddSeries() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqAddSeries
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.seriService.AddSeries(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description DeleteSeries
// @Tags DeleteSeries
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqDeleteSeries body param.ReqDeleteSeries true "888"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/series/delete [post]
func (h *handler) DeleteSeries() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqDeleteSeries
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.seriService.DeleteSeries(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description QuerySeries
// @Tags QuerySeries
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqQuerySeries body param.ReqQuerySeries true "6666"
// @Success 200 {object} response.responseSucess{data=param.RespQuerySeries}
// @Failure 400 {object} response.responseFailure
// @Router /v1/series/query [post]
func (h *handler) QuerySeries() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqQuerySeries
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.seriService.QuerySeries(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description ModifySeries
// @Tags ModifySeries
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifySeries body param.ReqModifySeries true "6666"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/series/modify [post]
func (h *handler) ModifySeries() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifySeries
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.seriService.ModifySeries(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description UpLoadGoods
// @Tags UpLoadGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqUpLoadGoods body param.ReqUpLoadGoods true "666"
// @Success 200 {object} response.responseSucess{data=param.RespUpLoadGoods}
// @Failure 400 {object} response.responseFailure
// @Router /v1/goods/upload [post]
func (h *handler) UpLoadGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqUpLoadGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.goodsService.UpLoadGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description SearchGoods
// @Tags SearchGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqSearchGoods body param.ReqSearchGoods true "666"
// @Success 200 {object} response.responseSucess{data=param.RespSearchGoods}
// @Failure 400 {object} response.responseFailure
// @Router /v1/goods/search [post]
func (h *handler) SearchGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqSearchGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.goodsService.SearchGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description AddGoods
// @Tags AddGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqAddGoods body param.ReqAddGoods true "ipId为商品对应IP;seriesId为商品对应系列;singleOrMuti(1为单一,2为两个组合,3为三个组合);multiIds为商品id集合;prizeIndex为A赏等;num保留字段固定填0;pkgStatus打包状态1.已经拆包,2未拆包"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/goods/create [post]
func (h *handler) AddGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqAddGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.goodsService.AddGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(struct {
			Id uint `json:"id"`
		}{Id: data}, context)
	}
}

// @Description DeleteGoods
// @Tags DeleteGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqDeleteGoods body param.ReqDeleteGoods true "id"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/goods/delete [post]
func (h *handler) DeleteGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqDeleteGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.goodsService.DeleteGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description QueryGoods
// @Tags QueryGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqQueryGoods body param.ReqQueryGoods true "pageSize每页大小,pageIndex第几页"
// @Success 200 {object} response.responseSucess{data=param.RespQueryGoods}
// @Failure 400 {object} response.responseFailure
// @Router /v1/goods/query [post]
func (h *handler) QueryGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqQueryGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.goodsService.QueryGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description ModifyGoods
// @Tags ModifyGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyGoods body param.ReqModifyGoods true "id"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/goods/modify [post]
func (h *handler) ModifyGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.goodsService.ModifyGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description AddBox
// @Tags AddBox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqAddBox body param.ReqAddBox true "看文档下面结构说明"
// @Success 200 {object} response.responseSucess{data=param.RespAddBox}
// @Failure 400 {object} response.responseFailure
// @Router /v1/box/create [post]
func (h *handler) AddBox() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqAddBox
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.boxService.AddBox(req)
		if err != nil {
			response.AbortWithBadRequestWithData(err, data, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}
func (h *handler) SetNormalPrizePosition() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqSetNormalPrizePosition
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.boxService.SetNormalPrizePosition(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description DeleteBox
// @Tags DeleteBox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqDeleteBox body param.ReqDeleteBox true "看文档下面结构说明"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/box/delete [post]
func (h *handler) DeleteBox() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqDeleteBox
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.boxService.DeleteBox(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description ModifyBox
// @Tags ModifyBox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyBox body param.ReqModifyBox true "看文档下面结构说明"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/box/modify [post]
func (h *handler) ModifyBox() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyBox
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.boxService.ModifyBox(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description ModifyBoxStatus
// @Tags ModifyBoxStatus
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyBoxStatus body param.ReqModifyBoxStatus true "看文档下面结构说明"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/box/modify/status [post]
func (h *handler) ModifyBoxStatus() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyBoxStatus
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.boxService.ModifyBoxStatus(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description QueryGoodsForBox
// @Tags QueryGoodsForBox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqQueryGoodsForBox body param.ReqQueryGoodsForBox true "看文档下面结构说明"
// @Success 200 {object} response.responseSucess{data=param.RespQueryGoodsForBox}
// @Failure 400 {object} response.responseFailure
// @Router /v1/box/goods/query [post]
func (h *handler) QueryGoodsForBox() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqQueryGoodsForBox
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.boxService.QueryGoodsForBox(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

func (h *handler) GoodsToBePrize() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqGoodsToBePrize
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.boxService.GoodsToBePrize(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description ModifyBoxGoods
// @Tags ModifyBoxGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyBoxGoods body param.ReqModifyBoxGoods true "看文档下面结构说明"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/box/goods/modify [post]
func (h *handler) ModifyBoxGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyBoxGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.boxService.ModifyBoxGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description DeleteBoxGoods
// @Tags DeleteBoxGoods
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqDeleteBoxGoods body param.ReqDeleteBoxGoods true "看文档下面结构说明"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/box/goods/delete [post]
func (h *handler) DeleteBoxGoods() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqDeleteBoxGoods
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.boxService.DeleteBoxGoods(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description AddFan
// @Tags AddFan
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqAddFan body param.ReqAddFan true "555"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/fan/create [post]
func (h *handler) AddFan() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqAddFan
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.fanService.AddFan(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description ModifyFanStatus
// @Tags ModifyFanStatus
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyFanStatus body param.ReqModifyFanStatus true "555"
// @Success 200 {object} response.responseSucess{data=int}
// @Failure 400 {object} response.responseFailure
// @Router /v1/fan/modify/status [post]
func (h *handler) ModifyFanStatus() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyFanStatus
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		err := h.fanService.ModifyFanStatus(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(nil, context)
	}
}

// @Description QueryFan
// @Tags QueryFan
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqQueryFan body param.ReqQueryFan true "pageSize每页大小,pageIndex第几页"
// @Success 200 {object} response.responseSucess{data=param.RespQueryFan}
// @Failure 400 {object} response.responseFailure
// @Router /v1/fan/query [post]
func (h *handler) QueryFan() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqQueryFan
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.fanService.QueryFan(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description ModifyFan
// @Tags ModifyFan
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyFan body param.ReqModifyFan true "666"
// @Success 200 {object} response.responseSucess{data=param.RespModifyFan}
// @Failure 400 {object} response.responseFailure
// @Router /v1/fan/modify [post]
func (h *handler) ModifyFan() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyFan
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.fanService.ModifyFan(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description ModifySaveFan
// @Tags ModifySaveFan
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifySaveFan body param.ReqModifySaveFan true "666"
// @Success 200 {object} response.responseSucess{data=param.RespModifySaveFan}
// @Failure 400 {object} response.responseFailure
// @Router /v1/fan/modify/save [post]
func (h *handler) ModifySaveFan() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifySaveFan
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.fanService.ModifySaveFan(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description QueryPrizePostion
// @Tags QueryPrizePostion
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqQueryPrizePostion body param.ReqQueryPrizePostion true "一次购买"
// @Success 200 {object} response.responseSucess{data=param.RespQueryPrizePostion}
// @Failure 400 {object} response.responseFailure
// @Router /v1/fan/queryPostion [post]
func (h *handler) QueryPrizePostion() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqQueryPrizePostion
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.fanService.QueryPrizePostion(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description ModifyGoodsPosition
// @Tags ModifyGoodsPosition
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqModifyGoodsPosition body param.ReqModifyGoodsPosition true "一次购买"
// @Success 200 {object} response.responseSucess{data=param.RespModifyGoodsPosition}
// @Failure 400 {object} response.responseFailure
// @Router /v1/fan/modifyPosition [post]
func (h *handler) ModifyGoodsPosition() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqModifyGoodsPosition
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.fanService.ModifyGoodsPosition(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description IsNew
// @Tags IsNew
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqIsNew body param.ReqIsNew true "666"
// @Success 200 {object} response.responseSucess{data=bool}
// @Failure 400 {object} response.responseFailure
// @Router /v1/user/isNew [post]
func (h *handler) IsNew() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqIsNew
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.userService.IsNew(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}

// @Description GetOpenId
// @Tags GetOpenId
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ReqGetOpenId body param.ReqGetOpenId true "666"
// @Success 200 {object} response.responseSucess{data=param.RespGetOpenId}
// @Failure 400 {object} response.responseFailure
// @Router /v1/user/query/openid [post]
func (h *handler) GetOpenId() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req param.ReqGetOpenId
		if err := context.ShouldBindJSON(&req); err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		data, err := h.userService.GetOpenId(req)
		if err != nil {
			response.AbortWithBadRequestWithError(err, context)
			return
		}
		response.ResposeSuccess(data, context)
	}
}
