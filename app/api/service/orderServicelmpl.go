package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"math"
	"time"
	"yifan/app/api/param"
	"yifan/app/db"
)

func (s *OrderServiceImpl) AddRemark(req param.ReqAddRemark) error {
	return s.db.GetDb().Model(&db.Order{}).Where("out_trade_no=?", req.OrderId).Update("remark", req.Remark).Error
}
func (s *OrderServiceImpl) PageOfOrder(req param.ReqPageOfOrder) (param.RespPageOfOrder, error) {
	DB := s.db.GetDb()
	var order []db.Order
	var total int64
	var resp param.RespPageOfOrder
	err := DB.Model(&db.Order{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespPageOfOrder{}, errors.New("服务正忙......")
	}
	if err = DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&order).Error; err != nil {
		return param.RespPageOfOrder{}, errors.New("服务正忙...")
	}
	for _, oneOrder := range order {
		var o param.Order
		var lugg []db.Luggage
		DB.Where("out_trade_no=?", oneOrder.OutTradeNo).Find(&lugg)
		for _, oneLugg := range lugg {
			o.Goods = append(o.Goods, param.Goodxs{
				IpID:           oneLugg.IpID,
				IpName:         oneLugg.IpName,
				SeriesID:       oneLugg.SeriesID,
				SeriesName:     oneLugg.SeriesName,
				PrizeName:      oneLugg.GoodName,
				PrizeIndexName: oneLugg.PrizeIndexName,
				PrizeId:        oneLugg.GoodID,
				Pic:            oneLugg.Pic,
			})
		}
		copier.Copy(&o, oneOrder)
		o.CreateTime = TimetToInt64(oneOrder.CreatedAt)
		resp.Orders = append(resp.Orders, o)
	}
	resp.Num = len(resp.Orders)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}

func (s *OrderServiceImpl) PageOfOrderCondition(req param.ReqPageOfOrderCondition) (param.RespPageOfOrderCondition, error) {
	DB := s.db.GetDb()
	if req.OrderId != "" {
		DB = DB.Model(&db.Order{}).Where("out_trade_no=?", req.OrderId)
	}
	if req.OrderStatus != "" {
		DB = DB.Model(&db.Order{}).Where("status=?", req.OrderStatus)
	}
	if req.Mobile != "" {
		DB = DB.Model(&db.Order{}).Where("user_mobile=?", req.Mobile)
	}
	if req.UserId != 0 {
		DB = DB.Model(&db.Order{}).Where("user_id=?", req.UserId)
	}
	if req.PayStyle != "" {
		DB = DB.Model(&db.Order{}).Where("pay_style=?", req.PayStyle)
	}
	var order []db.Order
	var total int64
	var resp param.RespPageOfOrderCondition
	err := DB.Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespPageOfOrderCondition{}, errors.New("服务正忙......")
	}
	if err = DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&order).Error; err != nil {
		return param.RespPageOfOrderCondition{}, errors.New("服务正忙...")
	}
	for _, oneOrder := range order {
		var o param.Order
		var lugg []db.Luggage
		s.db.GetDb().Where("out_trade_no=?", oneOrder.OutTradeNo).Find(&lugg)
		for _, oneLugg := range lugg {
			o.Goods = append(o.Goods, param.Goodxs{
				IpID:           oneLugg.IpID,
				IpName:         oneLugg.IpName,
				SeriesID:       oneLugg.SeriesID,
				SeriesName:     oneLugg.SeriesName,
				PrizeName:      oneLugg.GoodName,
				PrizeIndexName: oneLugg.PrizeIndexName,
				PrizeId:        oneLugg.GoodID,
				Pic:            oneLugg.Pic,
			})
		}
		copier.Copy(&o, oneOrder)
		o.CreateTime = TimetToInt64(oneOrder.CreatedAt)
		resp.Orders = append(resp.Orders, o)
	}
	resp.Num = len(resp.Orders)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}
func TimetToInt64(tTime time.Time) (nTime int64) {
	return time.Date(tTime.Year(), tTime.Month(), tTime.Day(), tTime.Hour(),
		tTime.Minute(), tTime.Second(), 0, tTime.Location()).Unix()
}

func (s *OrderServiceImpl) PageOfOrderDetail(req param.ReqPageOfOrderDetail) (param.RespPageOfOrderDetail, error) {
	DB := s.db.GetDb()
	var orderSrc db.Order
	var resp param.RespPageOfOrderDetail
	if err := DB.Where("out_trade_no=?", req.OrderId).First(&orderSrc).Error; err != nil {
		return param.RespPageOfOrderDetail{}, errors.New("服务正忙...")
	}
	var lugSrc []db.Luggage
	DB.Where("out_trade_no=?", req.OrderId).Find(&lugSrc)
	for _, oneLugg := range lugSrc {
		resp.Orders.Goods = append(resp.Orders.Goods, param.Goodxs{
			IpID:           oneLugg.IpID,
			IpName:         oneLugg.IpName,
			SeriesID:       oneLugg.SeriesID,
			SeriesName:     oneLugg.SeriesName,
			PrizeName:      oneLugg.GoodName,
			PrizeIndexName: oneLugg.PrizeIndexName,
			PrizeId:        oneLugg.GoodID,
			Pic:            oneLugg.Pic,
		})
	}
	resp.Orders.FinishTime = orderSrc.FinishTime
	resp.Orders.UserMobile = orderSrc.UserMobile
	resp.Orders.UserId = orderSrc.UserId
	resp.Orders.UserName = orderSrc.UserName
	resp.Orders.PrizeNum = orderSrc.PrizeNum
	resp.Orders.OutTradeNo = orderSrc.OutTradeNo
	resp.Orders.Price = orderSrc.Price
	resp.Orders.PayStyle = orderSrc.PayStyle
	resp.Orders.TansactionId = orderSrc.TransactionId
	resp.Orders.CreateTime = TimetToInt64(orderSrc.CreatedAt)
	return resp, nil
}
