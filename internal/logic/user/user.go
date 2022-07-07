package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"iyyzh/internal/dao"
	"iyyzh/internal/model"
	"iyyzh/internal/model/entity"
	"iyyzh/internal/service"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Creat 创建记录
func (s *sUser) Creat(ctx context.Context, req map[string]interface{}) (err error) {
	_, err1 := dao.User.Ctx(ctx).Data(req).InsertIgnore()
	if err1 != nil {
		return err1
	}
	return
}

// Update 更新记录
func (s *sUser) Update(ctx context.Context, req *entity.User) (err error) {
	res1, err1 := dao.User.Ctx(ctx).Data(req).Where(dao.User.Columns().Id, req.Id).Update()
	if err1 != nil {
		return err1
	}
	res2, err2 := res1.RowsAffected()
	if err2 != nil {
		return err2
	}
	if res2 == 0 {
		err = gerror.New("影响行数为0")
		return err
	}
	return
}

// React 根据id获取记录
func (s *sUser) React(ctx context.Context, id int) (res *entity.User, err error) {
	var One *entity.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&One)
	if err != nil {
		return nil, err
	}
	res = One
	return
}

// Delete 根据id删除单条
func (s *sUser) Delete(ctx context.Context, id int) (err error) {
	res1, err1 := dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, id).Delete()
	if err1 != nil {
		return err1
	}
	res2, err2 := res1.RowsAffected()
	if err2 != nil {
		return err2
	}
	if res2 == 0 {
		err = gerror.New("影响行数为0")
		return err
	}
	return
}

// DeleteByIds 根据id数组批量删除数据
func (s *sUser) DeleteByIds(ctx context.Context, ids []int) (err error) {
	res1, err1 := dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, ids).Delete()
	if err1 != nil {
		return err1
	}
	res2, err2 := res1.RowsAffected()
	if err2 != nil {
		return err2
	}
	if res2 == 0 {
		err = gerror.New("影响行数为0")
		return err
	}
	return
}

// GetList 分页获取记录
func (s *sUser) GetList(ctx context.Context, page *model.Page) (res []*entity.User, total int, err error) {
	var many []*entity.User
	start := (page.From - 1) * 10
	limit := page.Limit
	err = dao.User.Ctx(ctx).Where(page.Title).Limit(start, limit).OrderDesc(dao.User.Columns().Id).Scan(&many)
	if err != nil {
		return nil, 0, err
	}
	total, err = dao.User.Ctx(ctx).CountColumn(dao.User.Columns().Id)
	if err != nil {
		return nil, 0, err
	}
	res = many
	return
}
