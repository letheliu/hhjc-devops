package dbClient

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/letheliu/hhjc-devops/business/service/dbLinkService"
)

type DbClientController struct {
	dbLinkService dbLinkService.DbLinkService
}

func DbClientControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/dbClient")
		aus      = DbClientController{
			dbLinkService: dbLinkService.DbLinkService{},
		}
	)
	adinMenu.Get("/getDbLink", hero.Handler(aus.getDbLink))

	adinMenu.Post("/saveDbLink", hero.Handler(aus.saveDbLink))

	adinMenu.Post("/updateDbLink", hero.Handler(aus.updateDbLink))

	adinMenu.Post("/deleteDbLink", hero.Handler(aus.deleteDbLink))

	//execute sql
	adinMenu.Post("/execSql", hero.Handler(aus.execSql))

	//execute sql file
	adinMenu.Post("/exportSqlFile", hero.Handler(aus.exportSqlFile))
}

/*
*
查询 业务包
*/
func (aus *DbClientController) getDbLink(ctx iris.Context) {
	reslut := aus.dbLinkService.GetDbLinks(ctx)
	ctx.JSON(reslut)
}

/*
*
添加 业务包
*/
func (aus *DbClientController) saveDbLink(ctx iris.Context) {
	reslut := aus.dbLinkService.SaveDbLinks(ctx)

	ctx.JSON(reslut)
}

/*
*
修改 业务包
*/
func (aus *DbClientController) updateDbLink(ctx iris.Context) {
	reslut := aus.dbLinkService.UpdateDbLinks(ctx)

	ctx.JSON(reslut)
}

/*
*
删除 业务包
*/
func (aus *DbClientController) deleteDbLink(ctx iris.Context) {
	reslut := aus.dbLinkService.DeleteDbLinks(ctx)

	ctx.JSON(reslut)
}

/*
*
删除 业务包
*/
func (aus *DbClientController) execSql(ctx iris.Context) {
	reslut := aus.dbLinkService.ExecSql(ctx)

	ctx.JSON(reslut)
}

/*
*
export sql file
*/
func (aus *DbClientController) exportSqlFile(ctx iris.Context) {
	reslut := aus.dbLinkService.ExportSqlFile(ctx)

	ctx.JSON(reslut)
}
