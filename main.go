package main

import (
	"github.com/tiangaos/gops/config"
	"github.com/tiangaos/gops/model"
	"github.com/tiangaos/gops/server"

	"github.com/kataras/iris/v12"
)

func main() {
	startHttpServer()
}

func startHttpServer() {
	app := iris.New()
	app.HandleDir(config.ContextPath+"/", "./static/")
	tmpl := iris.HTML("./static", ".html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	app.Get(config.ContextPath+"/server", server.Index)

	app.Get("/projects", GetProjects)
	app.Post("/projects", AddProject)

	app.Run(iris.Addr(":9527"))
}

func AddProject(ctx iris.Context) {
	var p model.Project
	ctx.ReadJSON(&p)
	p.Save()
	ctx.JSON(p)
}

func GetProjects(ctx iris.Context) {
	ctx.JSON(model.GetAllProjects())
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>GO cicd</h1>")
}
