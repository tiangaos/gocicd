package server

import (
	"github.com/kataras/iris/v12"
	"github.com/tiangaos/gops/config"
)

func Index(ctx iris.Context) {
	ctx.ViewData("ContextPath", config.ContextPath)
	ctx.View("server.html")
}
