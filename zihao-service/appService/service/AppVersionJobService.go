package service

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/appService/dao"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/date"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appVersionJob"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"os"
	"os/exec"
	systemUser "os/user"
	"strconv"
)

type AppVersionJobService struct {
	appVersionJobDao dao.AppVersionJobDao
}

/**
查询 系统信息
*/
func (appVersionJobService *AppVersionJobService) GetAppVersionJobAll(appVersionJobDto appVersionJob.AppVersionJobDto) ([]*appVersionJob.AppVersionJobDto, error) {
	var (
		err               error
		appVersionJobDtos []*appVersionJob.AppVersionJobDto
	)

	appVersionJobDtos, err = appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	if err != nil {
		return nil, err
	}

	return appVersionJobDtos, nil

}

/**
查询 系统信息
*/
func (appVersionJobService *AppVersionJobService) GetAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err               error
		page              int64
		row               int64
		total             int64
		appVersionJobDto  = appVersionJob.AppVersionJobDto{}
		appVersionJobDtos []*appVersionJob.AppVersionJobDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	appVersionJobDto.Row = row * page

	appVersionJobDto.Page = (page - 1) * row

	appVersionJobDto.JobId = ctx.URLParam("jobId")
	appVersionJobDto.JobName = ctx.URLParam("jobName")
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	appVersionJobDto.TenantId = user.TenantId

	total, err = appVersionJobService.appVersionJobDao.GetAppVersionJobCount(appVersionJobDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	appVersionJobDtos, err = appVersionJobService.appVersionJobDao.GetAppVersionJobs(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDtos, total, row)

}

/**
保存 系统信息
*/
func (appVersionJobService *AppVersionJobService) SaveAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDto.JobId = seq.Generator()
	appVersionJobDto.State = appVersionJob.STATE_wait
	appVersionJobDto.PreJobTime = date.GetNowTimeString()
	appVersionJobDto.CurJobTime = date.GetNowTimeString()

	err = appVersionJobService.appVersionJobDao.SaveAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	tmpUser, _ := systemUser.Current()
	var path string = tmpUser.HomeDir + "/zihao/" + appVersionJobDto.JobId+"/"
	var fileName string = appVersionJobDto.JobId + ".sh"

	_, err = os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
	}

	//当前用户目录下生成 文件夹
	file, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	defer func() { file.Close() }()
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(path +  fileName)
	}
	_, err = file.WriteString("cd " + path+"\n")
	_, err = file.WriteString(appVersionJobDto.JobShell)

	if err != nil {
		fmt.Print("err=", err.Error())
	}

	return result.SuccessData(appVersionJobDto)

}

/**
修改 系统信息
*/
func (appVersionJobService *AppVersionJobService) UpdateAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	tmpUser, _ := systemUser.Current()
	var path string = tmpUser.HomeDir + "/zihao/" + appVersionJobDto.JobId + "/"
	var fileName string = appVersionJobDto.JobId + ".sh"

	_, err = os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
	}

	//当前用户目录下生成 文件夹
	file, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	defer func() { file.Close() }()
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(path+ fileName)
	}
	_, err = file.WriteString("cd " + path+"\n")
	_, err = file.WriteString(appVersionJobDto.JobShell)

	if err != nil {
		fmt.Print("err=", err.Error())
	}

	return result.SuccessData(appVersionJobDto)

}

/**
删除 系统信息
*/
func (appVersionJobService *AppVersionJobService) DoJob(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	appVersionJobDto.TenantId = user.TenantId
	appVersionJobDto.State = appVersionJob.STATE_doing
	err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	tmpUser, _ := systemUser.Current()
	var path string = tmpUser.HomeDir + "/zihao/" + appVersionJobDto.JobId+"/"
	var fileName string = path + appVersionJobDto.JobId + ".sh"

	jobShell :=  "nohup sh "+ fileName + " >" + path + appVersionJobDto.JobId + ".log &"
	cmd := exec.Command("bash", "-c",jobShell)
	fmt.Println(jobShell)
	//cmd := exec.Command("nohup echo 1")
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		appVersionJobDto.State = appVersionJob.STATE_error
		err = appVersionJobService.appVersionJobDao.UpdateAppVersionJob(appVersionJobDto)
	}
	return result.SuccessData(appVersionJobDto)

}

/**
构建
*/
func (appVersionJobService *AppVersionJobService) DeleteAppVersionJobs(ctx iris.Context) result.ResultDto {
	var (
		err              error
		appVersionJobDto appVersionJob.AppVersionJobDto
	)

	if err = ctx.ReadJSON(&appVersionJobDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = appVersionJobService.appVersionJobDao.DeleteAppVersionJob(appVersionJobDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(appVersionJobDto)

}
