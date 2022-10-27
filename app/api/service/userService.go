package service

import (
	"yifan/app/api/param"
	"yifan/app/db"
)

type UserService interface {
	IsNew(req param.ReqIsNew) (bool, error)
	GetOpenId(req param.ReqGetOpenId) (param.RespGetOpenId, error)
	UserList(req param.ReqUserList) (param.RespUserList, error)
}

type UserServiceImpl struct {
	db db.Repo
}

func NewUserService(db db.Repo) UserService {
	return &UserServiceImpl{
		db: db,
	}
}
