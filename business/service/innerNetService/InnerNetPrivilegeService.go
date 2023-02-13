package innerNetService

import (
	"github.com/kataras/iris/v12"
	"github.com/letheliu/hhjc-devops/business/dao/innerNetDao"
	"github.com/letheliu/hhjc-devops/common/seq"
	"github.com/letheliu/hhjc-devops/entity/dto/innerNet"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"strconv"
)

type InnerNetPrivilegeService struct {
	innerNetDao innerNetDao.InnerNetPrivilegeDao
}

// get db link
// all db by this user
func (innerNetService *InnerNetPrivilegeService) GetInnerNetPrivilegeAll(InnerNetPrivilegeDto innerNet.InnerNetPrivilegeDto) ([]*innerNet.InnerNetPrivilegeDto, error) {
	var (
		err                   error
		InnerNetPrivilegeDtos []*innerNet.InnerNetPrivilegeDto
	)

	InnerNetPrivilegeDtos, err = innerNetService.innerNetDao.GetInnerNetPrivileges(InnerNetPrivilegeDto)
	if err != nil {
		return nil, err
	}

	return InnerNetPrivilegeDtos, nil

}

/*
*
查询 系统信息
*/
func (innerNetService *InnerNetPrivilegeService) GetInnerNetPrivileges(ctx iris.Context) result.ResultDto {
	var (
		err          error
		page         int64
		row          int64
		total        int64
		innerNetDto  = innerNet.InnerNetPrivilegeDto{}
		innerNetDtos []*innerNet.InnerNetPrivilegeDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	innerNetDto.Row = row * page

	innerNetDto.Page = (page - 1) * row

	total, err = innerNetService.innerNetDao.GetInnerNetPrivilegeCount(innerNetDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	innerNetDtos, err = innerNetService.innerNetDao.GetInnerNetPrivileges(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDtos, total, row)

}

/*
*
保存 系统信息
*/
func (innerNetService *InnerNetPrivilegeService) SaveInnerNetPrivileges(ctx iris.Context) result.ResultDto {
	var (
		err         error
		innerNetDto innerNet.InnerNetPrivilegeDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}
	innerNetDto.PId = seq.Generator()
	//InnerNetPrivilegeDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = innerNetService.innerNetDao.SaveInnerNetPrivilege(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

/*
*
修改 系统信息
*/
func (innerNetService *InnerNetPrivilegeService) UpdateInnerNetPrivileges(ctx iris.Context) result.ResultDto {
	var (
		err         error
		innerNetDto innerNet.InnerNetPrivilegeDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	//innerNetDto.Id = ctx.FormValue("id")

	//innerNetDto.Name = ctx.FormValue("name")

	err = innerNetService.innerNetDao.UpdateInnerNetPrivilege(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}

/*
*
删除 系统信息
*/
func (innerNetService *InnerNetPrivilegeService) DeleteInnerNetPrivileges(ctx iris.Context) result.ResultDto {
	var (
		err         error
		innerNetDto innerNet.InnerNetPrivilegeDto
	)
	if err = ctx.ReadJSON(&innerNetDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = innerNetService.innerNetDao.DeleteInnerNetPrivilege(innerNetDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(innerNetDto)

}
