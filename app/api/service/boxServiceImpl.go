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

func (s *BoxServiceImpl) EachBox(box *param.Box, fanId uint, fanName string) ([]db.Prize, []db.PrizePosition,
	int32, int32, int32, int32, []define.PrizeIdIndexName, []define.PrizeIdIndexName, []define.PrizeIdIndexName,
	[]int, []int, []int, []PrizeIndexIdName) {
	oneBoxFirstPrizeNum := int32(0)
	oneBoxLastPrizeNum := int32(0)
	oneBoxGlobalPrizeNum := int32(0)
	oneBoxNormalPrizeNum := int32(0)

	firstEles := []define.PrizeIdIndexName{}
	lastEles := []define.PrizeIdIndexName{}
	globalEles := []define.PrizeIdIndexName{}

	firstPos := make([]int, 0)
	lastPos := make([]int, 0)
	globalPos := make([]int, 0)
	oneBoxFirstPrizeIndexIdName := []PrizeIndexIdName{}
	oneBoxLastPrizeIndexIdName := []PrizeIndexIdName{}
	oneBoxGlobalPrizeIndexIdName := []PrizeIndexIdName{}
	oneBoxNormalPrizeIndexIdName := []PrizeIndexIdName{}
	prizes := []db.Prize{}
	prizePositions := []db.PrizePosition{}
	for _, prizeEle := range box.Prizes {
		if prizeEle.PrizeIndexName == define.PrizeIndexNameFirst {
			prizeNum := int32(0)
			indexIdName := PrizeIndexIdName{}
			prizeNum, indexIdName, firstEles, firstPos = EachFirstPrize(&prizeEle)
			oneBoxFirstPrizeNum += prizeNum
			oneBoxFirstPrizeIndexIdName = append(oneBoxFirstPrizeIndexIdName, indexIdName)
		} else if prizeEle.PrizeIndexName == define.PrizeIndexNameLast {
			prizeNum := int32(0)
			indexIdName := PrizeIndexIdName{}
			prizeNum, indexIdName, lastEles, lastPos = EachLastPrize(&prizeEle)
			oneBoxLastPrizeNum += prizeNum
			oneBoxLastPrizeIndexIdName = append(oneBoxLastPrizeIndexIdName, indexIdName)
		} else if prizeEle.PrizeIndexName == define.PrizeIndexNameGlobal {
			prizeNum := int32(0)
			indexIdName := PrizeIndexIdName{}
			prizeNum, indexIdName, globalEles, globalPos = EachGlobalPrize(&prizeEle)
			oneBoxGlobalPrizeNum += prizeNum
			oneBoxGlobalPrizeIndexIdName = append(oneBoxGlobalPrizeIndexIdName, indexIdName)
		} else {
			prizeNum, indexIdName := EachNormalPrize(&prizeEle)
			oneBoxNormalPrizeNum += prizeNum
			oneBoxNormalPrizeIndexIdName = append(oneBoxNormalPrizeIndexIdName, indexIdName)
		}
		prizes = append(prizes, db.Prize{
			ID:             define.GetRandPrizeId(),
			GoodID:         prizeEle.PrizeId,
			GoodName:       prizeEle.PrizeName,
			FanId:          fanId,
			FanName:        fanName,
			PrizePositions: prizeEle.PrizePosition,
			IpID:           prizeEle.IpId,
			IpName:         prizeEle.IpName,
			SeriesID:       prizeEle.SeriId,
			SeriesName:     prizeEle.SeriName,
			Pic:            prizeEle.Pic,
			PrizeNum:       prizeEle.PrizeNum,
			PriczeLeftNum:  prizeEle.PrizeNum,
			PrizeIndex:     prizeEle.PrizeIndex,
			PrizeIndexName: prizeEle.PrizeIndexName,
			SingleOrMuti:   prizeEle.SingleOrMuti,
			Price:          prizeEle.Price,
			MultiIds:       prizeEle.MultiIds,
			PkgStatus:      prizeEle.PkgStatus,
			Remark:         prizeEle.Remark,
			SoldStatus:     define.YfPrizeStatusNotSoldOut,
		})
		prizePositions = append(prizePositions, db.PrizePosition{
			FanId:          fanId,
			FanName:        fanName,
			PrizeIndex:     prizeEle.PrizeIndex,
			PrizeIndexName: prizeEle.PrizeIndexName,
			GoodName:       prizeEle.PrizeName,
			GoodId:         prizeEle.PrizeId,
			Position:       prizeEle.PrizePosition,
		})
	}
	return prizes, prizePositions, oneBoxFirstPrizeNum, oneBoxLastPrizeNum, oneBoxGlobalPrizeNum, oneBoxNormalPrizeNum,
		firstEles, lastEles, globalEles, firstPos, lastPos, globalPos, oneBoxNormalPrizeIndexIdName
}
func EachNormalPrize(prize *param.Prize) (prizeNum int32, prizeIndexIdName PrizeIndexIdName) {
	for _, posNorEle := range prize.PrizePosition {
		prizeIndexIdName.Pos = append(prizeIndexIdName.Pos, posNorEle)
	}
	prizeNum = prize.PrizeNum
	prizeIndexIdName.PrizeIndex = prize.PrizeIndex
	prizeIndexIdName.PrizeName = prize.PrizeName
	prizeIndexIdName.PrizeId = prize.PrizeId
	prizeIndexIdName.PrizeNum = prize.PrizeNum
	return
}
func EachFirstPrize(prize *param.Prize) (prizeNum int32, prizeIndexIdName PrizeIndexIdName, eles []define.PrizeIdIndexName, pos []int) {
	for _, posSpeEle := range prize.PrizePosition {
		prizeIndexIdName.Pos = append(prizeIndexIdName.Pos, posSpeEle)
		pos = append(pos, posSpeEle)
	}
	for i := 0; i < int(prize.PrizeNum); i++ {
		eles = append(eles, define.PrizeIdIndexName{
			PrizeId:        prize.PrizeId,
			PrizeIndex:     prize.PrizeIndex,
			PrizeName:      prize.PrizeName,
			PrizeIndexName: prize.PrizeIndexName,
			Price:          prize.Price,
			Pic:            prize.Pic,
		})
	}
	prizeNum = prize.PrizeNum
	prizeIndexIdName.PrizeIndex = prize.PrizeIndex
	prizeIndexIdName.PrizeName = prize.PrizeName
	prizeIndexIdName.PrizeId = prize.PrizeId
	prizeIndexIdName.PrizeNum = prize.PrizeNum
	return
}
func EachLastPrize(prize *param.Prize) (prizeNum int32, prizeIndexIdName PrizeIndexIdName, eles []define.PrizeIdIndexName, pos []int) {
	for _, posSpeEle := range prize.PrizePosition {
		prizeIndexIdName.Pos = append(prizeIndexIdName.Pos, posSpeEle)
		pos = append(pos, posSpeEle)
	}
	for i := 0; i < int(prize.PrizeNum); i++ {
		eles = append(eles, define.PrizeIdIndexName{
			PrizeId:        prize.PrizeId,
			PrizeIndex:     prize.PrizeIndex,
			PrizeName:      prize.PrizeName,
			PrizeIndexName: prize.PrizeIndexName,
			Price:          prize.Price,
			Pic:            prize.Pic,
		})
	}
	prizeNum = prize.PrizeNum
	prizeIndexIdName.PrizeIndex = prize.PrizeIndex
	prizeIndexIdName.PrizeName = prize.PrizeName
	prizeIndexIdName.PrizeId = prize.PrizeId
	prizeIndexIdName.PrizeNum = prize.PrizeNum
	return
}
func EachGlobalPrize(prize *param.Prize) (prizeNum int32, prizeIndexIdName PrizeIndexIdName, eles []define.PrizeIdIndexName, pos []int) {
	for _, posSpeEle := range prize.PrizePosition {
		prizeIndexIdName.Pos = append(prizeIndexIdName.Pos, posSpeEle)
		pos = append(pos, posSpeEle)
	}
	for i := 0; i < int(prize.PrizeNum); i++ {
		eles = append(eles, define.PrizeIdIndexName{
			PrizeId:        prize.PrizeId,
			PrizeIndex:     prize.PrizeIndex,
			PrizeName:      prize.PrizeName,
			PrizeIndexName: prize.PrizeIndexName,
			Price:          prize.Price,
			Pic:            prize.Pic,
		})
	}
	prizeNum = prize.PrizeNum
	prizeIndexIdName.PrizeIndex = prize.PrizeIndex
	prizeIndexIdName.PrizeName = prize.PrizeName
	prizeIndexIdName.PrizeId = prize.PrizeId
	prizeIndexIdName.PrizeNum = prize.PrizeNum
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *BoxServiceImpl) SetSureTargetEles(sures []PrizeIndexIdName) (int32, []int32, []int32) { //所有商品总数
	size := int32(0)
	sure := make([]int32, 0)
	for _, ele := range sures {
		size += ele.PrizeNum
		for _, e := range ele.Pos {
			if e != 0 {
				sure = append(sure, ele.PrizeIndex)
			}
		}
	}
	target := make([]int32, size)
	for _, ele := range sures {
		for _, e := range ele.Pos {
			if e != 0 {
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
			if e != 0 {
				times = times + 1
			}
		}
		for i := int32(0); i < (ele.PrizeNum - times); i++ {
			left = append(left, ele.PrizeIndex)
		}
	}
	return left
}

//func (s *BoxServiceImpl) SetfirstRecordCache(ctx context.Context, firstRecord []int, fanId, boxId uint, seconds int64) { //(req.ActiveEndTime-time.Now().Unix()))
//	fanBoxfirstRecord := fmt.Sprintf("fanId%d-boxId%d-firstRecord", fanId, boxId)
//	for _, ele := range firstRecord {
//		s.cache.GetCache().RPush(ctx, fanBoxfirstRecord, ele)
//	}
//	s.cache.GetCache().Expire(ctx, fanBoxfirstRecord, time.Second*time.Duration(seconds))
//}
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

func (s *BoxServiceImpl) PkgBoxes(tx *gorm.DB, fanId uint, req param.ReqAddBox, reqbox param.Box, boxIndex int32, prizeNum, prizeLeftNum int32) (*db.Box, error) {
	box := &db.Box{
		ID:            define.GetRandBoxId(),
		FanId:         fanId,
		FanName:       req.FanTitle,
		BoxIndex:      boxIndex,
		PriczeNum:     prizeNum,
		PriczeLeftNum: prizeLeftNum,
		Status:        define.YfBoxStatusNotOnBoard,
	}
	err := tx.Create(box).Error
	if err != nil {
		return nil, errors.New("箱子创建失败...")
	}
	return box, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *BoxServiceImpl) SetTargetCache(pipe redis.Pipeliner, ctx context.Context, target []int32, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range target {
		rPush := pipe.RPush(ctx, define.RedisTarget(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(target))
	return
}
func (s *BoxServiceImpl) SetSureCache(pipe redis.Pipeliner, ctx context.Context, sure []int32, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range sure {
		rPush := pipe.RPush(ctx, define.RedisSure(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(sure))
	return
}
func (s *BoxServiceImpl) SetLeftCache(pipe redis.Pipeliner, ctx context.Context, left []int32, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	for _, ele := range left {
		rPush := pipe.RPush(ctx, define.RedisLeft(fanId, boxId), ele)
		rCmd = append(rCmd, rPush)
	}
	times = int64(len(left))
	return
}
func (s *BoxServiceImpl) SetUserCache(pipe redis.Pipeliner, ctx context.Context, userId int, fanId, boxId uint, seconds int64) (rCmd []*redis.IntCmd, times int64) { //(req.ActiveEndTime-time.Now().Unix()))
	rPush := pipe.RPush(ctx, define.RedisUser(fanId, boxId), userId)
	rCmd = append(rCmd, rPush)
	times = 1
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *BoxServiceImpl) AddBox(req param.ReqAddBox) (param.RespAddBox, error) {
	var (
		fan = &db.Fan{}
		ret = param.RespAddBox{}
		tx  = s.db.GetDb().Begin()
	)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	result := tx.Where("title=?", req.FanTitle).First(&fan)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return ret, errors.New("服务正忙...")
	}
	if result.RowsAffected != 0 {
		tx.Rollback()
		return ret, errors.New("此蕃已经存在...")
	}
	fanId := define.GetRandFanId()
	tx.Create(&db.Fan{
		ID:        fanId,
		Title:     req.Title,
		Status:    define.YfFanStatusNotOnBoard,
		Price:     req.FanPrice,
		SharePic:  req.SharePic,
		DetailPic: req.DetailPic,
		Rule:      req.Rule,
	})
	ctx := context.Background()
	//for index, boxEle := range req.BoxNum { //fMix, lMix, gMix
	for index := 0; index < req.BoxNum; index++ {
		boxEle := req.Boxes
		prizes, prizePositions, fNum, lNum, gNum, nNum, firstEles, lastEles, globalEles, firstPos, lastPos, _, nMix :=
			s.EachBox(&boxEle, fanId, req.FanTitle)
		box, err := s.PkgBoxes(tx, fanId, req, boxEle, int32(index+1), fNum+lNum+gNum+nNum, fNum+lNum+gNum+nNum)
		if err != nil {
			tx.Rollback()
			return ret, err
		}
		allPrizeNums := int32(0)
		target := []int32{}
		sures := []int32{}
		var a, b, c, d, e, f, g, h, i, j []*redis.IntCmd
		aTimes, bTimes, cTimes, dTimes, eTimes, fTimes, gTimes, hTimes, iTimes, jTimes := int64(0), int64(0), int64(0), int64(0), int64(0), int64(0), int64(0), int64(0), int64(0), int64(0)
		_, err = s.cache.Cache.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			a, aTimes = s.SetFirstPrizePosition(pipe, ctx, firstPos, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			b, bTimes = s.SetLastPrizePosition(pipe, ctx, lastPos, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			//c, cTimes = s.SetGlobalPrizePosition(pipe, ctx, globalPos, req.FanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			d, dTimes = s.SetFirstPrizeCache(pipe, ctx, firstEles, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			e, eTimes = s.SetLastPrizeCache(pipe, ctx, lastEles, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			f, fTimes = s.SetGlobalPrizeCache(pipe, ctx, globalEles, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			allPrizeNums, target, sures = s.SetSureTargetEles(nMix)
			g, gTimes = s.SetTargetCache(pipe, ctx, target, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			h, hTimes = s.SetSureCache(pipe, ctx, sures, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			lefts := s.SetLeftEles(nMix)
			i, iTimes = s.SetLeftCache(pipe, ctx, lefts, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			j, jTimes = s.SetUserCache(pipe, ctx, -1, fanId, box.ID, req.ActiveEndTime-time.Now().Unix())
			return nil
		})
		if err != nil {
			tx.Rollback()
			return ret, errors.New("服务正忙...")
		}
		err = isAllResultOk(a, b, c, d, e, f, g, h, i, j, aTimes, bTimes, cTimes, dTimes, eTimes, fTimes, gTimes, hTimes, iTimes, jTimes)
		if err != nil {
			tx.Rollback()
			return ret, errors.New("服务正忙...")
		}
		for nindex, x := range prizes {
			prizes[nindex].BoxID = &box.ID
			rate, _ := decimal.NewFromFloat32(float32(x.PriczeLeftNum)).Div(decimal.NewFromFloat32(float32(allPrizeNums))).Float64()
			prizes[nindex].PrizeRate = fmt.Sprintf("%.2f", 100*rate) + "%"
		}
		for nindex, _ := range prizePositions {
			prizePositions[nindex].BoxID = &box.ID
		}
		if err = tx.Model(&box).Association("Prizes").Append(&prizes); err != nil {
			tx.Rollback()
			return param.RespAddBox{}, errors.New("服务正忙...")
		}
		if err = tx.Model(&box).Association("PrizePositions").Append(&prizePositions); err != nil {
			tx.Rollback()
			return param.RespAddBox{}, errors.New("服务正忙...")
		}
	}
	tx.Commit()
	return ret, nil
}
func isAllResultOk(a, b, c, d, e, f, g, h, i, j []*redis.IntCmd, aTimes, bTimes, cTimes, dTimes, eTimes, fTimes, gTimes, hTimes, iTimes, jTimes int64) error {
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

	//isCOk := false
	//for _, ele := range c {
	//	if ele.Val() == cTimes {
	//		isCOk = true
	//	}
	//}
	//if !isCOk {
	//	return errors.New("服务正忙...")
	//}

	isDOk := false
	for _, ele := range d {
		if ele.Val() == dTimes {
			isDOk = true
		}
	}
	if !isDOk {
		return errors.New("服务正忙...")
	}

	//////////////////
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

	isGOk := false
	for _, ele := range g {
		if ele.Val() == gTimes {
			isGOk = true
		}
	}
	if !isGOk {
		return errors.New("服务正忙...")
	}

	isHOk := false
	for _, ele := range h {
		if ele.Val() == hTimes {
			isHOk = true
		}
	}
	if !isHOk {
		return errors.New("服务正忙...")
	}

	isIOk := false
	for _, ele := range i {
		if ele.Val() == iTimes {
			isIOk = true
		}
	}
	if !isIOk {
		return errors.New("服务正忙...")
	}

	isJOk := false
	for _, ele := range j {
		if ele.Val() == jTimes {
			isJOk = true
		}
	}
	if !isJOk {
		return errors.New("服务正忙...")
	}

	return nil
}
func (s *BoxServiceImpl) DeleteBox(req param.ReqDeleteBox) error {
	fmt.Println("DeleteBox......")
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
	//resp := param.RespEnterFan{}
	//ret := DB.Where("id=? or status IN ? ", fanId, []int{define.YfBoxStatusPrizeLeft, define.YfBoxStatusPrizeNotLeft}).First(&db.Fan{})
	//if ret.Error != nil && ret.Error != gorm.ErrRecordNotFound {
	//	return 0, 0
	//}
	//if ret.RowsAffected == 0 {
	//	return 0, 0
	//}
	//boxAll, result := s.EachBoxInfoByStatus(fanId, define.YfBoxStatusPrizeLeft, define.YfBoxStatusPrizeNotLeft) //上架有商品，上架无商品
	//if result.RowsAffected == 0 {
	//	return 0, 0
	//}
	//boxLeft, result := s.EachBoxInfoByStatus(fanId, define.YfBoxStatusPrizeLeft) //上架有商品
	//if result.RowsAffected == 0 {
	//	return 0, 0
	//}
	//for _, eleBox := range boxLeft {
	//	prize, err := s.EachPrizeInfoByStatus(fanId, &eleBox, define.YfPrizeStatusNotSoldOut)
	//	if err != nil {
	//		return 0, 0
	//	}
	//	if len(prize) == 0 {
	//		continue
	//	}
	//	oneBoxPrizeNum := int32(0)
	//	for _, elePrize := range prize {
	//		oneBoxPrizeNum += elePrize.PriczeLeftNum
	//	}
	//}
}
func (s *BoxServiceImpl) EachBoxInfoByStatus(fanId uint, status ...int) (boxes []db.Box, result *gorm.DB) { //
	tmpStatus := []int{0}
	for _, e := range status {
		tmpStatus = append(tmpStatus, e)
	}
	result = s.db.GetDb().Where("fan_id=? and status IN ?", fanId, tmpStatus).Find(&boxes)
	return
}

func (s *BoxServiceImpl) ModifyBox(req param.ReqModifyBox) error {
	fmt.Println("EditBox......")
	return nil
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
	m := make(map[string]interface{})
	m["good_id"] = req.NewGoodId
	m["good_name"] = req.NewGoodName
	m["prize_num"] = req.NewPrizeNum
	m["prize_index"] = req.NewPrizeIndex
	m["prize_index_name"] = req.NewPrizeIndexName
	m["pkg_status"] = req.NewPkgStatus
	m["single_or_muti"] = req.NewSingleOrMuti
	m["multi_ids"] = req.NewMultiIds
	tx := s.db.GetDb()
	err := tx.Table("yf_prize").Where("fan_id=? and good_id=? and prize_index=?", req.FanId, req.OldGoodId, req.OldPrizeIndex).
		Updates(&m).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("yf_prize_position").Where("fan_id=? and good_id=?", req.FanId, req.OldGoodId).
		Update("position", req.NewPrizePosition).
		Update("good_name", req.NewGoodName).
		Update("good_id", req.NewGoodId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (s *BoxServiceImpl) DeleteBoxGoods(req param.ReqDeleteBoxGoods) error {
	fmt.Println("delete prize", req.GoodId, req.FanId)
	err := s.db.GetDb().Model(&db.Prize{}).
		Where("fan_id=? and good_id=?", req.FanId, req.GoodId).Delete(&db.Prize{}).Error
	if err != nil {
		return errors.New("服务正忙...")
	}
	return nil

}
