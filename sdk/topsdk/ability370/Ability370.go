package ability370

import (
	"errors"
	"iyyzh-gf/sdk/topsdk"
	request2 "iyyzh-gf/sdk/topsdk/ability370/request"
	response2 "iyyzh-gf/sdk/topsdk/ability370/response"
	"iyyzh-gf/sdk/topsdk/util"
	"log"
)

type Ability370 struct {
	Client *topsdk.TopClient
}

func NewAbility370(client *topsdk.TopClient) *Ability370 {
	return &Ability370{client}
}

/*
   淘宝客-推广者-物料搜索
*/
func (ability *Ability370) TaobaoTbkDgMaterialOptional(req *request2.TaobaoTbkDgMaterialOptionalRequest) (*response2.TaobaoTbkDgMaterialOptionalResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability370 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.material.optional", req.ToMap(), req.ToFileMap())
	var respStruct = response2.TaobaoTbkDgMaterialOptionalResponse{}
	if err != nil {
		log.Println("taobaoTbkDgMaterialOptional error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-店铺搜索
*/
func (ability *Ability370) TaobaoTbkShopGet(req *request2.TaobaoTbkShopGetRequest) (*response2.TaobaoTbkShopGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability370 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.shop.get", req.ToMap(), req.ToFileMap())
	var respStruct = response2.TaobaoTbkShopGetResponse{}
	if err != nil {
		log.Println("taobaoTbkShopGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
