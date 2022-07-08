package defaultability

import (
	"errors"
	"iyyzh/sdk/topsdk"
	request2 "iyyzh/sdk/topsdk/defaultability/request"
	response2 "iyyzh/sdk/topsdk/defaultability/response"
	"iyyzh/sdk/topsdk/util"
	"log"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
   关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request2.TaobaoKfcKeywordSearchRequest, session string) (*response2.TaobaoKfcKeywordSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response2.TaobaoKfcKeywordSearchResponse{}
	if err != nil {
		log.Println("taobaoKfcKeywordSearch error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-淘口令预警及拦截查询
*/
func (ability *Defaultability) TaobaoTbkDgTpwdRiskReport(req *request2.TaobaoTbkDgTpwdRiskReportRequest) (*response2.TaobaoTbkDgTpwdRiskReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.tpwd.risk.report", req.ToMap(), req.ToFileMap())
	var respStruct = response2.TaobaoTbkDgTpwdRiskReportResponse{}
	if err != nil {
		log.Println("taobaoTbkDgTpwdRiskReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
