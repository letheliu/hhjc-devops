package appVersion

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/appService/service"
)

type AppVersionController struct{
	appVersionService service.AppVersionService
}


func AppVersionControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/appVersion")
		aus      = AppVersionController{appVersionService: service.AppVersionService{}}
	)
	adinMenu.Get("/getAppVersion", hero.Handler(aus.getAppVersion))

	adinMenu.Post("/saveAppVersion", hero.Handler(aus.saveAppVersion))

	adinMenu.Post("/updateAppVersion", hero.Handler(aus.updateAppVersion))

	adinMenu.Post("/deleteAppVersion", hero.Handler(aus.deleteAppVersion))

}

/**
查询 主机组
 */
func (aus *AppVersionController) getAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.GetAppVersions(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppVersionController) saveAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.SaveAppVersions(ctx)

	ctx.JSON(reslut)
}


/**
修改 主机组
*/
func (aus *AppVersionController) updateAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.UpdateAppVersions(ctx)

	ctx.JSON(reslut)
}


/**
删除 主机组
*/
func (aus *AppVersionController) deleteAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.DeleteAppVersions(ctx)

	ctx.JSON(reslut)
}
