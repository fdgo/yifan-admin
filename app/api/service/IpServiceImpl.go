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

func (s *IpServiceImpl) UpLoadIPs(req param.ReqUpLoadIPs) (param.RespUpLoadIPs, error) {
	//获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return param.RespUpLoadIPs{}, err
	}
	xlsxPath := dir + "/import.xlsx"
	//打开文件路径
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return param.RespUpLoadIPs{}, err
	}
	DB := s.db.GetDb()
	//读取每一个sheet
	for _, oneSheet := range xlsxFile.Sheets {
		if oneSheet.Name == "IP" {
			ids := []uint{}
			for index, r := range oneSheet.Rows {
				if index != 0 {
					for _, cell := range r.Cells {
						id := define.DealWithOneIp(DB, cell.Value)
						ids = append(ids, id)
					}
				}
			}
			for j, row := range oneSheet.Rows {
				if j == 0 {
					for k, _ := range row.Cells {
						if k+1 == len(row.Cells) {
							x := row.AddCell()
							x.Value = "IP对应ID"
						}
					}
				}
			}
			for j, row := range oneSheet.Rows {
				if j != 0 {
					for k, _ := range row.Cells {
						if k+1 == len(row.Cells) {
							x := row.AddCell()
							x.Value = strconv.Itoa(int(ids[j-1]))
						}
					}
				}
			}
			xlsxFile.Save(xlsxPath)
		}
	}
	return param.RespUpLoadIPs{}, nil
}

func (s *IpServiceImpl) SearchIP(req param.ReqSearchIP) (param.RespSearchIp, error) {
	DB := s.db.GetDb()
	ip := &db.Ip{}
	result := DB.Where("name=?", req.Search).First(&ip)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespSearchIp{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		ip = &db.Ip{}
		result = DB.Where("create_name=?", req.Search).First(&ip)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return param.RespSearchIp{}, errors.New("服务正忙...")
		}
		if result.RowsAffected == 0 {
			return param.RespSearchIp{}, errors.New("不存在此IP...")
		}
		return param.RespSearchIp{ID: ip.ID, Name: ip.Name,
			CreateName: ip.CreateName, CreateTime: ip.CreatedAt}, nil
	} else {
		return param.RespSearchIp{ID: ip.ID, Name: ip.Name,
			CreateName: ip.CreateName, CreateTime: ip.CreatedAt}, nil
	}
}
func (s *IpServiceImpl) AddIP(req param.ReqAddIP) (uint, error) {
	DB := s.db.GetDb()
	ip := &db.Ip{}
	result := DB.Where("name=?", req.Name).First(&ip)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return 0, errors.New("服务正忙...")
	}
	if result.RowsAffected != 0 {
		return 0, errors.New("该IP已经存在...")
	}
	ip = &db.Ip{ID: define.GetRandIpId(), Name: req.Name, CreateName: req.CreateName}
	err := DB.Create(ip).Error
	if err != nil {
		return 0, errors.New("服务正忙...")
	}
	return ip.ID, nil
}
func (s *IpServiceImpl) DeleteIP(req param.ReqDeleteIP) error {
	DB := s.db.GetDb()
	if *req.Id == 0 {
		return errors.New("IP不能为空 ...")
	}
	ip := &db.Ip{}
	result := DB.Where("id=?", *req.Id).First(&ip)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return errors.New("该IP不存在...")
	}
	err := DB.Model(&db.Ip{}).Where("id=?", req.Id).Delete(&db.Ip{}).Error
	if err != nil {
		return errors.New("服务正忙...")
	}
	return nil
}
func (s *IpServiceImpl) QueryIP(req param.ReqQueryIP) (param.RespQueryIP, error) {
	DB := s.db.GetDb()
	total := int64(0)
	err := DB.Model(&db.Ip{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespQueryIP{}, errors.New("服务正忙......")
	}
	ips := []*db.Ip{}
	if err := DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&ips).Error; err != nil {
		return param.RespQueryIP{}, errors.New("服务正忙...")
	}
	result := param.RespQueryIP{}
	for _, oneIp := range ips {
		var series []db.Series
		if err := DB.Model(oneIp).Association("Series").Find(&series); err != nil {
			return param.RespQueryIP{}, errors.New("服务正忙...")
		}
		var respIp param.RespIp
		respIp.ID = oneIp.ID
		respIp.Name = oneIp.Name
		respIp.CreateName = oneIp.CreateName
		respIp.CreateTime = oneIp.CreatedAt
		result.IpInfo.RespIp = append(result.IpInfo.RespIp, respIp)
	}
	result.IpInfo.Num = len(result.IpInfo.RespIp)
	result.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return result, nil
}
func (s *IpServiceImpl) ModifyIP(req param.ReqModifyIP) error {
	if *req.Id == 0 {
		return errors.New("IP不能为空 ...")
	}
	DB := s.db.GetDb()
	ip := &db.Ip{}
	result := DB.Where("id=?", *req.Id).First(&ip)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 {
		return errors.New("该IP不存在...")
	}
	m := make(map[string]interface{})
	if req.Name != "" {
		m["name"] = req.Name
	}
	if req.CreateName != "" {
		m["create_name"] = req.CreateName
	}
	err := DB.Model(&db.Ip{}).Where("id=?", *req.Id).Updates(m).Error
	if err != nil {
		return err
	}
	return nil
}
