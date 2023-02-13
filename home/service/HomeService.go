package service

import (
	"fmt"
	"github.com/kataras/iris/v12"
	appDao "github.com/letheliu/hhjc-devops/appService/dao"
	hostDao "github.com/letheliu/hhjc-devops/assets/dao"
	"github.com/letheliu/hhjc-devops/common/constants"
	"github.com/letheliu/hhjc-devops/entity/dto/appService"
	"github.com/letheliu/hhjc-devops/entity/dto/home"
	"github.com/letheliu/hhjc-devops/entity/dto/host"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/entity/dto/user"
	"github.com/shopspring/decimal"
	"strconv"
)

type HomeService struct {
	hostDaoImpl hostDao.HostDao
	appDaoImpl  appDao.AppServiceDao
}

// get platform data
func (homeService *HomeService) PlatformData(ctx iris.Context) interface{} {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	//get
	hostDto := host.HostDto{
		TenantId: user.TenantId,
	}
	//get host count
	hostCount, _ := homeService.hostDaoImpl.GetHostCount(hostDto)

	hostDto = host.HostDto{
		TenantId: user.TenantId,
	}
	//get cpu mem disk
	hostDto, _ = homeService.hostDaoImpl.GetHostCpuMemDistTotal(hostDto)

	//get app
	appServiceDto := appService.AppServiceDto{
		TenantId: user.TenantId,
	}
	appCount, _ := homeService.appDaoImpl.GetAppServiceCount(appServiceDto)

	//get container
	appServiceContainerDto := appService.AppServiceContainerDto{
		TenantId: user.TenantId,
	}
	dockerCount, _ := homeService.appDaoImpl.GetAppServiceContainerCount(appServiceContainerDto)
	mem, _ := strconv.ParseFloat(hostDto.Mem, 10)
	totalMemDec := decimal.NewFromFloat(mem)
	totalMemDec = totalMemDec.Div(decimal.NewFromInt(1024))
	memValue, _ := totalMemDec.Float64()

	disk, _ := strconv.ParseFloat(hostDto.Disk, 10)
	totalDiskDec := decimal.NewFromFloat(disk)
	totalDiskDec = totalDiskDec.Div(decimal.NewFromInt(1024))
	diskValue, _ := totalDiskDec.Float64()

	platfromDataDto := home.PlatformDataDto{
		HostCount:   hostCount,
		CpuCount:    hostDto.Cpu,
		MemCount:    fmt.Sprintf("%.2f", memValue),
		DiskCount:   fmt.Sprintf("%.2f", diskValue),
		AppCount:    appCount,
		DockerCount: dockerCount,
	}

	return result.SuccessData(platfromDataDto)
}
