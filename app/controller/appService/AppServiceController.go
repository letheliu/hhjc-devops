package appService

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/appService/service"
)

type AppServiceController struct {
	appServiceService  service.AppServiceService
	appVarGroupService service.AppVarGroupService
	appVarService      service.AppVarService
}

func AppServiceControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/appService")
		aus      = AppServiceController{appServiceService: service.AppServiceService{}}
	)
	adinMenu.Get("/getAppService", hero.Handler(aus.getAppService))

	adinMenu.Post("/saveAppService", hero.Handler(aus.saveAppService))

	adinMenu.Post("/updateAppService", hero.Handler(aus.updateAppService))

	adinMenu.Post("/deleteAppService", hero.Handler(aus.deleteAppService))

	adinMenu.Get("/getAppVarGroup", hero.Handler(aus.getAppVarGroup))

	adinMenu.Post("/saveAppVarGroup", hero.Handler(aus.saveAppVarGroup))

	adinMenu.Post("/updateAppVarGroup", hero.Handler(aus.updateAppVarGroup))

	adinMenu.Post("/deleteAppVarGroup", hero.Handler(aus.deleteAppVarGroup))

	adinMenu.Get("/getAppVar", hero.Handler(aus.getAppVar))

	adinMenu.Post("/saveAppVar", hero.Handler(aus.saveAppVar))

	adinMenu.Post("/updateAppVar", hero.Handler(aus.updateAppVar))

	adinMenu.Post("/deleteAppVar", hero.Handler(aus.deleteAppVar))

	adinMenu.Get("/getAppServiceVar", hero.Handler(aus.getAppServiceVar))

	adinMenu.Post("/saveAppServiceVar", hero.Handler(aus.saveAppServiceVar))

	adinMenu.Post("/updateAppServiceVar", hero.Handler(aus.updateAppServiceVar))

	adinMenu.Post("/deleteAppServiceVar", hero.Handler(aus.deleteAppServiceVar))

	adinMenu.Post("/deleteAppVar", hero.Handler(aus.deleteAppVar))

	adinMenu.Get("/getAppServiceHosts", hero.Handler(aus.getAppServiceHosts))

	adinMenu.Post("/saveAppServiceHosts", hero.Handler(aus.saveAppServiceHosts))

	adinMenu.Post("/updateAppServiceHosts", hero.Handler(aus.updateAppServiceHosts))

	adinMenu.Post("/deleteAppServiceHosts", hero.Handler(aus.deleteAppServiceHosts))

}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppService(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServices(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppService(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServices(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppService(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServices(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppService(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServices(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.GetAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.SaveAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.UpdateAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.DeleteAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppVar(ctx iris.Context) {
	reslut := aus.appVarService.GetAppVars(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppVar(ctx iris.Context) {
	reslut := aus.appVarService.SaveAppVars(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppVar(ctx iris.Context) {
	reslut := aus.appVarService.UpdateAppVars(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppVar(ctx iris.Context) {
	reslut := aus.appVarService.DeleteAppVars(ctx)

	ctx.JSON(reslut)
}



/**
查询 主机组
*/
func (aus *AppServiceController) getAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServiceVars(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServiceVars(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServiceVars(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServiceVars(ctx)

	ctx.JSON(reslut)
}



/**
查询 主机组
*/
func (aus *AppServiceController) getAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServiceHosts(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServiceHosts(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServiceHosts(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServiceHosts(ctx)

	ctx.JSON(reslut)
}
