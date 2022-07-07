package service

import (
	"context"
	v1 "iyyzh/api/v1"
)

var localTop iTop

func RegisterTop(i iTop) {
	localTop = i
}

func Top() iTop {
	if localTop == nil {
		panic("implement not found for interface iTop, forgot register?")
	}
	return localTop
}

type iTop interface {
	Search(ctx context.Context, top v1.Search) (val map[string]interface{}, err error)
	Tpwd(ctx context.Context, url string) (val map[string]interface{}, err error)
}
