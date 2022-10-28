package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"yifan/app/api/param"
	"yifan/app/db"
)

func (s *AdverServiceImpl) ActiveByMan(req param.ReqActiveByMan) (param.RespActiveByMan, error) {
	DB := s.db.GetDb()
	DB.Model(&db.AdverTab{}).Where("tab_name=? and tab_name_son=?", req.TabName, req.TabNameSon).Update("active_by_man", 2).
		Update("active_by_man_time", time.Now().Format("2006-01-02 15:04:05"))
	return param.RespActiveByMan{}, nil
}
func (s *AdverServiceImpl) SingleClick(req param.ReqSingleClick) (param.RespSingleClick, error) {
	DB := s.db.GetDb()
	var fans []db.Fan
	var resp param.RespSingleClick
	DB.Find(&fans)
	for _, one := range fans {
		resp.FanPicTitle = append(resp.FanPicTitle, param.FanPicTitle{
			FanId: one.ID,
			Pic:   one.SharePic,
			Title: one.Title,
		})
	}
	return resp, nil
}

func (s *AdverServiceImpl) DelBannerPic(req param.ReqDelBannerPic) error {
	return s.db.GetDb().Model(&db.Adver{}).Where("id=?", req.Id).Delete(&db.Adver{}).Error
}
func (s *AdverServiceImpl) SetBannerPic(req param.ReqSetBannerPic) error {
	for _, one := range req.Banners {
		s.db.GetDb().Create(&db.Adver{
			BannerTitle:       one.BannerTitle,
			Pic:               one.Pic,
			Remark:            one.Remark,
			BannerReleatedUrl: one.BannerReleatedUrl,
			ReleatedUrlType:   one.ReleatedUrlType,
			TipsAfterBanner:   one.TipsAfterBanner,
			ActiveBeginTime:   one.ActiveBeginTime,
			ActiveEndTime:     one.ActiveEndTime,
			IsHide:            true,
		})
	}
	return nil
}
func (s *AdverServiceImpl) GetBannerPic(req param.ReqGetBannerPic) (param.RespGetBannerPic, error) {
	DB := s.db.GetDb()
	var adver []db.Adver
	result := DB.Find(&adver)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespGetBannerPic{}, errors.New("服务正忙......")
	}
	var resp param.RespGetBannerPic
	for _, one := range adver {
		resp.Banners = append(resp.Banners, param.Banner{
			Id:                one.ID,
			BannerTitle:       one.BannerTitle,
			Pic:               one.Pic,
			BannerReleatedUrl: one.BannerReleatedUrl,
			ReleatedUrlType:   one.ReleatedUrlType,
			TipsAfterBanner:   one.TipsAfterBanner,
			ActiveBeginTime:   one.ActiveBeginTime,
			ActiveEndTime:     one.ActiveEndTime,
			Remark:            one.Remark,
			IsHide:            one.IsHide,
		})
	}
	return resp, nil
}
func (s *AdverServiceImpl) AddSecondTab(req param.ReqAddSecondTab) (param.RespAddSecondTab, error) {
	if req.TabTag != "推荐" && req.TabTag != "热门" {
		return param.RespAddSecondTab{}, errors.New(" 一级tab输入错误(推荐或者热门)...")
	}
	DB := s.db.GetDb()
	if req.TabTag == "推荐" {
		DB.Unscoped().Where("first_tab_name=?", req.TabTag).Delete(&db.AdverTab{})
		DB.Create(&db.AdverTab{
			FirstTabName:     req.TabTag,
			FirstTabNameSons: req.TabSons,
		})
	} else {
		DB.Unscoped().Where("second_tab_name=?", req.TabTag).Delete(&db.AdverTab{})
		DB.Create(&db.AdverTab{
			SecondTabName:     req.TabTag,
			SecondTabNameSons: req.TabSons,
		})
	}
	return param.RespAddSecondTab{}, nil
}
func (s *AdverServiceImpl) AddSecondTabSon(req param.ReqAddSecondTabSon) (param.RespAddSecondTabSon, error) {
	if req.TabTag != "推荐" && req.TabTag != "热门" {
		return param.RespAddSecondTabSon{}, errors.New(" 一级tab输入错误(推荐或者热门)...")
	}
	DB := s.db.GetDb()
	if req.TabTag == "推荐" {
		result := DB.Where("first_tab_name=?", req.TabTag).First(&db.AdverTab{})
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return param.RespAddSecondTabSon{}, errors.New("服务正忙......")
		}
		if result.RowsAffected == 0 {
			return param.RespAddSecondTabSon{}, errors.New("不存在:" + req.TabTag + "...")
		}
		var at db.AdverTab
		DB.Where("first_tab_name=?", req.TabTag).First(&at)
		isExist := false
		for _, one := range at.FirstTabNameSons {
			if one == req.TabSon {
				isExist = true
			}
		}
		if !isExist {
			return param.RespAddSecondTabSon{}, errors.New("不存在这样的二级tab:" + req.TabSon)
		}
		DB.Create(&db.AdverTab{
			TabName:         req.TabTag,
			TabNameSon:      req.TabSon,
			RedirectType:    req.RedirectType,
			RedirectAddress: req.RedirectAddress,
			ActiveBeginTime: req.ActiveBeginTime,
			ActiveEndTime:   req.ActiveEndTime,
			Remark:          req.Remark,
			Title:           req.Title,
			Pic:             req.Pic})
	} else {
		result := DB.Where("second_tab_name=?", req.TabTag).First(&db.AdverTab{})
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return param.RespAddSecondTabSon{}, errors.New("服务正忙......")
		}
		if result.RowsAffected == 0 {
			return param.RespAddSecondTabSon{}, errors.New("不存在:" + req.TabTag + "...")
		}

		var at db.AdverTab
		DB.Where("second_tab_name=?", req.TabTag).First(&at)
		isExist := false
		for _, one := range at.SecondTabNameSons {
			if one == req.TabSon {
				isExist = true
			}
		}
		if !isExist {
			return param.RespAddSecondTabSon{}, errors.New("不存在这样的二级tab:" + req.TabSon)
		}
		DB.Create(&db.AdverTab{
			TabName:         req.TabTag,
			TabNameSon:      req.TabSon,
			RedirectType:    req.RedirectType,
			RedirectAddress: req.RedirectAddress,
			ActiveBeginTime: req.ActiveBeginTime,
			ActiveEndTime:   req.ActiveEndTime,
			Remark:          req.Remark,
			Title:           req.Title,
			Pic:             req.Pic,
		})
	}
	return param.RespAddSecondTabSon{}, nil
}
func (s *AdverServiceImpl) QuerySecondTab(req param.ReqQuerySecondTab) (param.RespQuerySecondTab, error) {
	DB := s.db.GetDb()
	var at db.AdverTab
	if req.TabTag == "推荐" {
		DB.Where("first_tab_name=?", req.TabTag).First(&at)
		return param.RespQuerySecondTab{
			SecondTab: at.FirstTabNameSons,
		}, nil
	} else if req.TabTag == "热门" {
		DB.Where("second_tab_name=?", req.TabTag).First(&at)
		return param.RespQuerySecondTab{
			SecondTab: at.SecondTabNameSons,
		}, nil
	} else {
		return param.RespQuerySecondTab{}, errors.New("参数错误...")
	}

}
func (s *AdverServiceImpl) QuerySecondSonTab(req param.ReqQuerySecondSonTab) (param.RespQuerySecondSonTab, error) {
	DB := s.db.GetDb()
	var at []db.AdverTab
	if req.TabTag == "推荐" {
		DB.Where("tab_name=? and tab_name_son=? and is_hide=?", req.TabTag, req.TabSon, false).Find(&at)
	} else if req.TabTag == "热门" {
		DB.Where("tab_name=? and tab_name_son=? and is_hide=?", req.TabTag, req.TabSon, false).Find(&at)
	} else {
		return param.RespQuerySecondSonTab{}, errors.New("参数错误...")
	}
	var resp param.RespQuerySecondSonTab
	for _, ele := range at {
		resp.Tab = append(resp.Tab, param.Tab{
			TabTag:          req.TabTag,
			TabSon:          req.TabSon,
			RedirectType:    ele.RedirectType,
			RedirectAddress: ele.RedirectAddress,
			ActiveBeginTime: ele.ActiveBeginTime,
			ActiveEndTime:   ele.ActiveEndTime,
			Remark:          ele.Remark,
			Title:           ele.Title,
			Pic:             ele.Pic,
		})
	}
	return resp, nil
}

func (s *AdverServiceImpl) ShowOrHideBanner(req param.ReqShowOrHideBanner) error {
	s.db.GetDb().Model(&db.Adver{}).Where("id=?", req.BannerId).Update("is_hide", req.IsHide)
	return nil
}
func (s *AdverServiceImpl) ShowOrHideSecondTab(req param.ReqShowOrHideSecondTab) error {
	return s.db.GetDb().Model(&db.AdverTab{}).Where("tab_name=? and tab_name_son=?", req.TabTag, req.TabSon).Update("is_hide", req.IsHide).Error
}

func (s *AdverServiceImpl) ShowOrHideSecondTabSon(req param.ReqShowOrHideSecondTabSon) error {
	return s.db.GetDb().Model(&db.AdverTab{}).Where("id=? and tab_name=? and tab_name_son=?", req.Id, req.TabTag, req.TabSon).Update("is_son_hide", req.IsHide).Error
}

func (s *AdverServiceImpl) DeleteTabSon(req param.ReqDeleteTabSon) error {
	return s.db.GetDb().Model(&db.AdverTab{}).Where("id=?", req.Id).Delete(&db.AdverTab{}).Error
}
func (s *AdverServiceImpl) ModifyAndSaveSecondTab(req param.ReqModifyAndSaveSecondTab) (param.RespModifyAndSaveSecondTab, error) {
	fmt.Println("ModifyAndSaveSecondTab...")
	return param.RespModifyAndSaveSecondTab{}, nil
}
