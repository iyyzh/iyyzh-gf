package main

import (
	_ "iyyzh/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	_ "iyyzh/internal/logic"

	"iyyzh/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
