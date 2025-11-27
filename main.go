package main

import (
	_ "se_project/internal/logic"
	_ "se_project/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"se_project/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
