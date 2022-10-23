package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
	"path"
	"yifan/app/api/param"
	"yifan/app/db"
	"yifan/pkg/define"
)

func (s *FanServiceImpl) ModifyFanStatus(req param.ReqModifyFanStatus) error {
	tx := s.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	fan := &db.Fan{}
	result := tx.Where("id=?", req.FanId).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("此蕃不存在...")
	}
	if req.Status == define.YfFanStatusDelete {
		if fan.Status == define.YfFanStatusOnBoardByMan || fan.Status == define.YfFanStatusOnBoardAuto {
			tx.Rollback()
			return errors.New("此蕃正在上架中，不可删除，如要删除，请先下架...")
		} else {
			result = tx.Model(&fan).Update("status", req.Status)
			if result.Error != nil {
				tx.Rollback()
				return errors.New("服务正忙...")
			}
			if result.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("删除失败...")
			}
			err := tx.Model(&fan).Association("Boxs").Find(&fan.Boxs)
			if err != nil {
				tx.Rollback()
				return errors.New("服务正忙...")
			}
			for _, oneBox := range fan.Boxs {
				result = tx.Model(&oneBox).Update("status", define.YfBoxStatusDelete)
				if result.Error != nil {
					tx.Rollback()
					return errors.New("服务正忙...")
				}
				if result.RowsAffected == 0 {
					tx.Rollback()
					return errors.New("箱子删除失败...")
				}
				err = tx.Model(&oneBox).Association("Prizes").Find(&oneBox.Prizes)
				if err != nil {
					tx.Rollback()
					return errors.New("服务正忙...")
				}
				for _, onePrize := range oneBox.Prizes {
					result = tx.Model(&onePrize).Update("status", define.YfPrizeStatusDelete)
					if result.Error != nil {
						tx.Rollback()
						return errors.New("服务正忙...")
					}
					if result.RowsAffected == 0 {
						tx.Rollback()
						return errors.New("服务正忙...")
					}
					result = tx.Model(&onePrize).Where("id=?", onePrize.ID).Delete(&db.Prize{})
					if result.Error != nil {
						tx.Rollback()
						return errors.New("服务正忙...")
					}
					if result.RowsAffected == 0 {
						tx.Rollback()
						return errors.New("箱子删除失败...")
					}
				}
				result = tx.Model(&fan.Boxs).Where("id=?", oneBox.ID).Delete(&db.Box{})
				if result.Error != nil {
					tx.Rollback()
					return errors.New("服务正忙...")
				}
				if result.RowsAffected == 0 {
					tx.Rollback()
					return errors.New("箱子删除失败...")
				}
			}
			result = tx.Model(&fan).Where("id=?", req.FanId).Delete(&db.Fan{})
			if result.Error != nil {
				tx.Rollback()
				return errors.New("服务正忙...")
			}
			if result.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("删除失败...")
			}
			tx.Commit()
			return nil
		}
	} else {
		result = tx.Model(&fan).Update("status", req.Status)
		if result.Error != nil {
			tx.Rollback()
			return errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("服务正忙...")
		}
		err := tx.Model(&fan).Association("Boxs").Find(&fan.Boxs)
		if err != nil {
			tx.Rollback()
			return errors.New("服务正忙...")
		}
		if req.Status == define.YfFanStatusNotOffBoardByMan {
			req.Status = define.YfBoxStatusNotOnBoard
		}
		for _, oneBox := range fan.Boxs {
			result = tx.Model(&oneBox).Update("status", req.Status)
			if result.Error != nil {
				tx.Rollback()
				return errors.New("服务正忙...")
			}
			if result.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("服务正忙...")
			}
			err = tx.Model(&oneBox).Association("Prizes").Find(&oneBox.Prizes)
			if err != nil {
				tx.Rollback()
				return errors.New("服务正忙...")
			}
			for _, onePrize := range oneBox.Prizes {
				result = tx.Model(&onePrize).Update("status", req.Status)
				if result.Error != nil {
					tx.Rollback()
					return errors.New("服务正忙...")
				}
				if result.RowsAffected == 0 {
					tx.Rollback()
					return errors.New("服务正忙...")
				}
			}
		}
		tx.Commit()
		return nil
	}
}

func (s *FanServiceImpl) QueryFan(req param.ReqQueryFan) (param.RespQueryFan, error) {
	DB := s.db.GetDb()
	total := int64(0)
	err := DB.Model(&db.Fan{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespQueryFan{}, errors.New("服务正忙...")
	}
	fans := []*db.Fan{}
	if err := DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&fans).Error; err != nil {
		return param.RespQueryFan{}, errors.New("服务正忙...")
	}
	ret := param.RespQueryFan{}
	for _, one := range fans {
		boxes := []db.Box{}
		result := DB.Model(&one).Association("Boxs").Find(&boxes)
		if result != nil {
			return param.RespQueryFan{}, errors.New("服务正忙...")
		}
		totalNum := 0
		leftNum := 0
		for _, ele := range boxes {
			if ele.Status == define.YfBoxStatusPrizeLeft ||
				ele.Status == define.YfBoxStatusPrizeNotLeft {
				totalNum++
			}
			if ele.Status == define.YfBoxStatusPrizeLeft {
				leftNum++
			}
		}
		price, _ := decimal.NewFromFloat32(float32(one.Price)).Float64()
		ret.FanInfos.Fans = append(ret.FanInfos.Fans, param.Fan{
			ID:              one.ID,
			Title:           one.Title,
			Status:          one.Status,
			Price:           price,
			TotalBoxNum:     totalNum,
			LeftBoxNum:      leftNum,
			SharePic:        one.SharePic,
			DetailPic:       one.DetailPic,
			ActiveBeginTime: one.ActiveBeginTime,
			ActiveEndTime:   one.ActiveEndTime,
			CreateTime:      one.CreatedAt,
			WhoUpdate:       one.WhoUpdate,
		})
	}
	ret.FanInfos.Num = len(ret.FanInfos.Fans)
	ret.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return ret, nil
}

func (s *FanServiceImpl) ModifyFan(req param.ReqModifyFan) (param.RespModifyFan, error) {
	DB := s.db.GetDb()
	ret := param.RespModifyFan{}
	var fan db.Fan
	result := DB.Where("id=?", req.FanId).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return ret, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return ret, errors.New("不存在任何蕃...")
	}
	totalBox := int64(0)
	err := DB.Model(&db.Box{}).Where("fan_id=?", req.FanId).Count(&totalBox).Error
	if err != nil {
		return ret, errors.New("服务正忙...")
	}
	box := db.Box{}
	result = DB.Where("fan_id=?", req.FanId).Find(&box)
	if result.Error != nil {
		return ret, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return ret, errors.New("此番没有箱子...")
	}
	var prize []db.Prize
	err = DB.Model(&box).Association("Prizes").Find(&prize)
	if err != nil {
		return ret, errors.New("服务正忙...")
	}
	mf := []param.EachBoxPrize{}
	for _, ele := range prize {
		mf = append(mf, param.EachBoxPrize{
			PrizeId:        ele.ID,
			GoodId:         ele.GoodID,
			PrizeName:      ele.GoodName,
			PrizeNum:       ele.PrizeNum,
			PrizeIndex:     ele.PrizeIndex,
			PrizeLeftNum:   ele.PriczeLeftNum,
			PrizeRate:      ele.PrizeRate,
			PrizeIndexName: ele.PrizeIndexName,
			Position:       ele.Position,
			Remark:         ele.Remark,
			IpId:           ele.IpID,
			IpName:         ele.IpName,
			SeriId:         ele.SeriesID,
			SeriName:       ele.SeriesName,
			Pic:            ele.Pic,
			PkgStatus:      ele.PkgStatus,
			SingleOrMuti:   ele.SingleOrMuti,
			MultiIds:       ele.MultiIds,
		})
	}
	price, _ := decimal.NewFromFloat32(float32(fan.Price)).Float64()
	ret.EachBoxPrize = mf
	ret.FanId = box.FanId
	ret.FanName = box.FanName
	ret.FanPrice = price
	ret.BoxNum = int(totalBox)
	ret.WhoUpdate = fan.WhoUpdate
	ret.ActiveBeginTime = fan.ActiveBeginTime
	ret.ActiveEndTime = fan.ActiveEndTime
	ret.Rule = fan.Rule
	ret.Title = fan.Title
	return ret, nil
}
func (s *FanServiceImpl) GetNewBoxes(fanId uint, boxOld, boxNew *db.Box) (pz db.Prize, err error) {
	type GoodIdPrizeIndex struct {
		GoodId     uint
		PrizeIndex int32
	}
	toDeletes := []GoodIdPrizeIndex{}
	for _, oPrize := range boxOld.Prizes {
		isNeedToDelete := true
		for _, nPrize := range boxNew.Prizes {
			if nPrize.GoodID == oPrize.GoodID && nPrize.PrizeIndex == oPrize.PrizeIndex {
				isNeedToDelete = false //更新
				tx := s.db.GetDb().Begin()
				err = tx.Table("yf_prize").Where("fan_id=? and good_id=? and prize_index=?",
					fanId, nPrize.GoodID, nPrize.PrizeIndex).Updates(map[string]interface{}{
					"good_id":          nPrize.GoodID,
					"good_name":        nPrize.GoodName,
					"ip_id":            nPrize.IpID,
					"ip_name":          nPrize.IpName,
					"series_id":        nPrize.SeriesID,
					"series_name":      nPrize.SeriesName,
					"pic":              nPrize.Pic,
					"price":            nPrize.Price,
					"prize_num":        nPrize.PrizeNum,
					"pricze_left_num":  nPrize.PriczeLeftNum,
					"prize_index":      nPrize.PrizeIndex,
					"prize_index_name": nPrize.PrizeIndexName,
					"remark":           nPrize.Remark,
					"single_or_muti":   nPrize.SingleOrMuti,
					"multi_ids":        nPrize.MultiIds,
					"pkg_status":       nPrize.PkgStatus,
					"who_update":       nPrize.WhoUpdate,
				}).Error
				if err != nil {
					tx.Rollback()
					return db.Prize{}, errors.New("服务正忙......")
				}
				tx.Commit()
			}
		}
		if isNeedToDelete {
			toDeletes = append(toDeletes, GoodIdPrizeIndex{
				GoodId:     oPrize.GoodID,
				PrizeIndex: oPrize.PrizeIndex,
			})
		}
	}
	tx := s.db.GetDb()
	for _, delEle := range toDeletes {
		err = tx.
			Where("fan_id=? and good_id=? and prize_index=?",
				fanId, delEle.GoodId, delEle.PrizeIndex).Delete(&db.Prize{}).Error
		if err != nil {
			tx.Rollback()
			return db.Prize{}, errors.New("服务正忙......")
		}
	}
	tx.Commit()
	tx = s.db.GetDb().Begin()
	for _, nPrize := range boxNew.Prizes {
		isNeedToAdd := true
		for _, oPrize := range boxOld.Prizes {
			if nPrize.GoodID == oPrize.GoodID && nPrize.PrizeIndex == oPrize.PrizeIndex {
				isNeedToAdd = false
			}
		}
		if isNeedToAdd {
			pz = db.Prize{
				ID:             define.GetRandPrizeId(),
				GoodID:         nPrize.GoodID,
				GoodName:       nPrize.GoodName,
				FanId:          fanId,
				FanName:        nPrize.FanName,
				IpID:           nPrize.IpID,
				IpName:         nPrize.IpName,
				SeriesID:       nPrize.SeriesID,
				SeriesName:     nPrize.SeriesName,
				BoxID:          nPrize.BoxID,
				Pic:            nPrize.Pic,
				Price:          nPrize.Price,
				PrizeNum:       nPrize.PrizeNum,
				PriczeLeftNum:  nPrize.PriczeLeftNum,
				PrizeIndex:     nPrize.PrizeIndex,
				PrizeIndexName: nPrize.PrizeIndexName,
				PrizeRate:      nPrize.PrizeRate,
				Remark:         nPrize.Remark,
				SingleOrMuti:   nPrize.SingleOrMuti,
				MultiIds:       nPrize.MultiIds,
				SoldStatus:     nPrize.SoldStatus,
				PkgStatus:      nPrize.PkgStatus,
				WhoUpdate:      nPrize.WhoUpdate,
			}
			err = tx.Create(&pz).Error
			if err != nil {
				tx.Rollback()
				return db.Prize{}, errors.New("服务正忙......")
			}
		}
	}
	tx.Commit()
	return pz, nil
}

func (s *FanServiceImpl) ModifySaveFan(req param.ReqModifySaveFan) (param.RespModifySaveFan, error) {
	DB := s.db.GetDb()
	var resp param.RespModifySaveFan
	var fan db.Fan
	ret := DB.Where("id=?", req.FanID).First(&fan)
	if ret.Error != nil && ret.Error != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if ret.RowsAffected == 0 {
		return resp, errors.New("不存在该蕃或者该番未处于上架状态...")
	}
	err := DB.Model(&fan).Association("Boxs").Find(&fan.Boxs)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if len(fan.Boxs) == 0 {
		return resp, errors.New("该蕃没有箱...")
	}
	if req.TotalBoxNum < len(fan.Boxs) {
		return resp, errors.New("必须大于已有箱数...")
	}
	err = DB.Model(&fan.Boxs[0]).Association("Prizes").Find(&fan.Boxs[0].Prizes)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if len(fan.Boxs[0].Prizes) == 0 {
		return resp, errors.New("该箱没有奖品...")
	}
	OldBox := &db.Box{
		Price:     float64(fan.Price),
		FanId:     fan.ID,
		FanName:   fan.Boxs[0].FanName,
		WhoUpdate: fan.Boxs[0].WhoUpdate,
	}
	for _, ePz := range fan.Boxs[0].Prizes {
		var newPz db.Prize
		newPz.ID = ePz.ID
		newPz.GoodID = ePz.GoodID
		newPz.GoodName = ePz.GoodName
		newPz.Remark = ePz.Remark
		newPz.IpID = ePz.IpID
		newPz.PrizeIndexName = ePz.PrizeIndexName
		newPz.PrizeIndex = ePz.PrizeIndex
		newPz.PkgStatus = ePz.PkgStatus
		newPz.MultiIds = ePz.MultiIds
		newPz.SingleOrMuti = ePz.SingleOrMuti
		newPz.IpName = ePz.IpName
		newPz.SeriesName = ePz.SeriesName
		newPz.SeriesID = ePz.SeriesID
		newPz.Price = ePz.Price
		newPz.Pic = ePz.Pic
		newPz.Position = ePz.Position
		OldBox.Prizes = append(OldBox.Prizes, &newPz)
	}
	Box := &db.Box{
		Price:   req.FanPrice,
		FanId:   req.FanID,
		FanName: req.FanName,
	}
	for _, ePz := range req.Prizes {
		var newPz db.Prize
		newPz.GoodID = ePz.GoodId
		newPz.GoodName = ePz.GoodName
		newPz.Remark = ePz.Remark
		newPz.IpID = ePz.IpId
		newPz.PrizeIndexName = ePz.PrizeIndexName
		newPz.PrizeIndex = ePz.PrizeIndex
		newPz.PkgStatus = ePz.PkgStatus
		newPz.MultiIds = ePz.MultiIds
		newPz.SingleOrMuti = ePz.SingleOrMuti
		newPz.IpName = ePz.IpName
		newPz.SeriesName = ePz.SeriName
		newPz.SeriesID = ePz.SeriId
		newPz.Pic = ePz.Pic
		newPz.Position = ePz.Position
		Box.Prizes = append(Box.Prizes, &newPz)
	}
	if req.TotalBoxNum == len(fan.Boxs) {
		s.GetNewBoxes(req.FanID, OldBox, Box)
		return param.RespModifySaveFan{}, nil
	} else {
		pz, _ := s.GetNewBoxes(req.FanID, OldBox, Box)
		for i := 0; i < req.TotalBoxNum-len(fan.Boxs); i++ {
			newBox := Box
			newBox.ID = define.GetRandBoxId()
			tx := s.db.GetDb().Begin()
			err = tx.Create(newBox).Error
			if err != nil {
				tx.Rollback()
				return param.RespModifySaveFan{}, errors.New("服务正忙...")
			}
			if err = tx.Model(newBox).Association("Prizes").Append(&pz); err != nil {
				tx.Rollback()
				return param.RespModifySaveFan{}, errors.New("服务正忙...")
			}
			tx.Commit()
		}
		return param.RespModifySaveFan{}, nil
	}
}

func (s *FanServiceImpl) GetFanStatus(req param.ReqModifyFan) int {
	return 0
}
func (s *FanServiceImpl) GetBoxStatus(req param.ReqModifyFan) int {
	return 0
}
func (s *FanServiceImpl) GetPrizeStatus(req param.ReqModifyFan) int {
	return 0
}
func (s *FanServiceImpl) EachBoxInfoByStatus(req param.ReqEnterFan, status ...int) (boxes []db.Box, result *gorm.DB) { //
	tmpStatus := []int{0}
	for _, e := range status {
		tmpStatus = append(tmpStatus, e)
	}
	result = s.db.GetDb().Where("fan_id=? and status IN ?", req.FanId, tmpStatus).Find(&boxes)
	return
}
func (s *FanServiceImpl) EachPrizeInfoByStatus(fanId uint, box *db.Box, status ...int) (prizes []db.Prize, err error) { //
	tmpStatus := []int{0}
	for _, e := range status {
		tmpStatus = append(tmpStatus, e)
	}
	err = s.db.GetDb().Model(box).Where("fan_id=? and  box_id=? and sold_status IN ?",
		fanId, box.ID, tmpStatus).Association("Prizes").Find(&prizes)
	return
}
func (s *FanServiceImpl) QueryPrizePostion(req param.ReqQueryPrizePostion) (param.RespQueryPrizePostion, error) {
	//ret := param.RespQueryPrizePostion{}
	//DB := s.db.GetDb()
	//fan := db.Fan{}
	//{
	//	result := DB.Where("id=?", req.FanId).First(&fan)
	//	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
	//		return ret, errors.New("服务正忙...")
	//	}
	//	if result.RowsAffected == 0 {
	//		return ret, errors.New("该蕃不存在...")
	//	}
	//}
	//boxes := []db.Box{}
	//{
	//	result := DB.Where("fan_id=?", req.FanId).Find(&boxes)
	//	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
	//		return ret, errors.New("服务正忙...")
	//	}
	//	if result.RowsAffected == 0 {
	//		return ret, errors.New("此番没有箱子...")
	//	}
	//}
	//for _, ele := range boxes {
	//	prizes := []db.Prize{}
	//	result := DB.Model(ele).Where("good_name=? and prize_index_name=?", req.PrizeName, req.PrizeIndexName).Association("Prizes").Find(&prizes)
	//	if result != nil {
	//		return ret, errors.New("服务正忙...")
	//	}
	//	positions := []db.PrizePosition{}
	//	result = DB.Model(ele).Where("good_name=? and prize_index_name=?", req.PrizeName, req.PrizeIndexName).Association("PrizePositions").Find(&positions)
	//	if result != nil {
	//		return ret, errors.New("服务正忙...")
	//	}
	//	for _, pEle := range prizes {
	//		for _, pPEle := range positions {
	//			if pEle.PrizeIndexName == pPEle.PrizeIndexName && pEle.GoodName == pPEle.GoodName {
	//				ret.QueryPrizePostions = append(ret.QueryPrizePostions, param.QueryPrizePostion{
	//					FanId:          req.FanId,
	//					FanTitle:       fan.Title,
	//					BoxId:          ele.ID,
	//					PrizeNum:       pEle.PrizeNum,
	//					PrizeIndexName: pEle.PrizeIndexName,
	//					PrizeName:      pEle.GoodName,
	//					Position:       pPEle.Position,
	//				})
	//			}
	//		}
	//	}
	//}
	//return ret, nil
	return param.RespQueryPrizePostion{}, nil
}

func (s *FanServiceImpl) ModifyGoodsPosition(req param.ReqModifyGoodsPosition) (param.RespModifyGoodsPosition, error) {
	//DB := s.db.GetDb()
	//fan := db.Fan{}
	//result := DB.Where("id=?", req.FanId).First(&fan)
	//if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
	//	return param.RespModifyGoodsPosition{}, errors.New("服务正忙...")
	//}
	//if result.RowsAffected == 0 {
	//	return param.RespModifyGoodsPosition{}, errors.New("该蕃不存在...")
	//}
	//box := db.Box{}
	//result = DB.Where("fan_id=? and id=?", req.FanId, req.BoxId).Find(&box)
	//if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
	//	return param.RespModifyGoodsPosition{}, errors.New("服务正忙...")
	//}
	//if result.RowsAffected == 0 {
	//	return param.RespModifyGoodsPosition{}, errors.New("此番没有箱子...")
	//}
	//poses := []db.PrizePosition{}
	//DB.Where("fan_id=? and box_id=?", req.FanId, req.BoxId).First(&poses)

	return param.RespModifyGoodsPosition{}, nil
}
func (s *FanServiceImpl) FileUpload(c *gin.Context) (interface{}, error) {
	file, err := c.FormFile("fileName")
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		c.SaveUploadedFile(file, dst)
	}
	return dst, nil
}
