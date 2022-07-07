package top

import "iyyzh/internal/service"

type sTop struct {
}

func init() {
	service.RegisterTop(New())
}

func New() *sTop {
	return &sTop{}
}
