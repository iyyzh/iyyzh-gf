package controller

import (
	"context"
	"iyyzh/internal/model/entity"
	"iyyzh/internal/service"

	_ "github.com/gogf/gf/v2/frame/g"

	"iyyzh/api/v1"
)

// User 包装对外暴露对象
var User = cUser{}

type cUser struct{}

// Creat 创建记录
func (c cUser) Creat(ctx context.Context, req *v1.UserCreatReq) (res *v1.UserCreatRes, err error) {
	err = service.User().Creat(ctx, req.Form)
	if err != nil {
		return nil, err
	}
	res = &v1.UserCreatRes{
		IsSuccess: true,
	}
	return
}

// Update 更新记录
func (c cUser) Update(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	err = service.User().Update(ctx, req.Form)
	if err != nil {
		return nil, err
	}
	res = &v1.UserUpdateRes{
		IsSuccess: true,
	}
	return
}

// React 根据id获取记录
func (c cUser) React(ctx context.Context, req *v1.UserReadReq) (res *v1.UserReadRes, err error) {
	var Form *entity.User
	Form, err = service.User().React(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.UserReadRes{
		Form: Form,
	}
	return
}

// Delete 根据id删除单条
func (c cUser) Delete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.UserDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除数据
func (c cUser) DeleteByIds(ctx context.Context, req *v1.UserDeleteByIdsReq) (res *v1.UserDeleteByIdsRes, err error) {
	err = service.User().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &v1.UserDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// GetList 分页获取记录
func (c cUser) GetList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	var list []*entity.User
	var total int
	list, total, err = service.User().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &v1.UserListRes{
		List:  list,
		Total: total,
	}
	return
}
