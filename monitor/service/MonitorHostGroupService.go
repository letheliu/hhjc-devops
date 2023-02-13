package service

import (
	"github.com/kataras/iris/v12"
	"strconv"

	"github.com/letheliu/hhjc-devops/common/constants"
	"github.com/letheliu/hhjc-devops/common/date"
	"github.com/letheliu/hhjc-devops/common/seq"
	"github.com/letheliu/hhjc-devops/entity/dto/monitor"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/entity/dto/user"
	"github.com/letheliu/hhjc-devops/monitor/dao"
)

type MonitorHostGroupService struct {
	monitorHostGroupDao dao.MonitorHostGroupDao
}

/*
*
查询 系统信息
*/
func (monitorHostGroupService *MonitorHostGroupService) GetMonitorHostGroups(ctx iris.Context) result.ResultDto {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		err                  error
		page                 int64
		row                  int64
		total                int64
		monitorHostGroupDto  = monitor.MonitorHostGroupDto{TenantId: user.TenantId}
		monitorHostGroupDtos []*monitor.MonitorHostGroupDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	monitorHostGroupDto.Row = row * page

	monitorHostGroupDto.Page = (page - 1) * row

	total, err = monitorHostGroupService.monitorHostGroupDao.GetMonitorHostGroupCount(monitorHostGroupDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	monitorHostGroupDtos, err = monitorHostGroupService.monitorHostGroupDao.GetMonitorHostGroups(monitorHostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostGroupDtos, total, row)

}

/*
*
保存 系统信息
*/
func (monitorHostGroupService *MonitorHostGroupService) SaveMonitorHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err                 error
		monitorHostGroupDto monitor.MonitorHostGroupDto
	)

	if err = ctx.ReadJSON(&monitorHostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	monitorHostGroupDto.TenantId = user.TenantId
	monitorHostGroupDto.MhgId = seq.Generator()
	monitorHostGroupDto.State = "3302"
	monitorHostGroupDto.MonDate = date.GetNowTimeString()

	err = monitorHostGroupService.monitorHostGroupDao.SaveMonitorHostGroup(monitorHostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostGroupDto)

}

/*
*
修改 系统信息
*/
func (monitorHostGroupService *MonitorHostGroupService) UpdateMonitorHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err                 error
		monitorHostGroupDto monitor.MonitorHostGroupDto
	)

	if err = ctx.ReadJSON(&monitorHostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorHostGroupService.monitorHostGroupDao.UpdateMonitorHostGroup(monitorHostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostGroupDto)

}

/*
*
删除 系统信息
*/
func (monitorHostGroupService *MonitorHostGroupService) DeleteMonitorHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err                 error
		monitorHostGroupDto monitor.MonitorHostGroupDto
	)

	if err = ctx.ReadJSON(&monitorHostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorHostGroupService.monitorHostGroupDao.DeleteMonitorHostGroup(monitorHostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostGroupDto)

}

/*
*
启动组
*/
func (monitorHostGroupService *MonitorHostGroupService) StartMonitorHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err                 error
		monitorHostGroupDto monitor.MonitorHostGroupDto
	)

	if err = ctx.ReadJSON(&monitorHostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}

	monitorHostGroupDto.State = "3301"
	err = monitorHostGroupService.monitorHostGroupDao.UpdateMonitorHostGroup(monitorHostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostGroupDto)

}

/*
*
停止组
*/
func (monitorHostGroupService *MonitorHostGroupService) StopMonitorHostGroups(ctx iris.Context) result.ResultDto {
	var (
		err                 error
		monitorHostGroupDto monitor.MonitorHostGroupDto
	)

	if err = ctx.ReadJSON(&monitorHostGroupDto); err != nil {
		return result.Error("解析入参失败")
	}
	monitorHostGroupDto.State = "3302"

	err = monitorHostGroupService.monitorHostGroupDao.UpdateMonitorHostGroup(monitorHostGroupDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostGroupDto)

}
