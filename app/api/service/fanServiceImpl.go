package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
	"path"
	"strings"
	"time"
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
		}
		tx.Commit()
		return nil
	}
}
func (s *FanServiceImpl) QueryFanStatus(req param.ReqQueryFanStatus) (param.RespQueryFanStatus, error) {
	DB := s.db.GetDb()
	var fan db.Fan
	var resp param.RespQueryFanStatus
	result := DB.Order("created_at desc").First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespQueryFanStatus{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespQueryFanStatus{}, errors.New("没有蕃...")
	}
	resp.FanId = fan.ID
	resp.FanTitle = fan.Title
	DB.Model(&fan).Association("Boxs").Find(&fan.Boxs)
	for _, ele := range fan.Boxs {
		resp.BoxStatus = append(resp.BoxStatus, param.BoxStatus{
			BoxId:  ele.ID,
			Status: ele.Status,
		})
	}
	return resp, nil
}

func (s *FanServiceImpl) QueryFanStatusCondition(req param.ReqQueryFanStatusCondition) (param.RespQueryFanStatusCondition, error) {
	var fan db.Fan
	DB := s.db.GetDb()
	var resp param.RespQueryFanStatusCondition
	result := DB.Where("id=?", req.FanId).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespQueryFanStatusCondition{}, errors.New("服务正忙...")
	}
	resp.FanId = fan.ID
	resp.FanTitle = fan.Title
	if result.RowsAffected == 0 {
		return param.RespQueryFanStatusCondition{}, errors.New("没有蕃...")
	}

	boxes := []db.Box{}
	err := DB.Model(&fan).Association("Boxs").Find(&boxes)
	if err != nil {
		return param.RespQueryFanStatusCondition{}, errors.New("服务正忙...")
	}
	TotalBoxNum := 0
	TotalPrizeNum := int32(0)
	LeftPrizeNum := int32(0)
	for _, oneBox := range boxes {
		TotalBoxNum += 1
		err = DB.Model(&oneBox).Association("Prizes").Find(&oneBox.Prizes)
		if err != nil {
			return param.RespQueryFanStatusCondition{}, errors.New("服务正忙...")
		}
		for _, onePrize := range oneBox.Prizes {
			TotalPrizeNum += onePrize.PrizeNum
			LeftPrizeNum += onePrize.PriczeLeftNum
		}
	} //time.Unix(req.TimeRange[1], 0).Format("2006-01-02 15:04:05")
	price, _ := decimal.NewFromFloat32(float32(fan.Price)).Float64()
	resp.Status = fan.Status
	resp.Price = price
	resp.TotalBoxNum = TotalBoxNum
	resp.TotalPrizeNum = TotalPrizeNum
	resp.LeftPrizeNum = LeftPrizeNum
	resp.SharePic = fan.SharePic
	resp.DetailPic = fan.DetailPic
	resp.ActiveBeginTime = fan.ActiveBeginTime
	resp.ActiveEndTime = fan.ActiveEndTime
	resp.CreateTime = fan.CreatedAt
	resp.WhoUpdate = fan.WhoUpdate
	return resp, nil
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
	for _, oneFan := range fans {
		boxes := []db.Box{}
		result := DB.Model(&oneFan).Association("Boxs").Find(&boxes)
		if result != nil {
			return param.RespQueryFan{}, errors.New("服务正忙...")
		}
		TotalBoxNum := 0
		TotalPrizeNum := int32(0)
		LeftPrizeNum := int32(0)
		for _, oneBox := range boxes {
			TotalBoxNum += 1
			result = DB.Model(&oneBox).Association("Prizes").Find(&oneBox.Prizes)
			if result != nil {
				return param.RespQueryFan{}, errors.New("服务正忙...")
			}
			for _, onePrize := range oneBox.Prizes {
				TotalPrizeNum += onePrize.PrizeNum
				LeftPrizeNum += onePrize.PriczeLeftNum
			}
		}
		price, _ := decimal.NewFromFloat32(float32(oneFan.Price)).Float64()
		ret.FanInfos.Fans = append(ret.FanInfos.Fans, param.Fan{
			ID:              oneFan.ID,
			Title:           oneFan.Title,
			Status:          oneFan.Status,
			Price:           price,
			TotalBoxNum:     TotalBoxNum,
			TotalPrizeNum:   TotalPrizeNum,
			LeftPrizeNum:    LeftPrizeNum,
			SharePic:        oneFan.SharePic,
			DetailPic:       oneFan.DetailPic,
			ActiveBeginTime: oneFan.ActiveBeginTime,
			ActiveEndTime:   oneFan.ActiveEndTime,
			CreateTime:      oneFan.CreatedAt,
			WhoUpdate:       oneFan.WhoUpdate,
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
func (s *FanServiceImpl) EachBox(box *param.Box, fanId uint, fanName string) ([]*db.Prize, int32) {
	prizes := []*db.Prize{}
	prizeNum := int32(0)
	for _, prizeEle := range box.Prizes {
		if prizeEle.PrizeIndexName != define.PrizeIndexNameFirst &&
			prizeEle.PrizeIndexName != define.PrizeIndexNameGlobal &&
			prizeEle.PrizeIndexName != define.PrizeIndexNameLast {
			prizeNum += prizeEle.PrizeNum
		}
		prizes = append(prizes, &db.Prize{
			ID:                define.GetRandPrizeId(),
			GoodID:            prizeEle.GoodId,
			GoodName:          prizeEle.GoodName,
			FanId:             fanId,
			FanName:           fanName,
			Position:          prizeEle.Position,
			IpID:              prizeEle.IpId,
			IpName:            prizeEle.IpName,
			SeriesID:          prizeEle.SeriId,
			SeriesName:        prizeEle.SeriName,
			Pic:               prizeEle.Pic,
			PrizeNum:          prizeEle.PrizeNum,
			PriczeLeftNum:     prizeEle.PrizeNum,
			PrizeIndex:        prizeEle.PrizeIndex,
			PrizeIndexName:    prizeEle.PrizeIndexName,
			SingleOrMuti:      prizeEle.SingleOrMuti,
			Price:             prizeEle.Price,
			MultiIds:          prizeEle.MultiIds,
			PkgStatus:         prizeEle.PkgStatus,
			Remark:            prizeEle.Remark,
			TimeForSoldStatus: prizeEle.TimeForSoldStatus,
			SoldStatus:        define.YfPrizeStatusNotSoldOut,
		})
	}
	return prizes, prizeNum
}
func (s *FanServiceImpl) PkgBoxes(tx *gorm.DB, fanId uint, req param.ReqModifySaveFan, boxIndex int32, prizeNum int32) (*db.Box, error) {
	box := &db.Box{
		ID:            define.GetRandBoxId(),
		FanName:       req.Title,
		BoxIndex:      boxIndex,
		PriczeNum:     prizeNum,
		PriczeLeftNum: prizeNum,
		Status:        define.YfBoxStatusNotOnBoard,
	}
	err := tx.Create(box).Error
	if err != nil {
		return nil, errors.New("箱子创建失败...")
	}
	return box, nil
}
func (s *FanServiceImpl) ModifySaveFan(req param.ReqModifySaveFan) (param.RespModifySaveFan, error) {
	tx := s.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	{
		var fan db.Fan
		ret := tx.Where("id=?", req.FanID).First(&fan)
		if ret.Error != nil && ret.Error != gorm.ErrRecordNotFound {
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err := tx.Model(&db.FirstPrize{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.FirstPrize{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.LastPrize{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.LastPrize{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.GlobalPrize{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.GlobalPrize{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.Prize{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.Prize{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.Box{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.Box{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.Fan{}).Unscoped().Where("id=?", req.FanID).Delete(&db.Fan{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.Sure{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.Sure{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.Left{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.Left{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = tx.Model(&db.Target{}).Unscoped().Where("fan_id=?", req.FanID).Delete(&db.Target{}).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
	}
	var (
		fanIndex = 0
		fan      = &db.Fan{}
	)
	result := tx.Model(&db.Fan{}).Where("title=?", req.Title).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return param.RespModifySaveFan{}, errors.New("服务正忙...")
	}
	if result.RowsAffected != 0 {
		tx.Rollback()
		return param.RespModifySaveFan{}, errors.New("此蕃已经存在...")
	}
	result = tx.Table("yf_fan").Select("fan_index").Order("fan_index desc").First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return param.RespModifySaveFan{}, errors.New("服务正忙...")
	}
	if result.RowsAffected != 0 {
		fanIndex = fan.FanIndex
	}
	fanId := req.FanID
	fan = &db.Fan{
		ID:              fanId,
		Title:           req.Title,
		FanIndex:        fanIndex + 1,
		Status:          define.YfFanStatusNotOnBoard,
		Price:           float32(req.FanPrice),
		SharePic:        req.SharePic,
		DetailPic:       req.DetailPic,
		Rule:            req.Rule,
		ActiveBeginTime: req.ActiveBeginTime,
		ActiveEndTime:   req.ActiveEndTime,
	}
	err := tx.Create(fan).Error
	if err != nil {
		tx.Rollback()
		return param.RespModifySaveFan{}, errors.New("服务正忙...")
	}
	all := 0
	boxIds := []uint{}
	for index := 0; index < req.BoxNum; index++ {
		boxEle := req.Boxes
		prizes, allPrizeNum := s.EachBox(&boxEle, fanId, req.Title)
		if allPrizeNum == 0 {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("总奖品数为0...")
		}
		all = int(allPrizeNum)
		box, errx := s.PkgBoxes(tx, fanId, req, int32(index+1), allPrizeNum)
		if errx != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		boxIds = append(boxIds, box.ID)
		for nindex, ele := range prizes {
			prizes[nindex].BoxID = &box.ID
			prizes[nindex].BoxIndex = int(box.BoxIndex)
			if ele.PrizeIndexName != define.PrizeIndexNameFirst &&
				ele.PrizeIndexName != define.PrizeIndexNameLast &&
				ele.PrizeIndexName != define.PrizeIndexNameGlobal {
				req.Boxes.Prizes[nindex].Position = []int{-1, -1}
				rate, _ := decimal.NewFromFloat32(float32(ele.PriczeLeftNum)).Div(decimal.NewFromFloat32(float32(allPrizeNum))).Float64()
				prizes[nindex].PrizeRate = fmt.Sprintf("%.2f", 100*rate) + "%"
			}
		}
		if err = tx.Model(&box).Association("Prizes").Append(&prizes); err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
	}
	isPositionTwoOk := true
	tips1 := []param.Tips{}
	for n := 0; n < len(req.Boxes.Prizes); n++ {
		if len(req.Boxes.Prizes[n].Position) != 2 {
			isPositionTwoOk = false
			tips1 = append(tips1, param.Tips{
				PrizeName:      req.Boxes.Prizes[n].GoodName,
				PrizeIndex:     req.Boxes.Prizes[n].PrizeIndex,
				PrizeIndexName: req.Boxes.Prizes[n].PrizeIndexName,
				Message:        "特殊赏位置必须2位",
			})
		}
	}
	if !isPositionTwoOk {
		tx.Rollback()
		return param.RespModifySaveFan{Tips: tips1}, errors.New("位置必须2位...")
	}
	tips2 := []param.Tips{}
	isPositionOk := true
	for n := 0; n < len(req.Boxes.Prizes); n++ {
		if req.Boxes.Prizes[n].Position[1] > all {
			isPositionOk = false
			tips2 = append(tips2, param.Tips{
				PrizeName:      req.Boxes.Prizes[n].GoodName,
				PrizeIndex:     req.Boxes.Prizes[n].PrizeIndex,
				PrizeIndexName: req.Boxes.Prizes[n].PrizeIndexName,
				Message:        "位置最大值: " + fmt.Sprintf("%d", all),
			})
		}
	}
	if !isPositionOk {
		tx.Rollback()
		return param.RespModifySaveFan{Tips: tips2}, errors.New("位置范围错误...")
	}
	for n := 0; n < len(req.Boxes.Prizes); n++ {
		tmpPosition := "["
		for index := 0; index < len(req.Boxes.Prizes[n].Position); index++ {
			tmpPosition += fmt.Sprintf("%d,", req.Boxes.Prizes[n].Position[index])
		}
		positon := strings.TrimRight(tmpPosition, ",")
		positon += "]"
		err = tx.Table("yf_prize").Where("fan_id=? and good_id=? and prize_index_name=? and prize_index=?",
			fanId, req.Boxes.Prizes[n].GoodId, req.Boxes.Prizes[n].PrizeIndexName, req.Boxes.Prizes[n].PrizeIndex).
			Update("position", positon).Error
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
	}
	firstPos, lastPos, globalPos := []int{}, []int{}, []int{-1, -1}
	firstEles, lastEles, globalEles := []int{}, []int{}, []int{}
	for _, onePrize := range req.Boxes.Prizes {
		if onePrize.PrizeIndexName == "First" {
			for p := 0; p < len(onePrize.Position); p++ {
				firstPos = append(firstPos, onePrize.Position[p])
			}
			for i := 0; i < int(onePrize.PrizeNum); i++ {
				firstEles = append(firstEles, int(onePrize.PrizeIndex))
			}
		} else if onePrize.PrizeIndexName == "Last" {
			for p := 0; p < len(onePrize.Position); p++ {
				lastPos = append(lastPos, onePrize.Position[p])
			}
			for i := 0; i < int(onePrize.PrizeNum); i++ {
				lastEles = append(lastEles, int(onePrize.PrizeIndex))
			}
		} else if onePrize.PrizeIndexName == "全局" {
			for i := 0; i < int(onePrize.PrizeNum); i++ {
				globalEles = append(globalEles, int(onePrize.PrizeIndex))
			}
		}
	}
	for _, boxid := range boxIds {
		err = s.SetFirstPrizeAndPosition(tx, firstPos, firstEles, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = s.SetLastPrizeAndPosition(tx, lastPos, lastEles, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
		err = s.SetGlobalPrizeAndPostion(tx, globalPos, globalEles, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
		if err != nil {
			tx.Rollback()
			return param.RespModifySaveFan{}, errors.New("服务正忙...")
		}
	}
	tx.Commit()
	return param.RespModifySaveFan{}, nil
}
func (s *FanServiceImpl) SetFirstPrizeAndPosition(tx *gorm.DB, firstPosition db.GormList, firstEles db.GormList, fanId, boxId uint, seconds int64) error { //(req.ActiveEndTime-time.Now().Unix()))
	return tx.Save(&db.FirstPrize{
		FanId:       fanId,
		BoxId:       boxId,
		Pos:         firstPosition,
		PrizeIndexs: firstEles,
	}).Error
}
func (s *FanServiceImpl) SetLastPrizeAndPosition(tx *gorm.DB, lastPosition db.GormList, lastEles db.GormList, fanId, boxId uint, seconds int64) error { //(req.ActiveEndTime-time.Now().Unix()))
	return tx.Save(&db.LastPrize{
		FanId:       fanId,
		BoxId:       boxId,
		Pos:         lastPosition,
		PrizeIndexs: lastEles,
	}).Error
}
func (s *FanServiceImpl) SetGlobalPrizeAndPostion(tx *gorm.DB, globalPosition db.GormList, globalEles db.GormList, fanId, boxId uint, seconds int64) error { //(req.ActiveEndTime-time.Now().Unix()))
	return tx.Save(&db.GlobalPrize{
		FanId:       fanId,
		BoxId:       boxId,
		Pos:         globalPosition,
		PrizeIndexs: globalEles,
	}).Error
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
