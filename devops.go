package main

import (
	"github.com/letheliu/hhjc-devops/app/router"
	"github.com/letheliu/hhjc-devops/common/cache/factory"
	"github.com/letheliu/hhjc-devops/common/crontab"
	"github.com/letheliu/hhjc-devops/common/db/dbFactory"
	"github.com/letheliu/hhjc-devops/common/ip"
	"github.com/letheliu/hhjc-devops/common/jwt"
	"github.com/letheliu/hhjc-devops/common/kafka"
	"github.com/letheliu/hhjc-devops/config"
	"strconv"
)

/**
 * project address：https://github.com/letheliu/hhjc-devops.git
 * doc address: http://bbs.homecommunity.cn/document.html?docId=102022040475300252
 * author ：wuxw
 * email: 928255095@qq.com
 */
func main() {

	config.InitConfig()
	//support.InitLog()
	//support.InitValidator()
	//mysql.InitGorm()
	dbFactory.Init()
	factory.Init()
	//auth.InitAuth()
	jwt.InitJWT()

	// load qqwry ip data
	ip.InitIP()

	//初始化缓存信息
	factory.InitServiceSql()

	//初始化映射
	factory.InitMapping()

	// init kafka
	kafka.Init()

	//启动定时任务
	crontab.StartCrontab()

	app := iris.New()

	router.Hub(app)
	app.HandleDir("/", "./web")

	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.HTML("<h1>欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！</h1>")
	// 	app.Logger().Info("欢迎访问梓豪平台，这个是后台服务，请直接访问前段服务！")
	// })

	port := config.G_AppConfig.Port

	if port == 0 {
		port = 7000
	}

	app.Run(iris.Addr(":" + strconv.Itoa(port)))

}
