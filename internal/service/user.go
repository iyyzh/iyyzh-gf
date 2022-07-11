package service

import (
	"context"
	"iyyzh-gf/internal/model"
	"iyyzh-gf/internal/model/entity"
)

var localUser iUser

func RegisterUser(i iUser) {
	localUser = i
}

func User() iUser {
	if localUser == nil {
		panic("implement not found for interface iUser, forgot register?")
	}
	return localUser
}

type iUser interface {
	Creat(ctx context.Context, req map[string]interface{}) (err error)
	Update(ctx context.Context, req *entity.User) (err error)
	React(ctx context.Context, id int) (res *entity.User, err error)
	Delete(ctx context.Context, id int) (err error)
	DeleteByIds(ctx context.Context, ids []int) (err error)
	GetList(ctx context.Context, page *model.Page) (res []*entity.User, total int, err error)
}
