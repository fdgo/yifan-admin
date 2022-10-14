package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
	"strings"
	"time"
	"yifan/app/api/param"
	"yifan/app/db"
	"yifan/pkg/define"
)

type PrizeIndexIdName struct {
	PrizeIndex     int32
	PrizeIndexName string
	PrizeNum       int32
	PrizeId        uint
	PrizeName      string
	Pos            []int
}

func (s *BoxServiceImpl) EachBox(box *param.Box, fanId uint, fanName string) ([]*db.Prize, int32) {
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

//func EachNormalPrize(prize *param.Prize) (prizeNum int32, prizeIndexIdName PrizeIndexIdName) {
//	prizeNum = prize.PrizeNum
//	prizeIndexIdName.PrizeIndex = prize.PrizeIndex
//	prizeIndexIdName.PrizeName = prize.PrizeName
//	prizeIndexIdName.PrizeId = prize.PrizeId
//	prizeIndexIdName.PrizeNum = prize.PrizeNum
//	return
//}
//func EachFirstPrize(prize *param.Prize) (prizeNum int32, prizeIndexIdName PrizeIndexIdName, eles []define.PrizeIdIndexName) {
//	prizeNum = prize.PrizeNum
//	prizeIndexIdName.PrizeIndex = prize.PrizeIndex
//	prizeIndexIdName.PrizeName = prize.PrizeName
//	prizeIndexIdName.PrizeId = prize.PrizeId
//	prizeIndexIdName.PrizeNum = prize.PrizeNum
//	return
//}
//func EachLastPrize(prize *param.Prize) (prizeNum int32, prizeIndexIdName PrizeIndexIdName, eles []define.PrizeIdIndexName) {
//	prizeNum = prize.PrizeNum
//	prizeIndexIdName.PrizeIndex = prize.PrizeIndex
//	prizeIndexIdName.PrizeName = prize.PrizeName
//	prizeIndexIdName.PrizeId = prize.PrizeId
//	prizeIndexIdName.PrizeNum = prize.PrizeNum
//	return
//}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *BoxServiceImpl) SetSureTargetEles(sures []PrizeIndexIdName) (int32, []int32, []int32) { //所有商品总数
	size := int32(0)
	sure := make([]int32, 0)
	for _, ele := range sures {
		size += ele.PrizeNum
		for _, e := range ele.Pos {
			if e != -1 {
				sure = append(sure, ele.PrizeIndex)
			}
		}
	}
	target := make([]int32, size)
	for _, ele := range sures {
		for _, e := range ele.Pos {
			if e != -1 {
				target[e-1] = ele.PrizeIndex
			} else {
				target = append(target, ele.PrizeIndex)
			}
		}
	}
	return size, target, sure
}
func (s *BoxServiceImpl) SetLeftEles(lefts []PrizeIndexIdName) []int32 { //(req.ActiveEndTime-time.Now().Unix()))
	left := make([]int32, 0)
	for _, ele := range lefts {
		times := int32(0)
		for _, e := range ele.Pos {
			if e != -1 {
				times = times + 1
			}
		}
		for i := int32(0); i < (ele.PrizeNum - times); i++ {
			left = append(left, ele.PrizeIndex)
		}
	}
	return left
}

func (s *BoxServiceImpl) SetFirstPrizeCache(pipe redis.Pipeliner, ctx context.Context, firstPrize []define.PrizeIdIndexName, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range firstPrize {
		e, _ := json.Marshal(ele)
		rPush := pipe.RPush(ctx, define.RedisFirstPrize(fanId, boxId), e)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(firstPrize))
	return
}

func (s *BoxServiceImpl) SetLastPrizeCache(pipe redis.Pipeliner, ctx context.Context, lastPrize []define.PrizeIdIndexName, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range lastPrize {
		e, _ := json.Marshal(ele)
		rPush := pipe.RPush(ctx, define.RedisLastPrize(fanId, boxId), e)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(lastPrize))
	return
}
func (s *BoxServiceImpl) SetGlobalPrizeCache(pipe redis.Pipeliner, ctx context.Context, globalPrize []define.PrizeIdIndexName, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range globalPrize {
		e, _ := json.Marshal(ele)
		rPush := pipe.RPush(ctx, define.RedisGlobalPrize(fanId, boxId), e)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(globalPrize))
	return
}
func (s *BoxServiceImpl) SetGlobalPrizePosition(pipe redis.Pipeliner, ctx context.Context, globalPrize []int, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range globalPrize {
		rPush := pipe.RPush(ctx, define.RedisGlobalPrizePosition(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(globalPrize))
	return
}
func (s *BoxServiceImpl) SetFirstPrizePosition(pipe redis.Pipeliner, ctx context.Context, firstPosition []int, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range firstPosition {
		rPush := pipe.RPush(ctx, define.RedisFirstPrizePosition(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(firstPosition))
	return
}
func (s *BoxServiceImpl) SetLastPrizePosition(pipe redis.Pipeliner, ctx context.Context, lastPosition []int, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range lastPosition {
		rPush := pipe.RPush(ctx, define.RedisLastPrizePosition(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(lastPosition))
	return
}

func (s *BoxServiceImpl) PkgBoxes(tx *gorm.DB, fanId uint, req param.ReqAddBox, boxIndex int32, prizeNum int32) (*db.Box, error) {
	box := &db.Box{
		ID:            define.GetRandBoxId(),
		FanId:         fanId,
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

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *BoxServiceImpl) SetSureCache(pipe redis.Pipeliner, ctx context.Context, sure []int32, fanId, boxId uint) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	pipe.Del(ctx, define.RedisSure(fanId, boxId))
	for _, ele := range sure {
		rPush := pipe.RPush(ctx, define.RedisSure(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(sure))
	return
}
func (s *BoxServiceImpl) SetLeftCache(pipe redis.Pipeliner, ctx context.Context, left []int32, fanId, boxId uint) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	pipe.Del(ctx, define.RedisLeft(fanId, boxId))
	for _, ele := range left {
		rPush := pipe.RPush(ctx, define.RedisLeft(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(left))
	return
}
func (s *BoxServiceImpl) SetTargetCache(pipe redis.Pipeliner, ctx context.Context, target []int32, fanId, boxId uint) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	pipe.Del(ctx, define.RedisTarget(fanId, boxId))
	for _, ele := range target {
		rPush := pipe.RPush(ctx, define.RedisTarget(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(target))
	return
}

func (s *BoxServiceImpl) SetUserCache(pipe redis.Pipeliner, ctx context.Context, userId int, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	rPush := pipe.RPush(ctx, define.RedisUser(fanId, boxId), userId)
	rCmd = append(rCmd, rPush)
	times = 1
	return
}

func (s *BoxServiceImpl) AddBox(req param.ReqAddBox) (param.RespAddBox, error) {
	var (
		fanIndex = 0
		fan      = &db.Fan{}
		ret      = param.RespAddBox{}
		tx       = s.db.GetDb().Begin()
	)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	result := tx.Model(&db.Fan{}).Where("title=?", req.Title).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return ret, errors.New("服务正忙...")
	}
	if result.RowsAffected != 0 {
		tx.Rollback()
		return ret, errors.New("此蕃已经存在...")
	}
	result = tx.Table("yf_fan").Select("fan_index").Order("fan_index desc").First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return ret, errors.New("服务正忙...")
	}
	if result.RowsAffected != 0 {
		fanIndex = fan.FanIndex
	}
	fanId := define.GetRandFanId()
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
		return ret, errors.New("服务正忙...")
	}
	all := 0
	boxIds := []uint{}
	for index := 0; index < req.BoxNum; index++ {
		boxEle := req.Boxes
		prizes, allPrizeNum := s.EachBox(&boxEle, fanId, req.Title)
		if allPrizeNum == 0 {
			tx.Rollback()
			return ret, errors.New("总奖品数为0...")
		}
		all = int(allPrizeNum)
		box, errx := s.PkgBoxes(tx, fanId, req, int32(index+1), allPrizeNum)
		if errx != nil {
			tx.Rollback()
			return ret, errors.New("服务正忙...")
		}
		boxIds = append(boxIds, box.ID)
		for nindex, ele := range prizes {
			prizes[nindex].BoxID = &box.ID
			prizes[nindex].BoxIndex = int(box.BoxIndex)
			prizes[nindex].Status = box.Status
			if ele.PrizeIndexName != define.PrizeIndexNameFirst &&
				ele.PrizeIndexName != define.PrizeIndexNameLast &&
				ele.PrizeIndexName != define.PrizeIndexNameGlobal {
				rate, _ := decimal.NewFromFloat32(float32(ele.PriczeLeftNum)).Div(decimal.NewFromFloat32(float32(allPrizeNum))).Float64()
				prizes[nindex].PrizeRate = fmt.Sprintf("%.2f", 100*rate) + "%"
			}
		}
		if err = tx.Model(&box).Association("Prizes").Append(&prizes); err != nil {
			tx.Rollback()
			return param.RespAddBox{}, errors.New("服务正忙...")
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
		return param.RespAddBox{Tips: tips1}, errors.New("位置必须2位...")
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
		return param.RespAddBox{Tips: tips2}, errors.New("位置范围错误...")
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
			return param.RespAddBox{}, errors.New("服务正忙...")
		}
	}
	firstPos, lastPos := []int{}, []int{}
	firstEles, lastEles, globalEles := []define.PrizeIdIndexName{}, []define.PrizeIdIndexName{}, []define.PrizeIdIndexName{}
	for _, onePrize := range req.Boxes.Prizes {
		if onePrize.PrizeIndexName == "First" {
			for p := 0; p < len(onePrize.Position); p++ {
				firstPos = append(firstPos, onePrize.Position[p])
			}
			for i := 0; i < int(onePrize.PrizeNum); i++ {
				firstEles = append(firstEles, define.PrizeIdIndexName{
					PrizeId:           onePrize.GoodId,
					PrizeIndex:        onePrize.PrizeIndex,
					PrizeName:         onePrize.GoodName,
					PrizeIndexName:    onePrize.PrizeIndexName,
					Price:             onePrize.Price,
					FanId:             fanId,
					FanName:           fan.Title,
					GoodID:            onePrize.GoodId,
					GoodName:          onePrize.GoodName,
					Remark:            onePrize.Remark,
					SoldStatus:        onePrize.SoldStatus,
					TimeForSoldStatus: onePrize.TimeForSoldStatus,
					PkgStatus:         onePrize.PkgStatus,
					Pic:               onePrize.Pic,
					IpID:              onePrize.IpId,
					IpName:            onePrize.IpName,
					SeriesID:          onePrize.SeriId,
					SeriesName:        onePrize.SeriName,
				})
			}
		} else if onePrize.PrizeIndexName == "Last" {
			for p := 0; p < len(onePrize.Position); p++ {
				lastPos = append(lastPos, onePrize.Position[p])
			}
			for i := 0; i < int(onePrize.PrizeNum); i++ {
				lastEles = append(lastEles, define.PrizeIdIndexName{
					PrizeId:           onePrize.GoodId,
					PrizeIndex:        onePrize.PrizeIndex,
					PrizeName:         onePrize.GoodName,
					PrizeIndexName:    onePrize.PrizeIndexName,
					Price:             onePrize.Price,
					FanId:             fanId,
					FanName:           fan.Title,
					GoodID:            onePrize.GoodId,
					GoodName:          onePrize.GoodName,
					Remark:            onePrize.Remark,
					SoldStatus:        onePrize.SoldStatus,
					TimeForSoldStatus: onePrize.TimeForSoldStatus,
					PkgStatus:         onePrize.PkgStatus,
					Pic:               onePrize.Pic,
					IpID:              onePrize.IpId,
					IpName:            onePrize.IpName,
					SeriesID:          onePrize.SeriId,
					SeriesName:        onePrize.SeriName,
				})
			}
		} else if onePrize.PrizeIndexName == "全局" {
			for i := 0; i < int(onePrize.PrizeNum); i++ {
				globalEles = append(globalEles, define.PrizeIdIndexName{
					PrizeId:           onePrize.GoodId,
					PrizeIndex:        onePrize.PrizeIndex,
					PrizeName:         onePrize.GoodName,
					PrizeIndexName:    onePrize.PrizeIndexName,
					Price:             onePrize.Price,
					FanId:             fanId,
					FanName:           fan.Title,
					GoodID:            onePrize.GoodId,
					GoodName:          onePrize.GoodName,
					Remark:            onePrize.Remark,
					SoldStatus:        onePrize.SoldStatus,
					TimeForSoldStatus: onePrize.TimeForSoldStatus,
					PkgStatus:         onePrize.PkgStatus,
					Pic:               onePrize.Pic,
					IpID:              onePrize.IpId,
					IpName:            onePrize.IpName,
					SeriesID:          onePrize.SeriId,
					SeriesName:        onePrize.SeriName,
				})
			}
		}
	}
	ctx := context.Background()
	for _, boxid := range boxIds {
		//var a, b, c, d, e, f []*redis.IntCmd
		//aTimes, bTimes, cTimes, dTimes, eTimes, fTimes := int64(0), int64(0), int64(0), int64(0), int64(0), int64(0)
		_, err = s.cache.Cache.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			s.SetFirstPrizePosition(pipe, ctx, firstPos, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
			s.SetLastPrizePosition(pipe, ctx, lastPos, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
			s.SetFirstPrizeCache(pipe, ctx, firstEles, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
			s.SetLastPrizeCache(pipe, ctx, lastEles, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
			s.SetGlobalPrizeCache(pipe, ctx, globalEles, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
			s.SetUserCache(pipe, ctx, -1, fanId, boxid, req.ActiveEndTime-time.Now().Unix())
			return nil
		})
		if err != nil {
			tx.Rollback()
			return ret, errors.New("服务正忙...")
		}
		//err = isAllResultOk(a, b, c, d, e, f, aTimes, bTimes, cTimes, dTimes, eTimes, fTimes)
		//if err != nil {
		//	tx.Rollback()
		//	return ret, errors.New("服务正忙...")
		//}
	}
	tx.Commit()
	return param.RespAddBox{}, nil
}
func (s *BoxServiceImpl) GetOneBoxAllNormalPrize(tx *gorm.DB, fanId, boxId uint) (allNorPrizes int, prizes []db.Prize, err error) {
	err = tx.Where("fan_id=? and box_id=? and prize_index_name not IN ?",
		fanId, boxId, []string{define.PrizeIndexNameFirst,
			define.PrizeIndexNameLast,
			define.PrizeIndexNameGlobal}).Find(&prizes).Error
	if err != nil {
		return 0, prizes, err
	}
	for _, onePrize := range prizes {
		allNorPrizes += int(onePrize.PrizeNum)
	}
	return allNorPrizes, prizes, nil
}

func (s *BoxServiceImpl) OnePrizeNeedToSetPosition(tx *gorm.DB, onePrize param.Prizex, fanId, boxId uint, oneBoxAllNorPrizeNum int) (num int64, err error) {
	tmpPosition := "["
	for _, onePos := range onePrize.Pos { //幾個位置 (3,7)
		if onePos > oneBoxAllNorPrizeNum {
			return 0, errors.New("位置超出范围...")
		}
		tmpPosition += fmt.Sprintf("%d,", onePos)
	}
	positon := strings.TrimRight(tmpPosition, ",")
	positon += "]"
	err = tx.Table("yf_prize").
		Where("fan_id=? and box_id=? and prize_index_name=? and prize_index=?",
			fanId, boxId, onePrize.PrizeIndexName, onePrize.PrizeIndex).
		Update("position", positon).Error
	if err != nil {
		return 0, errors.New("服务正忙...")
	}
	var prize db.Prize
	result := tx.Model(&db.Prize{}).Where("fan_id=? and box_id=? and prize_index_name=? and prize_index=?",
		fanId, boxId, onePrize.PrizeIndexName, onePrize.PrizeIndex).First(&prize)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return 0, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("奖品不存在...")
	}
	return int64(prize.PriczeLeftNum), nil
}
func (s *BoxServiceImpl) PageOfPosition(req param.ReqPageOfPosition) (param.RespPageOfPosition, error) {
	var fan db.Fan
	result := s.db.GetDb().Where("fan_index=?", req.FanIndex).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespPageOfPosition{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespPageOfPosition{}, errors.New("沒有任何蕃...")
	}
	var boxes []db.Box
	err := s.db.GetDb().Model(&fan).Where("status=? and created_at Between ? and ?",
		req.Status, time.Unix(req.TimeRange[0], 0), time.Unix(req.TimeRange[1], 0).Format("2006-01-02 15:04:05")).Association("Boxs").Find(&boxes)
	if err != nil {
		return param.RespPageOfPosition{}, errors.New("服务正忙...")
	}
	if boxes == nil {
		return param.RespPageOfPosition{}, errors.New("沒有任何箱子...")
	}
	res := param.RespPageOfPosition{}

	for _, oneBox := range boxes {
		var prizes []db.Prize
		s.db.GetDb().Model(&oneBox).Association("Prizes").Find(&prizes)
		for _, onePrize := range prizes {
			tmpPosition := "["
			for _, p := range onePrize.Position {
				tmpPosition += fmt.Sprintf("%d,", p)
			}
			positon := strings.TrimRight(tmpPosition, ",")
			positon += "]"

			ele := param.Ele{
				FanId:          fan.ID,
				BoxId:          oneBox.ID,
				Num:            onePrize.PrizeNum,
				PrizeIndexName: onePrize.PrizeIndexName,
				PrizeName:      onePrize.GoodName,
				Status:         onePrize.Status,
				Postion:        positon,
			}
			res.Ele = append(res.Ele, ele)
		}
	}
	return res, nil
}

func (s *BoxServiceImpl) PageOfPositionCondition(req param.ReqPageOfPositionCondition) (param.RespPageOfPositionCondition, error) {
	if req.BoxIndex == -1 {
		var fan db.Fan
		result := s.db.GetDb().Where("fan_index=?", req.FanIndex).First(&fan)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return param.RespPageOfPositionCondition{}, errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			return param.RespPageOfPositionCondition{}, errors.New("沒有任何蕃...")
		}
		var boxes []db.Box
		err := s.db.GetDb().Model(&fan).Where("status=? and created_at Between ? and ? ", req.Status, time.Unix(req.TimeRange[0], 0).Format("2006-01-02 15:04:05"),
			time.Unix(req.TimeRange[1], 0).Format("2006-01-02 15:04:05")).Association("Boxs").Find(&boxes)
		if err != nil {
			return param.RespPageOfPositionCondition{}, errors.New("服务正忙...")
		}
		if boxes == nil {
			return param.RespPageOfPositionCondition{}, errors.New("沒有任何箱子...")
		}
		res := param.RespPageOfPositionCondition{}
		for _, oneBox := range boxes {
			var prizes []db.Prize
			sql := ""
			value := []interface{}{}
			if req.PrizeIndexName != "All" {
				sql += "prize_index_name=? "
				value = append(value, req.PrizeIndexName)
				if req.PrizeName != "All" {
					sql += " and good_name=? "
					value = append(value, req.PrizeName)
				}
				err = s.db.GetDb().Model(&oneBox).Where(sql, value).Association("Prizes").Find(&prizes)
			} else {
				if req.PrizeName != "All" {
					sql += "good_name=?"
					value = append(value, req.PrizeName)
					err = s.db.GetDb().Model(&oneBox).Where(sql, value).Association("Prizes").Find(&prizes)
				} else {
					err = s.db.GetDb().Model(&oneBox).Association("Prizes").Find(&prizes)
				}
			}
			if err != nil {
				return param.RespPageOfPositionCondition{}, errors.New("服务正忙...")
			}
			for _, onePrize := range prizes {
				tmpPosition := "["
				for _, p := range onePrize.Position {
					tmpPosition += fmt.Sprintf("%d,", p)
				}
				positon := strings.TrimRight(tmpPosition, ",")
				positon += "]"
				ele := param.Ele{
					FanId:          fan.ID,
					BoxId:          oneBox.ID,
					Num:            onePrize.PrizeNum,
					PrizeIndexName: onePrize.PrizeIndexName,
					PrizeName:      onePrize.GoodName,
					Status:         onePrize.Status,
					Postion:        positon,
				}
				res.Ele = append(res.Ele, ele)
			}
		}
		return res, nil
	} else {
		var fan db.Fan
		result := s.db.GetDb().Where("fan_index=?", req.FanIndex).First(&fan)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return param.RespPageOfPositionCondition{}, errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			return param.RespPageOfPositionCondition{}, errors.New("沒有任何蕃...")
		}
		var box db.Box
		err := s.db.GetDb().Model(&fan).Where("box_index=? and status=? and created_at Between ? and ?", req.BoxIndex, req.Status, time.Unix(req.TimeRange[0], 0).Format("2006-01-02 15:04:05"),
			time.Unix(req.TimeRange[1], 0).Format("2006-01-02 15:04:05")).Association("Boxs").Find(&box)
		if err != nil {
			return param.RespPageOfPositionCondition{}, errors.New("服务正忙...")
		}
		res := param.RespPageOfPositionCondition{}
		sql := ""
		value := []interface{}{}
		var prizes []db.Prize
		if req.PrizeIndexName != "All" {
			sql += "prize_index_name=? "
			value = append(value, req.PrizeIndexName)
			if req.PrizeName != "All" {
				sql += " and good_name=? "
				value = append(value, req.PrizeName)
			}
			err = s.db.GetDb().Model(&box).Where(sql, value).Association("Prizes").Find(&prizes)
		} else {
			if req.PrizeName != "All" {
				sql += "good_name=?"
				value = append(value, req.PrizeName)
				err = s.db.GetDb().Model(&box).Where(sql, value).Association("Prizes").Find(&prizes)
			} else {
				err = s.db.GetDb().Model(&box).Association("Prizes").Find(&prizes)
			}
		}
		if err != nil {
			return param.RespPageOfPositionCondition{}, errors.New("服务正忙...")
		}
		for _, onePrize := range prizes {
			tmpPosition := "["
			for _, p := range onePrize.Position {
				tmpPosition += fmt.Sprintf("%d,", p)
			}
			positon := strings.TrimRight(tmpPosition, ",")
			positon += "]"

			ele := param.Ele{
				FanId:          fan.ID,
				BoxId:          box.ID,
				Num:            onePrize.PrizeNum,
				PrizeIndexName: onePrize.PrizeIndexName,
				PrizeName:      onePrize.GoodName,
				Status:         onePrize.Status,
				Postion:        positon,
			}
			res.Ele = append(res.Ele, ele)
		}
		return res, nil
	}
}
func (s *BoxServiceImpl) SetNormalPrizePosition(req param.ReqSetNormalPrizePosition) error {
	tx := s.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	ctx := context.Background()
	for _, oneFan := range req.Fanex {
		for _, oneBox := range oneFan.Boxex {
			allNorPrizes, prizes, err := s.GetOneBoxAllNormalPrize(tx, oneFan.FanId, oneBox.BoxId)
			if err != nil {
				tx.Rollback()
				return err
			}
			target := make([]int32, allNorPrizes)
			sureIndexs := make([]int32, 0)
			leftIndexs := make([]int32, 0)
			oneBoxSamePos := []int{}
			for _, onePrize := range oneBox.Prizex {
				num, errx := s.OnePrizeNeedToSetPosition(tx, onePrize, oneFan.FanId, oneBox.BoxId, allNorPrizes)
				if errx != nil {
					tx.Rollback()
					return errx
				}
				onePrizeSamePos := []int{}
				for _, p := range onePrize.Pos {
					onePrizeSamePos = append(onePrizeSamePos, p)
					oneBoxSamePos = append(oneBoxSamePos, p)
					target[p-1] = int32(onePrize.PrizeIndex)
					sureIndexs = append(sureIndexs, int32(onePrize.PrizeIndex))
				}
				if len(onePrizeSamePos) > int(num) {
					tx.Rollback()
					return errors.New("奖品位置数量超过了奖品数量...")
				}
				isOnePrizePossame := define.IsHasSameEle(onePrizeSamePos)
				if isOnePrizePossame {
					tx.Rollback()
					return errors.New("同一奖品位置不要重复...")
				}
			}
			isOneBoxPrizePossame := define.IsHasSameEle(oneBoxSamePos)
			if isOneBoxPrizePossame {
				tx.Rollback()
				return errors.New("同一箱奖品位置不要重复...")
			}
			for index, wOnePrize := range prizes {
				for _, ele := range sureIndexs {
					if wOnePrize.PrizeIndex == int32(ele) {
						prizes[index].PriczeLeftNum = prizes[index].PriczeLeftNum - 1
					}
				}
			}
			for _, ele := range prizes {
				for a := 0; a < int(ele.PriczeLeftNum); a++ {
					leftIndexs = append(leftIndexs, ele.PrizeIndex)
				}
			}
			var a, b, c []*redis.IntCmd
			aTimes, bTimes, cTimes := int64(0), int64(0), int64(0)
			_, err = s.cache.Cache.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				a, aTimes = s.SetSureCache(pipe, ctx, sureIndexs, oneFan.FanId, oneBox.BoxId)
				b, bTimes = s.SetLeftCache(pipe, ctx, leftIndexs, oneFan.FanId, oneBox.BoxId)
				c, cTimes = s.SetTargetCache(pipe, ctx, target, oneFan.FanId, oneBox.BoxId)
				return nil
			})
			if err != nil {
				tx.Rollback()
				return err
			}
			err = isAllResultOkEx(a, b, c, aTimes, bTimes, cTimes)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}
func isAllResultOkEx(a, b, c []*redis.IntCmd, aTimes, bTimes, cTimes int64) error {
	isAOk := false
	for _, ele := range a {
		if ele.Val() == aTimes {
			isAOk = true
		}
	}
	if !isAOk {
		return errors.New("请不要重复设置...")
	}

	isBOk := false
	for _, ele := range b {
		if ele.Val() == bTimes {
			isBOk = true
		}
	}
	if !isBOk {
		return errors.New("请不要重复设置...")
	}

	isCOk := false
	for _, ele := range c {
		if ele.Val() == cTimes {
			isCOk = true
		}
	}
	if !isCOk {
		return errors.New("请不要重复设置...")
	}
	return nil
}
func isAllResultOk(a, b, c, d, e, f []*redis.IntCmd, aTimes, bTimes, cTimes, dTimes, eTimes, fTimes int64) error {
	isAOk := false
	for _, ele := range a {
		if ele.Val() == aTimes {
			isAOk = true
		}
	}
	if !isAOk {
		return errors.New("服务正忙...")
	}

	isBOk := false
	for _, ele := range b {
		if ele.Val() == bTimes {
			isBOk = true
		}
	}
	if !isBOk {
		return errors.New("服务正忙...")
	}

	isCOk := false
	for _, ele := range c {
		if ele.Val() == cTimes {
			isCOk = true
		}
	}
	if !isCOk {
		return errors.New("服务正忙...")
	}

	isDOk := false
	for _, ele := range d {
		if ele.Val() == dTimes {
			isDOk = true
		}
	}
	if !isDOk {
		return errors.New("服务正忙...")
	}

	isEOk := false
	for _, ele := range e {
		if ele.Val() == eTimes {
			isEOk = true
		}
	}
	if !isEOk {
		return errors.New("服务正忙...")
	}

	isFOk := false
	for _, ele := range f {
		if ele.Val() == fTimes {
			isFOk = true
		}
	}
	if !isFOk {
		return errors.New("服务正忙...")
	}

	return nil
}
func (s *BoxServiceImpl) DeleteBox(req param.ReqDeleteBox) error {
	return nil
}

func (s *BoxServiceImpl) EachPrizeInfo(fanId uint, box *db.Box) (prizes []db.Prize, err error) { //
	err = s.db.GetDb().Model(box).Where("fan_id=? and  box_id=?",
		fanId, box.ID).Association("Prizes").Find(&prizes)
	return
}
func (s *BoxServiceImpl) EachPrizeInfoByStatus(fanId uint, box *db.Box, status ...int) (prizes []db.Prize, err error) { //
	tmpStatus := []int{0}
	for _, e := range status {
		tmpStatus = append(tmpStatus, e)
	}
	err = s.db.GetDb().Model(box).Where("fan_id=? and  box_id=? and sold_status IN ?",
		fanId, box.ID, tmpStatus).Association("Prizes").Find(&prizes)
	return
}
func (s *BoxServiceImpl) GetEachBoxAllPrizes(DB *gorm.DB, boxId, boxIndex uint) (int, int) {
	return 0, 0
}
func (s *BoxServiceImpl) EachBoxInfoByStatus(fanId uint, status ...int) (boxes []db.Box, result *gorm.DB) { //
	tmpStatus := []int{0}
	for _, e := range status {
		tmpStatus = append(tmpStatus, e)
	}
	result = s.db.GetDb().Where("fan_id=? and status IN ?", fanId, tmpStatus).Find(&boxes)
	return
}
func (s *BoxServiceImpl) ModifyBoxStatus(req param.ReqModifyBoxStatus) error {
	if req.Status != define.YfBoxStatusPrizeLeft && req.Status != define.YfBoxStatusNotOnBoard {
		return errors.New("状态输入有误...")
	}
	DB := s.db.GetDb()
	box := &db.Box{}
	result := DB.Where("id=?", req.BoxId).First(&box)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return errors.New("此箱不存在...")
	}
	result = DB.Model(&box).Update("status", req.Status)
	if result.Error != nil {
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return errors.New("箱子状态更新失败...")
	}
	return nil
}
func (s *BoxServiceImpl) QueryGoodsForBox(req param.ReqQueryGoodsForBox) (param.RespQueryGoodsForBox, error) {
	fmt.Println("add prize")
	var goods []db.Goods
	DB := s.db.GetDb()
	total := int64(0)
	err := DB.Model(&db.Goods{}).Count(&total).Error
	if err = DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&goods).Error; err != nil {
		return param.RespQueryGoodsForBox{}, errors.New("服务正忙...")
	}
	var resp param.RespQueryGoodsForBox
	for _, one := range goods {
		resp.GInfo.Good = append(resp.GInfo.Good, param.Good{
			ID:           one.ID,
			IpID:         one.IpID,
			IpName:       one.IpName,
			SeriesID:     one.SeriesID,
			SeriesName:   one.SeriesName,
			Pic:          one.Pic,
			Price:        one.Price,
			Name:         one.Name,
			SingleOrMuti: one.SingleOrMuti,
			MultiIds:     one.MultiIds,
			PkgStatus:    one.PkgStatus,
			Introduce:    one.Introduce,
			Integral:     one.Integral,
			SoldStatus:   one.SoldStatus,
		})
	}
	resp.GInfo.Num = len(resp.GInfo.Good)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}
func (s *BoxServiceImpl) GoodsToBePrize(req param.ReqGoodsToBePrize) error {
	return nil
}
func (s *BoxServiceImpl) ModifyBoxGoods(req param.ReqModifyBoxGoods) error {
	//var prize db.Prize
	//result := s.db.GetDb().Where("fan_id=? and good_id=?", req.FanId, req.OldGoodId).First(&prize)
	//if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
	//	return errors.New("服务正忙...")
	//}
	//if result.RowsAffected == 0 {
	//	return errors.New("不存在该奖品...")
	//}
	//m := make(map[string]interface{})
	//m["good_id"] = req.NewGoodId
	//m["good_name"] = req.NewGoodName
	//m["prize_num"] = req.NewPrizeNum
	//m["prize_index"] = req.NewPrizeIndex
	//m["prize_index_name"] = req.NewPrizeIndexName
	//m["pkg_status"] = req.NewPkgStatus
	//m["remark"] = req.NewRemark
	//m["position"] = req.NewPrizePosition
	//err := s.db.GetDb().Table("yf_prize").Where("fan_id=? and good_id=?", req.FanId, req.OldGoodId).
	//	Updates(&m).Error
	//if err != nil {
	//	return err
	//}
	return nil
}
func (s *BoxServiceImpl) DeleteBoxGoods(req param.ReqDeleteBoxGoods) error {
	err := s.db.GetDb().Model(&db.Prize{}).
		Where("fan_id=? and good_id=?", req.FanId, req.GoodId).Delete(&db.Prize{}).Error
	if err != nil {
		return errors.New("服务正忙...")
	}
	return nil
}
