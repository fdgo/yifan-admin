package service

import (
	"yifan/app/api/param"
	"yifan/app/db"
)

type UserService interface {
	IsNew(req param.ReqIsNew) (bool, error)
	GetOpenId(req param.ReqGetOpenId) (param.RespGetOpenId, error)
	UserList(req param.ReqUserList) (param.RespUserList, error)
	UserListCondition(req param.ReqUserListCondition) (param.RespUserListCondition, error)
	Delever(req param.ReqDelever) (param.RespDelever, error)
	DeleverCondition(req param.ReqDeleverCondition) (param.RespDeleverCondition, error)
	DeleverDetail(req param.ReqDeleverDetail) (param.RespDeleverDetail, error)
	SetDelId(req param.ReqSetDelId) error
}

type UserServiceImpl struct {
	db db.Repo
}

func NewUserService(db db.Repo) UserService {
	return &UserServiceImpl{
		db: db,
	}
}
