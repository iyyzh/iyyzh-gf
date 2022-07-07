package top

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "iyyzh/api/v1"
	"iyyzh/internal/consts"
	"topsdk"
	"topsdk/ability370"
	"topsdk/ability370/domain"
	"topsdk/ability370/request"
)

func (s *sTop) Search(ctx context.Context, top v1.Search) (val map[string]interface{}, err error) {
	client := topsdk.NewDefaultTopClient(consts.AppKey, consts.AppSecret, consts.ServerUrlHttp, 20000, 20000)
	ability := ability370.NewAbility370(&client)
	taobaoTbkDgMaterialOptionalUcrowdrankitems := domain.TaobaoTbkDgMaterialOptionalUcrowdrankitems{}
	taobaoTbkDgMaterialOptionalUcrowdrankitems.SetCommirate(1234)
	taobaoTbkDgMaterialOptionalUcrowdrankitems.SetPrice("10.12")
	taobaoTbkDgMaterialOptionalUcrowdrankitems.SetItemId("542808901898")

	req := request.TaobaoTbkDgMaterialOptionalRequest{}
	req.SetQ(top.Q)
	req.SetAdzoneId(consts.AdzoneId)
	req.SetPageSize(top.PageSize)
	req.SetPageNo(top.PageNo)
	req.SetSort(top.Sort)
	req.SetItemloc(top.Itemloc)
	req.SetHasCoupon(top.HasCoupon)

	resp, err := ability.TaobaoTbkDgMaterialOptional(&req)
	if err != nil {
		return nil, err
	} else {
		return gconv.Map(resp.Body), nil
	}
}
