package main

import (
	_ "iyyzh-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	_ "iyyzh-gf/internal/logic"

	"iyyzh-gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
