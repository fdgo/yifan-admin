package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"

	//"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
	"strconv"
	"yifan/app/api/param"
	"yifan/app/db"
	"yifan/pkg/define"
	"yifan/pkg/randId"
)

func (s *FanServiceImpl) AddFan(req param.ReqAddFan) (uint, error) {
	DB := s.db.GetDb()
	fan := &db.Fan{}
	result := DB.Where("name=?", req.FanName).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return 0, errors.New("服务正忙...")
	}
	if result.RowsAffected != 0 {
		return 0, errors.New("此蕃已经存在...")
	}
	fanId := define.GetRandFanId()
	if err := DB.Create(&db.Fan{
		ID:              fanId,
		Name:            req.FanName,
		Status:          req.Status,
		Price:           req.FanPrice,
		Pic:             req.Pic,
		WhoUpdate:       req.WhoCreated,
		ActiveBeginTime: req.OnActiveTime,
		ActiveEndTime:   req.OffActiveTime,
	}).Error; err != nil {
		return 0, errors.New("服务正忙...")
	}
	return fanId, nil
}

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
		ret.FanInfos.Fans = append(ret.FanInfos.Fans, param.Fan{
			ID:              one.ID,
			Name:            one.Name,
			Status:          one.Status,
			Price:           one.Price,
			TotalBoxNum:     totalNum,
			LeftBoxNum:      leftNum,
			Pic:             one.Pic,
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
		pos := db.PrizePosition{}
		DB.Model(&db.PrizePosition{}).Where("fan_id=? and box_id=? and prize_index=?",
			req.FanId, box.ID, ele.PrizeIndex).First(&pos)
		mf = append(mf, param.EachBoxPrize{
			PrizeId:        ele.ID,
			GoodId:         ele.GoodID,
			PrizeName:      ele.GoodName,
			PrizeNum:       ele.PrizeNum,
			PrizeIndex:     ele.PrizeIndex,
			PrizeLeftNum:   ele.PriczeLeftNum,
			PrizeRate:      ele.PrizeRate,
			PrizeIndexName: ele.PrizeIndexName,
			Position:       pos.Position,
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
	ret.EachBoxPrize = mf
	ret.Type = box.Type
	ret.Rule = box.Rule
	ret.Title = box.Title
	ret.FanId = box.FanId
	ret.FanName = box.FanName
	ret.FanPrice = box.Price
	ret.ActiveBeginTime = box.ActiveBeginTime
	ret.ActiveEndTime = box.ActiveEndTime
	ret.DetailPic = box.DetailPic
	ret.SharePic = box.SharePic
	ret.BoxNum = int(totalBox)
	ret.WhoUpdate = fan.WhoUpdate
	return ret, nil
}
func (s *FanServiceImpl) ModifySaveFan(req param.ReqModifySaveFan) (param.RespModifySaveFan, error) {
	DB := s.db.GetDb()
	var resp param.RespModifySaveFan
	var fan db.Fan
	ret := DB.Where("id=? and status IN ?", req.FanID, []int{define.YfFanStatusNotOnBoard,
		define.YfFanStatusNotOffBoardByMan, define.YfFanStatusNotOffBoardAuto}).First(&fan)
	if ret.Error != nil && ret.Error != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if ret.RowsAffected == 0 {
		return resp, errors.New("不存在该蕃或者该番处于上架状态...")
	}
	err := DB.Model(&fan).Association("Boxs").Find(&fan.Boxs)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if req.TotalBoxNum > len(fan.Boxs) {

	} else {

	}
	//for _, oneBox := range fan.Boxs {
	//	err := DB.Model(&oneBox).Association("Prizes").Find(&fan.Boxs)
	//	err := DB.Model(&oneBox).Association("PrizePositions").Find(&fan.Boxs)
	//}
	return param.RespModifySaveFan{}, nil
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
	ret := param.RespQueryPrizePostion{}
	DB := s.db.GetDb()
	fan := db.Fan{}
	{
		result := DB.Where("id=?", req.FanId).First(&fan)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return ret, errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			return ret, errors.New("该蕃不存在...")
		}
	}
	boxes := []db.Box{}
	{
		result := DB.Where("fan_id=?", req.FanId).Find(&boxes)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return ret, errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			return ret, errors.New("此番没有箱子...")
		}
	}
	for _, ele := range boxes {
		prizes := []db.Prize{}
		result := DB.Model(ele).Where("good_name=? and prize_index_name=?", req.PrizeName, req.PrizeIndexName).Association("Prizes").Find(&prizes)
		if result != nil {
			return ret, errors.New("服务正忙...")
		}
		positions := []db.PrizePosition{}
		result = DB.Model(ele).Where("good_name=? and prize_index_name=?", req.PrizeName, req.PrizeIndexName).Association("PrizePositions").Find(&positions)
		if result != nil {
			return ret, errors.New("服务正忙...")
		}
		for _, pEle := range prizes {
			for _, pPEle := range positions {
				if pEle.PrizeIndexName == pPEle.PrizeIndexName && pEle.GoodName == pPEle.GoodName {
					ret.QueryPrizePostions = append(ret.QueryPrizePostions, param.QueryPrizePostion{
						FanId:          req.FanId,
						FanName:        fan.Name,
						BoxId:          ele.ID,
						BoxTitle:       ele.Title,
						PrizeNum:       pEle.PrizeNum,
						PrizeIndexName: pEle.PrizeIndexName,
						PrizeName:      pEle.GoodName,
						Position:       pPEle.Position,
					})
				}
			}
		}
	}
	return ret, nil
}

func (s *FanServiceImpl) ModifyGoodsPosition(req param.ReqModifyGoodsPosition) (param.RespModifyGoodsPosition, error) {
	DB := s.db.GetDb()
	fan := db.Fan{}
	result := DB.Where("id=?", req.FanId).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespModifyGoodsPosition{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespModifyGoodsPosition{}, errors.New("该蕃不存在...")
	}
	box := db.Box{}
	result = DB.Where("fan_id=? and id=?", req.FanId, req.BoxId).Find(&box)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespModifyGoodsPosition{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespModifyGoodsPosition{}, errors.New("此番没有箱子...")
	}
	poses := []db.PrizePosition{}
	DB.Where("fan_id=? and box_id=?", req.FanId, req.BoxId).First(&poses)

	return param.RespModifyGoodsPosition{}, nil
}
func (s *FanServiceImpl) Buy(req param.ReqBuy) (param.RespBuy, error) {
	if req.Times != 1 && req.Times != 5 && req.Times != 10 && req.Times != -1 {
		return param.RespBuy{}, errors.New("times有误...")
	}
	DB := s.db.GetDb()
	fan := db.Fan{}
	result := DB.Where("id=?", req.FanId).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespBuy{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespBuy{}, errors.New("该蕃不存在...")
	}
	box := db.Box{}
	result = DB.Where("id=?", req.BoxId).First(&box)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespBuy{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespBuy{}, errors.New("该箱不存在...")
	}
	var prize []db.Prize
	err := DB.Model(&box).Where("sold_status=?", define.YfPrizeStatusNotSoldOut).Association("Prizes").Find(&prize)
	if err != nil {
		return param.RespBuy{}, errors.New("服务正忙...")
	}
	maxLeft := int32(0)
	for _, ele := range prize {
		maxLeft += ele.PriczeLeftNum
	}
	if req.Times != -1 {
		if maxLeft < int32(req.Times) {
			return param.RespBuy{}, errors.New("库存不足...")
		} else {
			return param.RespBuy{
				Money: float64(req.Times) * fan.Price,
			}, nil

		}
	} else {
		return param.RespBuy{
			Money: float64(maxLeft) * fan.Price,
		}, nil
	}
}

func (s *FanServiceImpl) BuySure(req param.ReqBuySure) (param.RespBuySures, error) {
	var resp param.RespBuySures
	if req.Times != 1 && req.Times != 5 && req.Times != 10 && req.Times != -1 {
		return resp, errors.New("times有误...")
	}
	type tmpFan struct {
		Id   uint
		Name string
	}
	var fan tmpFan
	DB := s.db.GetDb()
	result := DB.Table("yf_fan").Select("id", "name").Where("id=? and status IN ?", req.FanId, []int{define.YfFanStatusOnBoardByMan, define.YfFanStatusOnBoardAuto}).Scan(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return resp, errors.New("商品已经售罄...")
	}
	type tmpBox struct {
		Id    uint
		Title string
	}
	var box tmpBox
	result = DB.Table("yf_box").Select("id", "title").Where("fan_id=? and id=? and status IN ?", req.FanId, req.BoxId, []int{define.YfBoxStatusPrizeLeft, define.YfBoxStatusPrizeNotLeft}).Scan(&box)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return resp, errors.New("商品已经售罄...")
	}
	var prizes []*db.TmpPrize
	result = DB.Table("yf_prize").Select("id", "good_id", "good_name", "fan_id", "fan_name", "box_id", "prize_num", "pricze_left_num", "prize_index", "prize_index_name").Where("fan_id=? and box_id=? and sold_status IN ? and prize_index_name<>? and prize_index_name<>? and prize_index_name<>? and pricze_left_num>?",
		req.FanId, req.BoxId, []int{define.YfPrizeStatusNotSoldOut}, define.PrizeIndexNameFirst, define.PrizeIndexNameLast, define.PrizeIndexNameGlobal, 0).Scan(&prizes)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return resp, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return resp, errors.New("商品已经售罄...")
	}
	num := int32(0)
	for _, ele := range prizes {
		num += ele.PriczeLeftNum
	}
	if req.Times != -1 {
		if num < int32(req.Times) {
			return resp, errors.New("库存不足...")
		}
		for i := 0; i < req.Times; i++ {
			var res param.RespBuySure
			var errx error
			res, prizes, errx = s.makesure(req, fan.Name, box.Title, prizes)
			if errx != nil {
				return resp, errx
			}
			resp.BuySures = append(resp.BuySures, param.BuySure{
				Index:          i + 1,
				PrizeIndex:     res.BuySure.PrizeIndex,
				PrizeIndexName: res.BuySure.PrizeIndexName,
				PrizeName:      res.BuySure.PrizeName,
				Price:          res.BuySure.Price,
				Pic:            res.BuySure.Pic,
			})
		}
		prizesLeftNums := int32(0)
		for _, ele := range prizes {
			if ele.PrizeIndexName != "Last" && ele.PrizeIndexName != "First" && ele.PrizeIndexName != "全局" {
				prizesLeftNums += ele.PriczeLeftNum
			}
		}
		for _, ele := range prizes {
			rate := float64(0)
			if prizesLeftNums != 0 {
				rate, _ = decimal.NewFromFloat32(float32(ele.PriczeLeftNum)).Div(decimal.NewFromFloat32(float32(prizesLeftNums))).Float64()
			} else {
				rate = 0
			}
			DB.Model(&db.Prize{}).Where("fan_id=? and box_id=? and sold_status IN ? and prize_index_name<>? and prize_index_name<>? and prize_index_name<>? and prize_index=?",
				req.FanId, req.BoxId, []int{define.YfPrizeStatusNotSoldOut},
				define.PrizeIndexNameFirst, define.PrizeIndexNameLast, define.PrizeIndexNameGlobal,
				ele.PrizeIndex).Update("prize_rate", fmt.Sprintf("%.2f", 100*rate)+"%")
		}

	} else {
		for i := 0; i < int(num); i++ {
			var res param.RespBuySure
			var errx error
			res, prizes, errx = s.makesure(req, fan.Name, box.Title, prizes)
			if errx != nil {
				return resp, errx
			}
			resp.BuySures = append(resp.BuySures, param.BuySure{
				Index:          i + 1,
				PrizeIndex:     res.BuySure.PrizeIndex,
				PrizeIndexName: res.BuySure.PrizeIndexName,
				PrizeName:      res.BuySure.PrizeName,
				Price:          res.BuySure.Price,
				Pic:            res.BuySure.Pic,
			})
		}
		prizesLeftNums := int32(0)
		for _, ele := range prizes {
			if ele.PrizeIndexName != "Last" && ele.PrizeIndexName != "First" && ele.PrizeIndexName != "全局" {
				prizesLeftNums += ele.PriczeLeftNum
			}
		}
		for _, ele := range prizes {
			rate := float64(0)
			if prizesLeftNums != 0 {
				rate, _ = decimal.NewFromFloat32(float32(ele.PriczeLeftNum)).Div(decimal.NewFromFloat32(float32(prizesLeftNums))).Float64()
			} else {
				rate = 0
			}
			DB.Model(&db.Prize{}).Where("fan_id=? and box_id=? and sold_status IN ? and prize_index_name<>? and prize_index_name<>? and prize_index_name<>? and prize_index=?",
				req.FanId, req.BoxId, []int{define.YfPrizeStatusNotSoldOut},
				define.PrizeIndexNameFirst, define.PrizeIndexNameLast, define.PrizeIndexNameGlobal,
				ele.PrizeIndex).Update("prize_rate", fmt.Sprintf("%.2f", 100*rate)+"%")
		}
	}
	ctx := context.Background()
	nindex, targetLen := s.getCurrentPrizeIndexCache(context.Background(), fan.Id, box.Id)
	go s.DealWithFirstPrize(ctx, fan.Id, box.Id, fan.Name, box.Title, nindex, targetLen)
	go s.DealWithLastPrize(ctx, fan.Id, box.Id, fan.Name, box.Title, nindex, targetLen)
	go s.DealWithGlobalPrize(ctx, fan.Id, box.Id, fan.Name, box.Title, nindex, targetLen)
	return resp, nil
}

func (s *FanServiceImpl) getFirstPrizeEleCache(ctx context.Context, fanId, boxId uint) (resp []define.PrizeIdIndexName, err error) { //(req.ActiveEndTime-time.Now().Unix()))
	lRange := s.cache.GetCache().LRange(ctx, define.RedisFirstPrize(fanId, boxId), 0, -1).Val()
	if len(lRange) == 0 {
		return []define.PrizeIdIndexName{}, errors.New("First已经发完...")
	}
	for _, ele := range lRange {
		e := define.PrizeIdIndexName{}
		json.Unmarshal([]byte(ele), &e)
		resp = append(resp, e)
	}
	return
}
func (s *FanServiceImpl) getLastPrizeEleCache(ctx context.Context, fanId, boxId uint) (resp []define.PrizeIdIndexName, err error) { //(req.ActiveEndTime-time.Now().Unix()))
	lRange := s.cache.GetCache().LRange(ctx, define.RedisLastPrize(fanId, boxId), 0, -1).Val()
	if len(lRange) == 0 {
		return []define.PrizeIdIndexName{}, errors.New("Last已经发完...")
	}
	for _, ele := range lRange {
		e := define.PrizeIdIndexName{}
		json.Unmarshal([]byte(ele), &e)
		resp = append(resp, e)
	}
	return
}
func (s *FanServiceImpl) getGlobalPrizeEleCache(ctx context.Context, fanId, boxId uint) (resp []define.PrizeIdIndexName, err error) {
	lRange := s.cache.GetCache().LRange(ctx, define.RedisGlobalPrize(fanId, boxId), 0, -1).Val()
	if len(lRange) == 0 {
		return []define.PrizeIdIndexName{}, errors.New("全局赏已经发完...")
	}
	for _, ele := range lRange {
		e := define.PrizeIdIndexName{}
		json.Unmarshal([]byte(ele), &e)
		resp = append(resp, e)
	}
	return
}

func (s *FanServiceImpl) DealWithFirstPrize(ctx context.Context, fanId, boxId uint, fanName string, Title string, current, targetLen int64) { //(req.ActiveEndTime-time.Now().Unix()))
	prizeStable, err := s.getFirstPrizeEleCache(ctx, fanId, boxId)
	if err != nil {
		return
	}
	//获取first奖出奖需求位置范围
	firstPos := s.getFirstPrizePositionCache(context.Background(), fanId, boxId)
	if len(firstPos) != 2 {
		return
	}
	if current >= int64(firstPos[1]) {
		firstTarget := []int{0}
		firstBegin, firstEnd := 0, 0
		firstBegin = firstPos[0]
		firstEnd = firstPos[1]
		firstTarget = define.GetRandRums(firstBegin, firstEnd, len(prizeStable))
		s.setSpecialPositionRecord(ctx, fanId, boxId, firstTarget)
		tmpPrizeIndexName := ""
		fmt.Println("**************************随机位置坐标:", prizeStable, "*****************************")
		tx := s.db.GetDb().Begin()
		for i := 0; i < len(prizeStable); i++ {
			tmpPrizeIndexName = prizeStable[i].PrizeIndexName
			order := db.Order{
				ID:              randId.RandID(),
				FanId:           fanId,
				FanName:         fanName,
				BoxId:           boxId,
				BoxTitle:        Title,
				Position:        fmt.Sprintf("出奖规则:前[%d-%d]用户", firstPos[0], firstPos[1]),
				PrizeIndex:      uint(prizeStable[i].PrizeIndex),
				PrizeName:       prizeStable[i].PrizeName,
				PrizeIndexName:  prizeStable[i].PrizeIndexName,
				UserId:          111111111,
				UserName:        "test111111",
				FirstLastGlobal: firstTarget[i],
				Num:             1,
			}
			err = tx.Create(&order).Error
			if err != nil {
				tx.Rollback()
				return
			}
			s.cache.GetCache().LTrim(ctx, define.RedisFirstPrize(fanId, boxId), 1, -1).Err()
		}
		err = tx.Model(&db.Prize{}).Where("fan_id=? and box_id=? and prize_index_name=?", fanId, boxId, tmpPrizeIndexName).
			Update("prize_num", 0).Update("pricze_left_num", 0).Error
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}
}
func (s *FanServiceImpl) DealWithLastPrize(ctx context.Context, fanId, boxId uint, fanName string, Title string, current, targetLen int64) { //(req.ActiveEndTime-time.Now().Unix()))
	prizeStable, err := s.getLastPrizeEleCache(ctx, fanId, boxId)
	if err != nil {
		return
	}
	//获取last奖出奖需求位置范围
	lastPos := s.getLastPrizePositionCache(context.Background(), fanId, boxId)
	if len(lastPos) != 2 {
		return
	}
	begin := targetLen - int64(lastPos[1])
	end := targetLen - int64(lastPos[0])
	fmt.Println("current:", current, "******", "begin:", begin, "******", "end:", end)
	if current >= begin && current <= end {
		rd := s.getSpecialPositionRecord(ctx, fanId, boxId)
	again:
		lastTarget := define.GetRandRums(int(begin), int(end), len(prizeStable))
		for _, a := range rd {
			for _, b := range lastTarget {
				if a == b {
					goto again
				}
			}
		}
		s.setSpecialPositionRecord(ctx, fanId, boxId, lastTarget)
		tx := s.db.GetDb().Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		tmpPrizeIndexName := ""
		for i := 0; i < len(prizeStable); i++ {
			tmpPrizeIndexName = prizeStable[i].PrizeIndexName
			order := db.Order{
				ID:              randId.RandID(),
				FanId:           fanId,
				FanName:         fanName,
				BoxId:           boxId,
				BoxTitle:        Title,
				Position:        fmt.Sprintf("出奖规则:后[%d]用户]", lastPos[1]),
				PrizeIndex:      uint(prizeStable[i].PrizeIndex),
				PrizeName:       prizeStable[i].PrizeName,
				PrizeIndexName:  prizeStable[i].PrizeIndexName,
				UserId:          111111111,
				UserName:        "test111111",
				FirstLastGlobal: lastTarget[i],
				Num:             1,
			}
			err = tx.Create(&order).Error
			if err != nil {
				tx.Rollback()
				return
			}
			s.cache.GetCache().LTrim(ctx, define.RedisLastPrize(fanId, boxId), 1, -1).Err()
		}
		err = tx.Model(&db.Prize{}).Where("fan_id=? and box_id=? and prize_index_name=?", fanId, boxId, tmpPrizeIndexName).
			Update("prize_num", 0).Update("pricze_left_num", 0).Error
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}

}
func (s *FanServiceImpl) DealWithGlobalPrize(ctx context.Context, fanId, boxId uint, fanName string, Title string, current, targetLen int64) { //(req.ActiveEndTime-time.Now().Unix()))
	prizeStable, err := s.getGlobalPrizeEleCache(ctx, fanId, boxId)
	if err != nil {
		return
	}
	////获取global奖出奖需求位置范围
	//globalPos := s.getGlobalPrizePositionCache(context.Background(), fanId, boxId)
	//if len(globalPos) != 2 {
	//	return
	//}
	if current == targetLen {
		rd := s.getSpecialPositionRecord(ctx, fanId, boxId)
	again:
		globalTarget := define.GetRandRums(1, int(targetLen), len(prizeStable))
		for _, a := range rd {
			for _, b := range globalTarget {
				if a == b {
					goto again
				}
			}
		}
		s.setSpecialPositionRecord(ctx, fanId, boxId, globalTarget)
		tx := s.db.GetDb().Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		tmpPrizeIndexName := ""
		for i := 0; i < len(prizeStable); i++ {
			tmpPrizeIndexName = prizeStable[i].PrizeIndexName
			order := db.Order{
				ID:              randId.RandID(),
				FanId:           fanId,
				FanName:         fanName,
				BoxId:           boxId,
				BoxTitle:        Title,
				Position:        fmt.Sprintf("出奖规则:所有用户"),
				PrizeIndex:      uint(prizeStable[i].PrizeIndex),
				PrizeName:       prizeStable[i].PrizeName,
				PrizeIndexName:  prizeStable[i].PrizeIndexName,
				UserId:          111111111,
				UserName:        "test111111",
				FirstLastGlobal: globalTarget[i],
				Num:             1,
			}
			err = tx.Create(&order).Error
			if err != nil {
				tx.Rollback()
				return
			}
			s.cache.GetCache().LTrim(ctx, define.RedisGlobalPrize(fanId, boxId), 1, -1).Err()
		}
		err = tx.Model(&db.Prize{}).Where("fan_id=? and box_id=? and prize_index_name=?", fanId, boxId, tmpPrizeIndexName).
			Update("prize_num", 0).Update("pricze_left_num", 0).Error
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}
}
func (s *FanServiceImpl) SetRecordCache(ctx context.Context, pos int64, ele int, fanId, boxId uint) { //(req.ActiveEndTime-time.Now().Unix()))
	fanBoxfirstRecord := fmt.Sprintf("fanId%d-boxId%d-firstRecord", fanId, boxId)
	ret, _ := json.Marshal(struct {
		PrizeIndex int
		Pos        int64
	}{
		PrizeIndex: ele,
		Pos:        pos,
	})
	s.cache.GetCache().RPush(ctx, fanBoxfirstRecord, ret)
	//s.cache.GetCache().Expire(ctx, fanBoxfirstRecord, time.Second*time.Duration(seconds))
}
func (s *FanServiceImpl) IsRecordExist(ctx context.Context, pos int64, fanId, boxId uint, pIn define.PrizeIdIndexName) bool {
	fanBoxfirstRecord := fmt.Sprintf("fanId%d-boxId%d-firstRecord", fanId, boxId)
	lRange := s.cache.GetCache().LRange(ctx, fanBoxfirstRecord, 0, -1).Val()
	for _, ele := range lRange {
		xxx := struct {
			PrizeIndex int
			Pos        int64
		}{}
		json.Unmarshal([]byte(ele), &xxx)
		if xxx.Pos == pos && int32(xxx.PrizeIndex) == pIn.PrizeIndex {
			return true
		}
	}
	return false
}
func (s *FanServiceImpl) GetFirstPrizeCache(ctx context.Context, fanId, boxId uint) (rets []int) {
	fanBoxFirstPrize := fmt.Sprintf("fanId%d-boxId%d-firstPrize", fanId, boxId)
	lRange := s.cache.GetCache().LRange(ctx, fanBoxFirstPrize, 0, -1).Val()
	for _, ele := range lRange {
		n, _ := strconv.Atoi(ele)
		rets = append(rets, n)
	}
	return
}

func (s *FanServiceImpl) getFirstPrizePositionCache(ctx context.Context, fanId, boxId uint) (rets []int) {
	lRange := s.cache.GetCache().LRange(ctx, define.RedisFirstPrizePosition(fanId, boxId), 0, -1).Val()
	for _, ele := range lRange {
		n, _ := strconv.Atoi(ele)
		rets = append(rets, n)
	}
	return
}
func (s *FanServiceImpl) setSpecialPositionRecord(ctx context.Context, fanId, boxId uint, pos []int) {
	for i := 0; i < len(pos); i++ {
		s.cache.GetCache().RPush(ctx, define.RedisSpecealRecordPosition(fanId, boxId), pos[i])
	}
}
func (s *FanServiceImpl) getSpecialPositionRecord(ctx context.Context, fanId, boxId uint) (rets []int) {
	lRange := s.cache.GetCache().LRange(ctx, define.RedisSpecealRecordPosition(fanId, boxId), 0, -1).Val()
	for _, ele := range lRange {
		n, _ := strconv.Atoi(ele)
		rets = append(rets, n)
	}
	return
}
func (s *FanServiceImpl) getLastPrizePositionCache(ctx context.Context, fanId, boxId uint) (rets []int) {
	lRange := s.cache.GetCache().LRange(ctx, define.RedisLastPrizePosition(fanId, boxId), 0, -1).Val()
	for _, ele := range lRange {
		n, _ := strconv.Atoi(ele)
		rets = append(rets, n)
	}
	return
}

func (s *FanServiceImpl) getGlobalPrizePositionCache(ctx context.Context, fanId, boxId uint) (rets []int) {
	lRange := s.cache.GetCache().LRange(ctx, define.RedisGlobalPrizePosition(fanId, boxId), 0, -1).Val()
	for _, ele := range lRange {
		n, _ := strconv.Atoi(ele)
		rets = append(rets, n)
	}
	return
}
func (s *FanServiceImpl) getCurrentPrizeIndexCache(ctx context.Context, fanId, boxId uint) (int64, int64) { //(req.ActiveEndTime-time.Now().Unix()))
	targetLen := s.cache.GetCache().LLen(ctx, define.RedisTarget(fanId, boxId)).Val()
	leftLen := s.cache.GetCache().LLen(ctx, define.RedisLeft(fanId, boxId)).Val()
	sureLen := s.cache.GetCache().LLen(ctx, define.RedisSure(fanId, boxId)).Val()
	return targetLen - (leftLen + sureLen), targetLen
}

//client.LRem(ctx, fanBoxLeft, 1, val)

func (s *FanServiceImpl) getfirstRecordCache(ctx context.Context, index int64, fanId, boxId uint) []int { //(req.ActiveEndTime-time.Now().Unix()))
	fanBoxfirstRecord := fmt.Sprintf("fanId%d-boxId%d-firstRecord", fanId, boxId)
	//targetLen := s.cache.GetCache().LLen(ctx, fanBoxfirstRecord).Val()
	//val, _ := s.cache.GetCache().LIndex(ctx, fanBoxfirstRecord, index-1).Int()
	rets := make([]int, 0)
	lRange := s.cache.GetCache().LRange(ctx, fanBoxfirstRecord, 0, index-1).Val()
	for _, ele := range lRange {
		n, _ := strconv.Atoi(ele)
		rets = append(rets, n)
	}
	return rets
}

func (s *FanServiceImpl) makesure(req param.ReqBuySure, fanName string, boxTitle string, prizes []*db.TmpPrize) (resp param.RespBuySure, leftPrizes []*db.TmpPrize, err error) {
	prizeIndex := 0
	client := s.cache.GetCache()
	tx := s.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	ctx := context.Background()
	user := client.LLen(context.Background(), define.RedisUser(req.FanId, req.BoxId)).Val()
	Value, _ := client.LIndex(ctx, define.RedisTarget(req.FanId, req.BoxId), user-1).Int()
	if Value == 0 {
		//从left中抽取随机抽取
		Leftlen := client.LLen(ctx, define.RedisLeft(req.FanId, req.BoxId)).Val()
		rId := int64(0)
		if Leftlen != 0 {
			rId = define.RandPrizeIndex(Leftlen)
		} else {
			rId = 0
		}
		val, _ := client.LIndex(ctx, define.RedisLeft(req.FanId, req.BoxId), rId).Int()
		client.LRem(ctx, define.RedisLeft(req.FanId, req.BoxId), 1, val)
		prizeIndex = val
	} else {
		client.LRem(ctx, define.RedisSure(req.FanId, req.BoxId), 1, Value)
		prizeIndex = Value
	}
	tmpPrize := db.TmpPrize{}
	for _, ele := range prizes {
		if ele.PrizeIndex == int32(prizeIndex) {
			tmpPrize = *ele
			leftPrizes = append(leftPrizes, &db.TmpPrize{
				Id:             ele.Id,
				GoodID:         ele.GoodID,
				GoodName:       ele.GoodName,
				FanId:          ele.FanId,
				FanName:        ele.FanName,
				BoxID:          ele.BoxID,
				PrizeNum:       ele.PrizeNum,
				PriczeLeftNum:  ele.PriczeLeftNum - 1,
				PrizeIndex:     ele.PrizeIndex,
				PrizeIndexName: ele.PrizeIndexName,
			})
		} else {
			leftPrizes = append(leftPrizes, &db.TmpPrize{
				Id:             ele.Id,
				GoodID:         ele.GoodID,
				GoodName:       ele.GoodName,
				FanId:          ele.FanId,
				FanName:        ele.FanName,
				BoxID:          ele.BoxID,
				PrizeNum:       ele.PrizeNum,
				PriczeLeftNum:  ele.PriczeLeftNum,
				PrizeIndex:     ele.PrizeIndex,
				PrizeIndexName: ele.PrizeIndexName,
			})
		}

	}
	client.RPush(context.Background(), define.RedisUser(req.FanId, req.BoxId), "111").Result()
	err = tx.Model(&db.Prize{}).Where("fan_id=? and box_id=? and sold_status IN ? and prize_index_name<>? and prize_index_name<>? and prize_index_name<>? and prize_index=?",
		req.FanId, req.BoxId, []int{define.YfPrizeStatusNotSoldOut}, define.PrizeIndexNameFirst, define.PrizeIndexNameLast, define.PrizeIndexNameGlobal, prizeIndex).
		Update("pricze_left_num", tmpPrize.PriczeLeftNum-1).Error
	if err != nil {
		tx.Rollback()
		return param.RespBuySure{}, leftPrizes, errors.New("服务正忙...")
	}
	order := db.Order{
		ID:             randId.RandID(),
		FanId:          req.FanId,
		FanName:        fanName,
		BoxId:          req.BoxId,
		BoxTitle:       boxTitle,
		PrizeIndex:     uint(prizeIndex),
		PrizeName:      tmpPrize.GoodName,
		PrizeIndexName: tmpPrize.PrizeIndexName,
		UserId:         111111111,
		UserName:       "test111111",
		Num:            1,
		OrderType:      "orderType",
		PayStyle:       "payStyle",
		PayStatus:      "未支付",
	}
	err = tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return param.RespBuySure{}, leftPrizes, errors.New("服务正忙...")
	}
	tx.Commit()
	return param.RespBuySure{BuySure: param.BuySure{
		PrizeIndexName: order.PrizeIndexName,
		PrizeIndex:     prizeIndex,
		PrizeName:      order.PrizeName,
		Price:          order.Price,
		Pic:            order.Pic,
	}}, leftPrizes, nil
}

func (s *FanServiceImpl) BuyQuery(req param.ReqBuyQuery) (param.RespBuyQuerys, error) {
	results := param.RespBuyQuerys{}
	orders := []db.Order{}
	ret := s.db.GetDb().Order("created_at desc").Find(&orders)
	if ret.Error != nil {
		return param.RespBuyQuerys{}, errors.New("服务正忙...")
	}
	if ret.RowsAffected == 0 {
		return param.RespBuyQuerys{}, nil
	}
	for n, ele := range orders {
		results.RespBuyQuerys = append(results.RespBuyQuerys, param.RespBuyQuery{
			Index:          n + 1,
			PrizeIndexName: ele.PrizeIndexName,
			PrizeIndex:     int(ele.PrizeIndex),
			Price:          ele.Price,
			Pic:            ele.Pic,
			PrizeName:      ele.PrizeName,
		})
	}
	return results, nil
}
