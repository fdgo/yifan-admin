package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"math"
	"yifan/app/api/param"
	"yifan/app/db"
)

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
		copier.Copy(&o, oneOrder)
		resp.Orders = append(resp.Orders, o)
	}
	resp.Num = len(resp.Orders)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}
func (s *OrderServiceImpl) PageOfOrderCondition(req param.ReqPageOfOrderCondition) (param.RespPageOfOrderCondition, error) {
	fmt.Println("222222222222")
	return param.RespPageOfOrderCondition{}, nil
}
