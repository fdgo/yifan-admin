package service

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"gorm.io/gorm"
	"math"
	"os"
	"strconv"
	"yifan/app/api/param"
	"yifan/app/db"
	"yifan/pkg/define"
)

func (s *SeriServiceImpl) UpLoadSeries(req param.ReqUpLoadSeries) (param.RespUpLoadSeries, error) {
	//获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return param.RespUpLoadSeries{}, err
	}
	xlsxPath := dir + "/import.xlsx"
	//打开文件路径
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return param.RespUpLoadSeries{}, err
	}
	DB := s.db.GetDb()
	type IpSer struct {
		IpId  uint
		SerId uint
	}
	for _, oneSheet := range xlsxFile.Sheets {
		if oneSheet.Name == "系列" {
			var ipSer []IpSer
			for index, row := range oneSheet.Rows {
				if index != 0 {
					ipx := db.Ip{}
					serId := uint(0)
					for n, cell := range row.Cells {
						if n == 0 {
							DB.Where("name=?", cell.Value).First(&ipx)
						}
						if n == 1 {
							serId = define.DealWithOneSeries(DB, ipx.Name, cell.Value)
						}
					}
					ipSer = append(ipSer, IpSer{
						IpId:  ipx.ID,
						SerId: serId,
					})
				}
			}
			for j, row := range oneSheet.Rows {
				if j == 0 {
					for k, _ := range row.Cells {
						if k+1 == len(row.Cells) {
							x := row.AddCell()
							x.Value = "IP对应ID"
							xlsxFile.Save(xlsxPath)
						}
					}
				}
			}
			for j, row := range oneSheet.Rows {
				if j == 0 {
					for k, _ := range row.Cells {
						if k+1 == len(row.Cells) {
							x := row.AddCell()
							x.Value = "系列对应ID"
							xlsxFile.Save(xlsxPath)
						}
					}
				}
			}
			for j, row := range oneSheet.Rows {
				if j != 0 {
					for k, _ := range row.Cells {
						if k+1 == len(row.Cells) {
							x := row.AddCell()
							x.Value = strconv.Itoa(int(ipSer[j-1].IpId))
							xlsxFile.Save(xlsxPath)
						}
					}
				}
			}
			for j, row := range oneSheet.Rows {
				if j != 0 {
					for k, _ := range row.Cells {
						if k+1 == len(row.Cells) {
							x := row.AddCell()
							x.Value = strconv.Itoa(int(ipSer[j-1].SerId))
							xlsxFile.Save(xlsxPath)
						}
					}
				}
			}
		}
	}
	return param.RespUpLoadSeries{}, nil
}

func (s *SeriServiceImpl) ManySerUpload(req param.ReqUpLoadSeries) (param.RespUpLoadSeries, error) {
	//ret := param.RespUpLoadSeries{}
	//DB := s.db.GetDb()
	//for _, ele := range req.ReqAddSeries {
	//	if *ele.IpId == 0 {
	//		ret.SeriIdNames = append(ret.SeriIdNames, param.SeriIdName{
	//			IpId: *ele.IpId,
	//			Name: ele.Name,
	//			Tip:  "缺少IP的id参数...",
	//		})
	//		continue
	//	}
	//	ip := &db.Ip{}
	//	result := DB.Where("id=?", ele.IpId).Find(&ip)
	//	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
	//		return param.RespUpLoadSeries{}, errors.New("服务正忙...")
	//	}
	//	if result.RowsAffected == 0 {
	//		ret.SeriIdNames = append(ret.SeriIdNames, param.SeriIdName{
	//			IpId: *ele.IpId,
	//			Name: ele.Name,
	//			Tip:  "库中不存在这样的IP...",
	//		})
	//		continue
	//	}
	//	ser := []db.Series{}
	//	rt := DB.Where("name=?", ele.Name).Find(&ser)
	//	if rt.Error != nil && rt.Error != gorm.ErrRecordNotFound {
	//		return param.RespUpLoadSeries{}, errors.New("服务正忙...")
	//	}
	//	if ser != nil { //存在该系列
	//		isOk := true
	//		for _, oneSer := range ser {
	//			if *oneSer.IpID == *ele.IpId && oneSer.Name == ele.Name {
	//				ret.SeriIdNames = append(ret.SeriIdNames, param.SeriIdName{
	//					IpId: *ele.IpId,
	//					Name: ele.Name,
	//					Tip:  "该系列已经创建...",
	//				})
	//				isOk = false
	//				break
	//			}
	//		}
	//		if !isOk {
	//			continue
	//		}
	//		tmpSer := &db.Series{ID: define.GetRandSeriesId(), Name: ele.Name, IpName: ele.IpName,
	//			CreateName: ele.CreateName, IpID: ele.IpId}
	//		err := DB.Create(tmpSer).Error
	//		if err != nil {
	//			return param.RespUpLoadSeries{}, errors.New("服务正忙...")
	//		}
	//	} else { //不存在该系列
	//		tmpSer := &db.Series{ID: define.GetRandSeriesId(), Name: ele.Name, IpName: ele.IpName,
	//			CreateName: ele.CreateName, IpID: ele.IpId}
	//		err := DB.Create(tmpSer).Error
	//		if err != nil {
	//			return param.RespUpLoadSeries{}, errors.New("服务正忙...")
	//		}
	//	}
	//}
	//return ret, nil
	return param.RespUpLoadSeries{}, nil
}

func (s *SeriServiceImpl) SearchSeries(req param.ReqSearchSeries) (param.RespSearchSeries, error) {
	DB := s.db.GetDb()
	ser := []db.Series{}
	result := DB.Where("name=? or create_name=? or ip_name=?",
		req.Search, req.Search, req.Search).Find(&ser)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespSearchSeries{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return param.RespSearchSeries{}, errors.New("不存在此系列...")
	} else {
		ret := param.RespSearchSeries{}
		for _, ele := range ser {
			ret.SerInfo = append(ret.SerInfo, param.SerInfo{
				Id:         ele.ID,
				Name:       ele.Name,
				CreateName: ele.CreateName,
				IpId:       *ele.IpID,
				IpName:     ele.IpName,
				CreateTime: ele.CreatedAt,
			})
		}
		return ret, nil
	}
}
func (s *SeriServiceImpl) AddSeries(req param.ReqAddSeries) (uint, error) {
	if *req.IpId == 0 {
		return 0, errors.New("IP不能为空 ...")
	}
	ip := &db.Ip{}
	DB := s.db.GetDb()
	result := DB.Where("id=?", req.IpId).Find(&ip)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return 0, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("该IP不存在...")
	}
	ser := []db.Series{}
	ret := DB.Where("name=?", req.Name).Find(&ser)
	if ret.Error != nil && ret.Error != gorm.ErrRecordNotFound {
		return 0, errors.New("服务正忙...")
	}
	if ser != nil { //存在该系列
		for _, ele := range ser {
			if *ele.IpID == *req.IpId && ele.Name == req.Name {
				return 0, errors.New("该系列已经创建...")
			}
		}
		tmpSer := &db.Series{ID: define.GetRandSeriesId(), Name: req.Name, IpName: req.IpName,
			CreateName: req.CreateName, IpID: req.IpId}
		err := DB.Create(tmpSer).Error
		if err != nil {
			return 0, errors.New("服务正忙...")
		}
		return tmpSer.ID, nil
	} else { //不存在该系列
		tmpSer := &db.Series{ID: define.GetRandSeriesId(), Name: req.Name, IpName: req.IpName,
			CreateName: req.CreateName, IpID: req.IpId}
		err := DB.Create(tmpSer).Error
		if err != nil {
			return 0, errors.New("服务正忙...")
		}
		return tmpSer.ID, nil
	}
}
func (s *SeriServiceImpl) DeleteSeries(req param.ReqDeleteSeries) error {
	if *req.SerId == 0 {
		return errors.New("系列id不能为空 ...")
	}
	DB := s.db.GetDb()
	ser := &db.Series{}
	result := DB.Where("id=?", *req.SerId).Find(&ser)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return errors.New("该系列不存在...")
	}
	err := DB.Model(&db.Series{}).Where("id=?", req.SerId).Delete(&db.Series{}).Error
	if err != nil {
		return errors.New("服务正忙...")
	}
	return nil
}

func (s *SeriServiceImpl) QuerySeries(req param.ReqQuerySeries) (param.RespQuerySeries, error) {
	DB := s.db.GetDb()
	total := int64(0)
	err := DB.Model(&db.Series{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespQuerySeries{}, errors.New("服务正忙......")
	}
	resp := param.RespQuerySeries{}
	ser := []*db.Series{}
	if err := DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&ser).Error; err != nil {
		return param.RespQuerySeries{}, errors.New("服务正忙...")
	}
	for _, ele := range ser {
		resp.ServieInfo.Servies = append(resp.ServieInfo.Servies, param.Servies{
			Id:         &ele.ID,
			Name:       ele.Name,
			CreateName: ele.CreateName,
			IpId:       ele.IpID,
			IpName:     ele.IpName,
			CreateTime: ele.CreatedAt,
		})

	}
	resp.ServieInfo.Num = len(resp.ServieInfo.Servies)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}

func (s *SeriServiceImpl) ModifySeries(req param.ReqModifySeries) error {
	if *req.Id == 0 {
		return errors.New("系列id不能为空 ...")
	}
	DB := s.db.GetDb()
	ser := &db.Series{}
	result := DB.Where("id=?", *req.Id).First(&ser)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return errors.New("该系列不存在...")
	}
	m := make(map[string]interface{})
	if req.Name != "" {
		m["name"] = req.Name
	}
	if req.CreateName != "" {
		m["create_name"] = req.CreateName
	}
	if req.IpId != 0 {
		m["ip_id"] = req.IpId
	}
	err := DB.Model(&db.Series{}).Where("id=?", *req.Id).Updates(m).Error
	if err != nil {
		return err
	}
	return nil
}
