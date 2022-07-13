package top

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"iyyzh-gf/sdk/topsdk"
	"iyyzh-gf/sdk/topsdk/ability375"
	"iyyzh-gf/sdk/topsdk/ability375/request"
)

func (s *sTop) Tpwd(ctx context.Context, url string) (val map[string]interface{}, err error) {
	client := topsdk.NewDefaultTopClient("33893453", "8430ec3b30768bfe3403f45d4033c66f", "http://gw.api.taobao.com/router/rest", 20000, 20000)
	ability := ability375.NewAbility375(&client)

	req := request.TaobaoTbkTpwdCreateRequest{}
	req.SetUrl("http:" + url)

	resp, err := ability.TaobaoTbkTpwdCreate(&req)
	if err != nil {
		return nil, err
	} else {
		return gconv.Map(resp.Body), nil
	}
}
