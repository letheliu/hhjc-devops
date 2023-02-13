package service

import (
	"github.com/kataras/iris/v12"
	"strconv"

	"github.com/letheliu/hhjc-devops/appService/dao"
	"github.com/letheliu/hhjc-devops/common/constants"
	"github.com/letheliu/hhjc-devops/common/seq"
	"github.com/letheliu/hhjc-devops/entity/dto/appVersion"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/entity/dto/user"
)

type AppVersionAttrService struct {
	appVersionAttrDao dao.AppVersionAttrDao
}

/*
*
查询 系统信息
*/
func (appVersionAttrService *AppVersionAttrService) GetAppVersionAttrAll(appVersionAttrDto appVersion.AppVersionAttrDto) ([]*appVersion.AppVersionAttrDto, error) {
	var (
		err                error
		appVersionAttrDtos []*appVersion.AppVersionAttrDto
	)

	appVersionAttrDtos, err = appVersionAttrService.appVersionAttrDao.GetAppVersionAttrs(appVersionAttrDto)
	if err != nil {
		return nil, err
	}

	return appVersionAttrDtos, nil

}

/*
*
查询 系统信息
*/
func (appVersionAttrService *AppVersionAttrService) GetAppVersionAttrs(ctx iris.Context) result.ResultDto {
	var (
		err                error
		page               int64
		row                int64
		total              int64
		appVersionAttrDto  = appVersion.AppVersionAttrDto{}
		appVersionAttrDtos []*appVersion.AppVersionAttrDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVersionAttrDto.Row = row * page

	appVersionAttrDto.Page = (page - 1) * row

	appVersionAttrDto.AvId = ctx.URLParam("avId")
	appVersionAttrDto.Version = ctx.URLParam("version")
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionAttrDto.TenantId = user.TenantId

	total, err = appVersionAttrService.appVersionAttrDao.GetAppVersionAttrCount(appVersionAttrDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVersionAttrDtos, err = appVersionAttrService.appVersionAttrDao.GetAppVersionAttrs(appVersionAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionAttrDtos, total, row)

}

/*
*
保存 系统信息
*/
func (appVersionAttrService *AppVersionAttrService) SaveAppVersionAttrs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		appVersionAttrDto appVersion.AppVersionAttrDto
	)

	if err = ctx.ReadJSON(&appVersionAttrDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionAttrDto.TenantId = user.TenantId
	appVersionAttrDto.AvId = seq.Generator()

	err = appVersionAttrService.appVersionAttrDao.SaveAppVersionAttr(appVersionAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionAttrDto)

}

/*
*
修改 系统信息
*/
func (appVersionAttrService *AppVersionAttrService) UpdateAppVersionAttrs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		appVersionAttrDto appVersion.AppVersionAttrDto
	)

	if err = ctx.ReadJSON(&appVersionAttrDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionAttrService.appVersionAttrDao.UpdateAppVersionAttr(appVersionAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionAttrDto)

}

/*
*
删除 系统信息
*/
func (appVersionAttrService *AppVersionAttrService) DeleteAppVersionAttrs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		appVersionAttrDto appVersion.AppVersionAttrDto
	)

	if err = ctx.ReadJSON(&appVersionAttrDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionAttrService.appVersionAttrDao.DeleteAppVersionAttr(appVersionAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionAttrDto)

}
