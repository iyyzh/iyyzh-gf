package model

// OptionData 字典配置值
type OptionData struct {
	Label    string        `json:"label"`
	Value    int           `json:"value"`
	Children []*OptionData `json:"children"      ` // 子路由
}

// Page 分页
type Page struct {
	Title interface{} `json:"title" in:"query"   dc:"筛选条件"`
	From  int         `json:"from"  in:"query"  v:"bail|required|integer" dc:"页数"`
	Limit int         `json:"limit" in:"query"  v:"bail|required|integer" dc:"条数"`
}
