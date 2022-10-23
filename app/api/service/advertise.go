package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type AdverService interface {
	ActiveByMan(req param.ReqActiveByMan) (param.RespActiveByMan, error)
	SingleClick(req param.ReqSingleClick) (param.RespSingleClick, error)

	SetBannerPic(req param.ReqSetBannerPic) error
	GetBannerPic(req param.ReqGetBannerPic) (param.RespGetBannerPic, error)
	AddSecondTab(req param.ReqAddSecondTab) (param.RespAddSecondTab, error)
	AddSecondTabSon(req param.ReqAddSecondTabSon) (param.RespAddSecondTabSon, error)
	QuerySecondTab(req param.ReqQuerySecondTab) (param.RespQuerySecondTab, error)
	QuerySecondSonTab(req param.ReqQuerySecondSonTab) (param.RespQuerySecondSonTab, error)
	ShowOrHideSecondTab(req param.ReqShowOrHideSecondTab) error
	ModifyAndSaveSecondTab(req param.ReqModifyAndSaveSecondTab) (param.RespModifyAndSaveSecondTab, error)
}

type AdverServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewAdverService(db db.Repo, cache *cache.CacheRepo) AdverService {
	return &AdverServiceImpl{
		db:    db,
		cache: cache,
	}
}
