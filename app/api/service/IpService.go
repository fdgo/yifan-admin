package service

import (
	"yifan/app/api/param"
	"yifan/app/cache"
	"yifan/app/db"
)

type IpService interface {
	UpLoadIPs(req param.ReqUpLoadIPs) (param.RespUpLoadIPs, error)
	SearchIP(req param.ReqSearchIP) (param.RespSearchIp, error)
	AddIP(req param.ReqAddIP) (uint, error)
	DeleteIP(req param.ReqDeleteIP) error
	QueryIP(req param.ReqQueryIP) (param.RespQueryIP, error)
	ModifyIP(req param.ReqModifyIP) error
}

type IpServiceImpl struct {
	db    db.Repo
	cache *cache.CacheRepo
}

func NewIpService(db db.Repo, cache *cache.CacheRepo) IpService {
	return &IpServiceImpl{
		db:    db,
		cache: cache,
	}
}
