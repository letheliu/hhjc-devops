package k8sContainerScheduling

import (
	"github.com/letheliu/hhjc-devops/entity/dto/appService"
	"github.com/letheliu/hhjc-devops/entity/dto/host"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
)

// k8s 调度器
func Scheduling(hosts []*host.HostDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {
	return result.Success(), nil
}

// default  stop
// base on mem
// add by wuxw 2021-12-07
func StopContainer(containerDto *appService.AppServiceContainerDto, appServiceDto *appService.AppServiceDto) (result.ResultDto, error) {
	return result.Success(), nil
}
