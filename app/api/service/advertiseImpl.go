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
			Title:           one.Title,
			AdvPic:          one.AdvPic,
			FanId:           one.FanId,
			FanTitle:        one.FanTitle,
			RedirectUrl:     one.RedirectUrl,
			RedirectType:    one.RedirectType,
			ActiveBeginTime: one.ActiveBeginTime,
			ActiveEndTime:   one.ActiveEndTime,
			IsHide:          one.IsHide,
			Remark:          one.Remark,
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
			Id:              one.ID,
			Title:           one.Title,
			AdvPic:          one.AdvPic,
			FanId:           one.FanId,
			FanTitle:        one.FanTitle,
			RedirectUrl:     one.RedirectUrl,
			RedirectType:    one.RedirectType,
			ActiveBeginTime: one.ActiveBeginTime,
			ActiveEndTime:   one.ActiveEndTime,
			IsHide:          one.IsHide,
			Remark:          one.Remark,
		})
	}
	return resp, nil
}
func (s *AdverServiceImpl) AddSecondTab(req param.ReqAddSecondTab) (param.RespAddSecondTab, error) {
	if req.TabTag != "推荐" && req.TabTag != "热门" {
		return param.RespAddSecondTab{}, errors.New(" 一级tab输入错误(推荐或者热门)...")
	}
	DB := s.db.GetDb()
	DB.Unscoped().Where("tab_tag=?", req.TabTag).Delete(&db.AdverTab{})
	DB.Create(&db.AdverTab{
		TabTag:   req.TabTag,
		TabSon1:  req.TabSon1,
		TabSon2:  req.TabSon2,
		TabSon3:  req.TabSon3,
		TabSon4:  req.TabSon4,
		TabSon5:  req.TabSon5,
		TabSon6:  req.TabSon6,
		TabSon7:  req.TabSon7,
		TabSon8:  req.TabSon8,
		TabSon9:  req.TabSon9,
		TabSon10: req.TabSon10,
		IsHide1:  req.IsHide1,
		IsHide2:  req.IsHide2,
		IsHide3:  req.IsHide3,
		IsHide4:  req.IsHide4,
		IsHide5:  req.IsHide5,
		IsHide6:  req.IsHide6,
		IsHide7:  req.IsHide7,
		IsHide8:  req.IsHide8,
		IsHide9:  req.IsHide9,
		IsHide10: req.IsHide10,
	})
	return param.RespAddSecondTab{}, nil
}
func (s *AdverServiceImpl) AddSecondTabSon(req param.ReqAddSecondTabSon) (param.RespAddSecondTabSon, error) {
	DB := s.db.GetDb()
	result := DB.Where("tab_tag=? and tab_son=?", req.TabTag, req.TabSon).First(&db.AdverContent{})
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespAddSecondTabSon{}, errors.New("服务正忙......")
	}
	DB.Create(&db.AdverContent{
		TabTag:          req.TabTag,
		TabSon:          req.TabSon,
		IsHide:          true,
		ActiveBeginTime: 0,
		ActiveEndTime:   0,
		Title:           req.Title,
		FanId:           req.FanId,
		FanTitle:        req.FanTitle,
		AdvPic:          req.AdvPic,
		Remark:          req.Remark,
		RedirectUrl:     req.RedirectUrl,
		RedirectType:    req.RedirectType,
	})
	return param.RespAddSecondTabSon{}, nil
}
func (s *AdverServiceImpl) QuerySecondTab(req param.ReqQuerySecondTab) (param.RespQuerySecondTab, error) {
	DB := s.db.GetDb()
	var at db.AdverTab
	var resp param.RespQuerySecondTab
	DB.Where("tab_tag=?", req.TabTag).Find(&at)
	resp.TabSon1 = at.TabSon1
	resp.TabSon2 = at.TabSon2
	resp.TabSon3 = at.TabSon3
	resp.TabSon4 = at.TabSon4
	resp.TabSon5 = at.TabSon5
	resp.TabSon6 = at.TabSon6
	resp.TabSon7 = at.TabSon7
	resp.TabSon8 = at.TabSon8
	resp.TabSon9 = at.TabSon9
	resp.TabSon10 = at.TabSon10
	resp.IsHide1 = at.IsHide1
	resp.IsHide2 = at.IsHide2
	resp.IsHide3 = at.IsHide3
	resp.IsHide4 = at.IsHide4
	resp.IsHide5 = at.IsHide5
	resp.IsHide6 = at.IsHide6
	resp.IsHide7 = at.IsHide7
	resp.IsHide8 = at.IsHide8
	resp.IsHide9 = at.IsHide9
	resp.IsHide10 = at.IsHide10
	return resp, nil
}
func (s *AdverServiceImpl) QuerySecondSonTab(req param.ReqQuerySecondSonTab) (param.RespQuerySecondSonTab, error) {
	DB := s.db.GetDb()
	var adverConts []db.AdverContent
	DB.Where("tab_tag=? and tab_son=?", req.TabTag, req.TabSon).Find(&adverConts)
	var resp param.RespQuerySecondSonTab
	for _, ele := range adverConts {
		resp.TabContent = append(resp.TabContent, param.TabContent{
			Id:              ele.ID,
			TabTag:          req.TabTag,
			TabSon:          req.TabSon,
			ActiveBeginTime: ele.ActiveBeginTime,
			ActiveEndTime:   ele.ActiveEndTime,
			Remark:          ele.Remark,
			Title:           ele.Title,
			AdvPic:          ele.AdvPic,
			IsHide:          ele.IsHide,
			RedirectUrl:     ele.RedirectUrl,
			FanId:           ele.FanId,
			FanTitle:        ele.FanTitle,
			RedirectType:    ele.RedirectType,
		})
	}
	return resp, nil
}

func (s *AdverServiceImpl) ShowOrHideBanner(req param.ReqShowOrHideBanner) error {
	s.db.GetDb().Model(&db.Adver{}).Where("id=?", req.BannerId).Update("is_hide", req.IsHide)
	return nil
}
func (s *AdverServiceImpl) ShowOrHideSecondTab(req param.ReqShowOrHideSecondTab) error {
	if req.IsHide1 == 1 || req.IsHide1 == 2 {
		if req.IsHide1 == 1 {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide1", true)
		} else {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide1", false)
		}
	}
	if req.IsHide2 == 1 || req.IsHide2 == 2 {
		if req.IsHide2 == 1 {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide2", true)
		} else {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide2", false)
		}
	}
	if req.IsHide3 == 1 || req.IsHide3 == 2 {
		if req.IsHide3 == 1 {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide3", true)
		} else {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide3", false)
		}
	}
	if req.IsHide4 == 1 || req.IsHide4 == 2 {
		if req.IsHide4 == 1 {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide4", true)
		} else {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide4", false)
		}
	}
	if req.IsHide5 == 1 || req.IsHide5 == 2 {
		if req.IsHide5 == 1 {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide5", true)
		} else {
			s.db.GetDb().Model(&db.AdverTab{}).Where("tab_tag=?", req.TabTag).Update("is_hide5", false)
		}
	}
	return nil
}

func (s *AdverServiceImpl) ShowOrHideSecondTabSon(req param.ReqShowOrHideSecondTabSon) error {
	return s.db.GetDb().Model(&db.AdverContent{}).Where("id=? and tab_tag=? and tab_son=?", req.Id, req.TabTag, req.TabSon).Update("is_hide", req.IsHide).Error
}

func (s *AdverServiceImpl) DeleteTabSon(req param.ReqDeleteTabSon) error {
	return s.db.GetDb().Model(&db.AdverTab{}).Where("id=?", req.Id).Delete(&db.AdverTab{}).Error
}
func (s *AdverServiceImpl) ModifyAndSaveSecondTab(req param.ReqModifyAndSaveSecondTab) (param.RespModifyAndSaveSecondTab, error) {
	fmt.Println("ModifyAndSaveSecondTab...")
	return param.RespModifyAndSaveSecondTab{}, nil
}
