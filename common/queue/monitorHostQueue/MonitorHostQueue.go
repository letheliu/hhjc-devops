package monitorHostQueue

import (
	"sync"

	"github.com/letheliu/hhjc-devops/entity/dto/monitor"
	"github.com/letheliu/hhjc-devops/monitor/service"
)

var lock sync.Mutex
var que chan monitor.MonitorHostDto

/*
*
初始化
*/
func initQueue() {
	if que != nil {
		return
	}
	lock.Lock()
	defer func() {
		lock.Unlock()
	}()
	if que != nil {
		return
	}
	que = make(chan monitor.MonitorHostDto, 100)

	go readData(que)

}

func SendData(host monitor.MonitorHostDto) {
	initQueue()
	que <- host
}

func readData(que chan monitor.MonitorHostDto) {
	for {
		select {
		case data := <-que:
			dealData(data)
		}
	}
}

func dealData(host monitor.MonitorHostDto) {
	var (
		monitorEventService service.MonitorEventService
		monitorEventDto     monitor.MonitorEventDto
		cpuThreshold        = host.CpuThreshold
		memThreshold        = host.MemThreshold
		diskThreshold       = host.DiskThreshold

		remark string
	)

	if host.CpuRate > cpuThreshold {
		remark += (" cpu使用率告警：阀值 " + cpuThreshold + ", 当前 " + host.CpuRate + ";")
	}
	if host.MemRate > memThreshold {
		remark += (" 内存使用率告警：阀值 " + memThreshold + ", 当前 " + host.MemRate + ";")
	}
	if host.DiskRate > diskThreshold {
		remark += (" 磁盘使用率告警：阀值 " + diskThreshold + ", 当前 " + host.DiskRate + ";")
	}

	if remark == "" {
		return
	}

	remark = "主机：" + host.Name + remark

	monitorEventDto = monitor.MonitorEventDto{
		EventType:      "1001",
		EventObjId:     host.HostId,
		EventObjName:   host.Name,
		TenantId:       host.TenantId,
		ThresholdValue: cpuThreshold,
		CurValue:       host.CpuRate,
		Remark:         remark,
		NoticeType:     host.NoticeType,
	}
	monitorEventService.SaveMonitorEvents(monitorEventDto)
}
