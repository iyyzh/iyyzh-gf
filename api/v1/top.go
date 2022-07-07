package v1

/**
淘宝客
*/
import (
	"github.com/gogf/gf/v2/frame/g"
)

type Search struct {
	Q         string `json:"q" in:"query"   dc:"搜索内容"`
	PageSize  int32  `json:"pageSize" in:"query"   dc:"分页大小"`
	PageNo    int32  `json:"PageNo" in:"query"   dc:"分页页码"`
	Sort      string `json:"sort" in:"query"   dc:"排序"`
	Itemloc   string `json:"itemloc" in:"query"   dc:"商品所在地区"`
	HasCoupon bool   `json:"hasCoupon" in:"query"   dc:"是否有优惠券"`
}

type TopSearchReq struct {
	g.Meta `path:"/top/search" tags:"物料搜索" method:"get" summary:"淘宝优惠商品信息"`
	Search
}

type TopSearchRes struct {
	Form map[string]interface{} `json:"form"   dc:"商品对象集合"`
	Flag bool                   `json:"flag"   dc:"是否成功"`
}

type TopTpwdReq struct {
	g.Meta `path:"/top/tpwd" tags:"淘口令" method:"post" summary:"优惠链接"`
	Url    string `json:"url" in:"query"   dc:"转换地址"`
}

type TopTpwdRes struct {
	Form map[string]interface{} `json:"form"   dc:"淘口令"`
	Flag bool                   `json:"flag"   dc:"是否成功"`
}
