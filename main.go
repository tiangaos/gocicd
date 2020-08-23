package main

import (
	"github.com/tiangaos/gocicd/model"

	"github.com/kataras/iris/v12"
)

func main() {
	startHttpServer()
}

func startHttpServer() {
	app := iris.New()
	app.Get("/", index)
	app.Get("/projects", GetProjects)
	app.Post("/projects", AddProject)

	app.Run(iris.Addr(":6600"))
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
