package v1

/**
用户
*/
import (
	"github.com/gogf/gf/v2/frame/g"
	"iyyzh/internal/model"
	"iyyzh/internal/model/entity"
)

// UserCreatReq 新建
type UserCreatReq struct {
	g.Meta `path:"/user/creat" method:"post" summary:"UserCreat" tags:"User"`
	Form   map[string]interface{} `json:"form" in:"query"  dc:"Creat对象"`
}
type UserCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess"   dc:"是否成功"`
}

// UserUpdateReq 修改
type UserUpdateReq struct {
	g.Meta `path:"/user/update" method:"post" summary:"UserUpdate" tags:"User"`
	Form   *entity.User `json:"form" in:"query"  dc:"Update对象"`
}
type UserUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess"   dc:"是否成功"`
}

// UserReadReq 读取
type UserReadReq struct {
	g.Meta `path:"/user/read" method:"get" summary:"UserRead" tags:"User"`
	Id     int `json:"id" in:"query"  dc:"主键ID"`
}
type UserReadRes struct {
	g.Meta `mime:"json" example:"obj"`
	Form   *entity.User `json:"form"   dc:"对象"`
}

// UserDeleteReq 删除
type UserDeleteReq struct {
	g.Meta `path:"/user/delete" method:"post" summary:"UserDelete" tags:"User"`
	Id     int `json:"id" in:"query"  dc:"主键ID"`
}
type UserDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess"   dc:"是否成功"`
}

// UserDeleteByIdsReq 批量删除
type UserDeleteByIdsReq struct {
	g.Meta `path:"/user/deleteByIds" method:"post" summary:"UserDeleteByIdsReq" tags:"User"`
	Ids    []int `json:"ids" in:"query"  dc:"主键ID数组"`
}
type UserDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess"   dc:"是否成功"`
}

// UserListReq 分页查询
type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" summary:"UserList" tags:"User"`
	Page   *model.Page `json:"page" in:"query"  dc:"分页数据"`
}
type UserListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.User `json:"list"   dc:"对象数组"`
	Total  int            `json:"total" dc:"记录总数"`
}
