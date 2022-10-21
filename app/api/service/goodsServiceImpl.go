package service

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"gorm.io/gorm"
	"math"
	"os"
	"yifan/app/api/param"
	"yifan/app/db"
	"yifan/pkg/define"
)

func (s *GoodsServiceImpl) ManyGoodsUpload() {
	//获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	xlsxPath := dir + "/import.xlsx"
	//打开文件路径
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, oneSheet := range xlsxFile.Sheets {
		if oneSheet.Name == "商品" {
			//读取每个sheet下面的行数据
			for index, row := range oneSheet.Rows {
				//读取每个cell的内容
				var g define.Goods
				for i, oneCell := range row.Cells {
					if i == 0 {
						g.Ip = oneCell.String()
					}
					if i == 1 {
						g.Series = oneCell.String()
					}
					if i == 2 {
						g.GoodsName = oneCell.String()
					}
					if i == 3 {
						g.PkgStatus = oneCell.String()
					}
					if i == 4 {
						g.PreStore = oneCell.String()
					}
					if i == 5 {
						g.Integral, _ = oneCell.Int()
					}
					if i == 6 {
						g.Pic = oneCell.String()
					}
					if i == 7 {
						g.Price, _ = oneCell.Float()
					}
					if i == 8 {
						g.CreatTime = oneCell.String()
					}
				}
				if index != 0 {
					define.DealWithOneGood(g)
				}
				//row.AddCell().Value = "测试一下新增"
			}
		}
	}
}
func (s *GoodsServiceImpl) UpLoadGoods(req param.ReqUpLoadGoods) (param.RespUpLoadGoods, error) {
	DB := s.db.GetDb()
	ret := param.RespUpLoadGoods{}
	for _, ele := range req.UpLoadGoods {
		var ip db.Ip
		{ //查看当前商品所属ip是否已经存在
			result := DB.Find(&ip, "name = ?", ele.IpName)
			if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
				return param.RespUpLoadGoods{}, errors.New("服务正忙...")
			}
			if result.RowsAffected == 0 {
				ret.IpIdSerId = append(ret.IpIdSerId, param.IpIdSerId{
					IpName:   ele.IpName,
					SerName:  ele.SeriesName,
					GoodName: ele.GoodsName,
					Tip:      "请先创建IP...",
				})
				continue
			}
		}
		var series db.Series
		{ //查看当前商品所属series是否已经存在
			result := DB.Find(&series, "name = ?", ele.SeriesName)
			if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
				return param.RespUpLoadGoods{}, errors.New("服务正忙...")
			}
			if result.RowsAffected == 0 {
				ret.IpIdSerId = append(ret.IpIdSerId, param.IpIdSerId{
					IpName:   ele.IpName,
					SerName:  ele.SeriesName,
					GoodName: ele.GoodsName,
					Tip:      "请先创建Series...",
				})
				continue
			}
		}
		var goods []db.Goods
		result := DB.Find(&goods, "name = ?", ele.GoodsName)
		if result.Error != nil {
			return param.RespUpLoadGoods{}, errors.New("服务正忙...")
		}
		isOk := true
		for _, oneGood := range goods {
			if (*oneGood.IpID == ip.ID) && (*oneGood.SeriesID == series.ID) &&
				(oneGood.Name == ele.GoodsName) {
				ret.IpIdSerId = append(ret.IpIdSerId, param.IpIdSerId{
					IpName:   ele.IpName,
					SerName:  ele.SeriesName,
					GoodName: ele.GoodsName,
					Tip:      "该商品已经存在...",
				})
				isOk = false
			}
		}
		if !isOk {
			continue
		}
		rId := define.GetRandGoodId()
		gs := &db.Goods{
			ID:           rId,
			IpID:         &ip.ID,
			IpName:       ip.Name,
			SeriesID:     &series.ID,
			SeriesName:   series.Name,
			Pic:          ele.Pic,
			Price:        ele.Price,
			Name:         ele.GoodsName,
			PkgStatus:    ele.PkgStatus,
			Integral:     ele.Integral,
			PreStore:     ele.PreStore,
			Introduce:    ele.Introduce,
			SingleOrMuti: ele.SingleOrMuti,
			MultiIds:     ele.MultiIds,
		}
		if err := DB.Create(gs).Error; err != nil {
			return param.RespUpLoadGoods{}, errors.New("服务正忙......")
		}
	}
	return ret, nil
}

func (s *GoodsServiceImpl) SearchGoods(req param.ReqSearchGoods) (param.RespSearchGoods, error) {
	DB := s.db.GetDb()
	goods := []db.Goods{}
	result := DB.Where("name=? or ip_name=? or series_name=?",
		req.Search, req.Search, req.Search).Find(&goods)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespSearchGoods{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespSearchGoods{}, errors.New("不存在此商品...")
	} else {
		ret := param.RespSearchGoods{}
		for _, ele := range goods {
			ret.Goods.GoodsId = ele.ID
			ret.Goods.Pic = ele.Pic
			ret.Goods.Price = ele.Price
			ret.Goods.Name = ele.Name
			ret.Goods.PkgStatus = ele.PkgStatus
			ret.Goods.Introduce = ele.Introduce
			ret.Goods.CreateTime = ele.CreatedAt
			ret.Goods.IpID = ele.IpID
			ret.Goods.IpName = ele.IpName
			ret.Goods.SeriesID = ele.SeriesID
			ret.Goods.SeriesName = ele.SeriesName
			ret.Goods.SingleOrMuti = ele.SingleOrMuti
			ret.Goods.MultiIds = ele.MultiIds
			ret.Goods.Integral = ele.Integral
			ret.Goods.PreStore = ele.PreStore
			ret.Goods.WhoUpdate = ele.WhoUpdate
		}
		return ret, nil
	}
}

//IP，系列，商品，商品拆解状态，现货/预售，图片，积分，单个/多个，零售价
func (s *GoodsServiceImpl) AddGoods(req param.ReqAddGoods) (uint, error) {
	DB := s.db.GetDb()
	var ip db.Ip
	{ //查看当前商品所属ip是否已经存在
		result := DB.Find(&ip, "id = ?", *req.IpId)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return 0, errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			return 0, errors.New("请先创建IP...")
		}
	}
	var series db.Series
	{ //查看当前商品所属series是否已经存在
		result := DB.Find(&series, "id = ?", *req.SeriesId)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return 0, errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			return 0, errors.New("请先创建Series...")
		}
	}
	var goods []db.Goods
	result := DB.Find(&goods, "name = ?", req.GoodsName)
	if result.Error != nil {
		return 0, errors.New("服务正忙...")
	}
	for _, oneGood := range goods {
		if (*oneGood.IpID == *req.IpId) && (*oneGood.SeriesID == *req.SeriesId) &&
			(oneGood.Name == req.GoodsName) {
			return 0, fmt.Errorf("该商品已经存在: IP:%v,系列:%v,商品:%v", ip.Name, series.Name, req.GoodsName)
		}
	}
	rId := define.GetRandGoodId()
	gs := &db.Goods{
		ID:           rId,
		IpID:         req.IpId,
		IpName:       req.IpName,
		SeriesID:     req.SeriesId,
		SeriesName:   req.SeriesName,
		Pic:          req.Pic,
		Price:        req.Price,
		Name:         req.GoodsName,
		SingleOrMuti: req.SingleOrMuti,
		MultiIds:     req.MultiIds,
		PkgStatus:    req.PkgStatus,
		Integral:     req.Integral,
		PreStore:     req.PreStore,
	}
	if err := DB.Create(gs).Error; err != nil {
		return 0, errors.New("服务正忙......")
	}
	return gs.ID, nil
}

func (s *GoodsServiceImpl) DeleteGoods(req param.ReqDeleteGoods) error {
	goods := &db.Goods{}
	tx := s.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id=?", *req.GoodsId).First(goods).Error; err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	if goods.ID == 0 {
		tx.Rollback()
		return errors.New("该商品不存在")
	}
	if err := s.db.GetDb().Model(&db.Goods{}).
		Where("id=?", *req.GoodsId).Delete(&db.Goods{}).Error; err != nil {
		return errors.New("服务正忙......")
	}
	return nil
}
func (s *GoodsServiceImpl) QueryGoods(req param.ReqQueryGoods) (param.RespQueryGoods, error) {
	total := int64(0)
	err := s.db.GetDb().Model(&db.Goods{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespQueryGoods{}, errors.New("服务正忙...")
	}
	if total == 0 {
		return param.RespQueryGoods{}, errors.New("不存在任何商品...")
	}
	resp := param.RespQueryGoods{}
	goods := []*db.Goods{}
	err = s.db.GetDb().Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("id desc").Find(&goods).Error
	if err != nil {
		return param.RespQueryGoods{}, errors.New("服务正忙...")
	}
	for _, ele := range goods {
		mids := []int{}
		for _, id := range ele.MultiIds {
			mids = append(mids, id)
		}
		resp.GoodsInfo.Goods = append(resp.GoodsInfo.Goods, param.Goods{
			GoodsId:      ele.ID,
			Pic:          ele.Pic,
			Price:        ele.Price,
			Name:         ele.Name,
			PkgStatus:    ele.PkgStatus,
			Introduce:    ele.Introduce,
			CreateTime:   ele.CreatedAt,
			IpID:         ele.IpID,
			IpName:       ele.IpName,
			SeriesID:     ele.SeriesID,
			SeriesName:   ele.SeriesName,
			SingleOrMuti: ele.SingleOrMuti,
			MultiIds:     mids,
			Integral:     ele.Integral,
			PreStore:     ele.PreStore,
			WhoUpdate:    ele.WhoUpdate,
		})
	}
	resp.GoodsInfo.Num = len(resp.GoodsInfo.Goods)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}

func (s *GoodsServiceImpl) ModifyGoods(req param.ReqModifyGoods) error {
	if *req.GoodsId == 0 {
		return errors.New("商品id不可为0")
	}
	DB := s.db.GetDb()
	goods := db.Goods{}
	result := DB.Where("id=?", *req.GoodsId).First(&goods)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return errors.New("该商品不存在...")
	}
	err := s.db.GetDb().Model(&goods).Where("id=?", *req.GoodsId).Updates(map[string]interface{}{
		"ip_id":             *req.IpID,
		"series_id":         *req.SeriesID,
		"pic":               req.Pic,
		"price":             req.Price,
		"name":              req.Name,
		"single_or_muti":    req.SingleOrMuti,
		"multi_ids":         req.MultiIds,
		"pkg_status":        req.PkgStatus,
		"introduce":         req.Introduce,
		"integral":          req.Integral,
		"preStore":          req.PreStore,
		"active_begin_time": req.ActiveBeginTime,
		"active_end_time":   req.ActiveEndTime,
		"who_update":        req.WhoUpdate,
	}).Error
	if err != nil {
		return errors.New("服务正忙......")
	}
	return nil
}
