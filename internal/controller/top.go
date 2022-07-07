package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"iyyzh/internal/service"

	"iyyzh/api/v1"
)

var Top = cTop{}

type cTop struct{}

func (c *cTop) Search(ctx context.Context, req *v1.TopSearchReq) (res *v1.TopSearchRes, err error) {
	val, err := service.Top().Search(ctx, req.Search)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res = &v1.TopSearchRes{
		Form: val,
		Flag: true,
	}
	return
}

func (c *cTop) Tpwd(ctx context.Context, req *v1.TopTpwdReq) (res *v1.TopTpwdRes, err error) {
	val, err := service.Top().Tpwd(ctx, req.Url)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res = &v1.TopTpwdRes{
		Form: val,
		Flag: true,
	}
	return
}
