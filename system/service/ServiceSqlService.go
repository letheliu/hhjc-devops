package service

import (
	"github.com/kataras/iris/v12"
	"strconv"

	"github.com/letheliu/hhjc-devops/common/encrypt"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/entity/dto/serviceSql"
	"github.com/letheliu/hhjc-devops/system/dao"
)

type ServiceSqlService struct {
	serviceSqlDao dao.ServiceSqlDao
}

/*
*
查询 系统信息
*/
func (serviceSqlService *ServiceSqlService) GetServiceSqls(ctx iris.Context) result.ResultDto {
	var (
		err           error
		page          int64
		row           int64
		total         int64
		serviceSqlDto = serviceSql.ServiceSqlDto{SqlCode: ctx.URLParam("sqlCode"),
			SqlId: ctx.URLParam("sqlId"), Remark: ctx.URLParam("remark")}
		serviceSqlDtos []*serviceSql.ServiceSqlDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	serviceSqlDto.Row = row * page

	serviceSqlDto.Page = (page - 1) * row

	total, err = serviceSqlService.serviceSqlDao.GetServiceSqlCount(serviceSqlDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	serviceSqlDtos, err = serviceSqlService.serviceSqlDao.GetServiceSqls(serviceSqlDto)
	for _, tmpServiceSqlDto := range serviceSqlDtos {
		tmpServiceSqlDto.SqlText = encrypt.Decode(tmpServiceSqlDto.SqlText)
	}

	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDtos, total, row)

}

/*
*
保存 系统信息
*/
func (serviceSqlService *ServiceSqlService) SaveServiceSqls(ctx iris.Context) result.ResultDto {
	var (
		err           error
		serviceSqlDto serviceSql.ServiceSqlDto
	)

	if err = ctx.ReadJSON(&serviceSqlDto); err != nil {
		return result.Error("解析入参失败")
	}
	serviceSqlDto.SqlText = encrypt.Encode(serviceSqlDto.SqlText)

	err = serviceSqlService.serviceSqlDao.SaveServiceSql(serviceSqlDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDto)

}

/*
*
修改 系统信息
*/
func (serviceSqlService *ServiceSqlService) UpdateServiceSqls(ctx iris.Context) result.ResultDto {
	var (
		err           error
		serviceSqlDto serviceSql.ServiceSqlDto
	)

	if err = ctx.ReadJSON(&serviceSqlDto); err != nil {
		return result.Error("解析入参失败")
	}

	serviceSqlDto.SqlText = encrypt.Encode(serviceSqlDto.SqlText)

	err = serviceSqlService.serviceSqlDao.UpdateServiceSql(serviceSqlDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDto)

}

/*
*
删除 系统信息
*/
func (serviceSqlService *ServiceSqlService) DeleteServiceSqls(ctx iris.Context) result.ResultDto {
	var (
		err           error
		serviceSqlDto serviceSql.ServiceSqlDto
	)

	if err = ctx.ReadJSON(&serviceSqlDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = serviceSqlService.serviceSqlDao.DeleteServiceSql(serviceSqlDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(serviceSqlDto)

}
