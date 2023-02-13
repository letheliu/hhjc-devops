package main

import (
	"github.com/letheliu/hhjc-devops/app/router"
	"github.com/letheliu/hhjc-devops/common/task"
	"github.com/letheliu/hhjc-devops/config"
)

func main() {
	//加载配置文件

	config.InitProp("conf/zihao.properties")
	go task.SlaveHealth()

	//这里暂不开启 以免影响 操作比较复杂
	//go task.SlaveFireWall()
	app := iris.New()
	router.HubSlave(app)
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>欢迎访问华恒DevOps平台slave</h1>")
		app.Logger().Info("欢迎访问华恒DevOps平台slave")
	})
	app.Run(iris.Addr(":7001"))

}
