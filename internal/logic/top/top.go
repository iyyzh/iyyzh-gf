package top

import "iyyzh-gf/internal/service"

type sTop struct {
}

func init() {
	service.RegisterTop(New())
}

func New() *sTop {
	return &sTop{}
}
