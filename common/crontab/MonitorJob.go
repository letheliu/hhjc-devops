package crontab

import (
	"github.com/robfig/cron/v3"
	"reflect"
	"sync"

	task2 "github.com/letheliu/hhjc-devops/common/task"
	"github.com/letheliu/hhjc-devops/entity/dto/monitor"
	"github.com/letheliu/hhjc-devops/monitor/dao"
)

var lock sync.Mutex

var c *cron.Cron

type MonitorJob struct {
	monitorHostGroupDao dao.MonitorHostGroupDao
	monitorTaskDao      dao.MonitorTaskDao
}

func (task MonitorJob) init() {

	lock.Lock()
	defer lock.Unlock()
	if c != nil {
		return
	}

	c = cron.New(cron.WithSeconds())

}

// 启动多个任务
func (task MonitorJob) Start() error {
	var (
		hostGroups []*monitor.MonitorHostGroupDto
		taskDtos   []*monitor.MonitorTaskDto
	)

	task.init()

	//停止 所有定时器
	c.Stop()

	//查询host_group
	var monitorHostGroup = monitor.MonitorHostGroupDto{
		State: "3301",
	}
	hostGroups, _ = task.monitorHostGroupDao.GetMonitorHostGroups(monitorHostGroup)

	for _, item := range hostGroups {

		//AddJob方法
		c.AddJob(item.MonCron, task2.HostGroupTask{
			MonitorHostGroupDto: item,
		})
	}
	var taskDto = monitor.MonitorTaskDto{
		State: "002",
	}
	taskDtos, _ = task.monitorTaskDao.GetMonitorTasks(taskDto)

	for _, item := range taskDtos {

		//AddJob方法
		c.AddJob(item.TaskCron, task2.MonitorCommonTask{
			MonitorTaskDto: item,
		})
	}
	//启动计划任务
	c.Start()

	return nil
}
func (job MonitorJob) hasInMonitorTask(dto monitor.MonitorTaskDto) (bool, cron.EntryID) {
	entryies := c.Entries()

	for i := 0; i < len(entryies); i++ {

		if reflect.TypeOf(entryies[i].Job).Name() != "MonitorCommonTask" {
			continue
		}
		id := entryies[i].Job.(task2.MonitorCommonTask).MonitorTaskDto.TaskId
		if id == dto.TaskId {
			return true, entryies[i].ID
		}
	}

	return false, -1
}

func (job MonitorJob) hasInHostGroup(dto monitor.MonitorHostGroupDto) (bool, cron.EntryID) {
	entryies := c.Entries()

	for i := 0; i < len(entryies); i++ {
		if reflect.TypeOf(entryies[i].Job).Name() != "HostGroupTask" {
			continue
		}
		id := entryies[i].Job.(task2.HostGroupTask).MonitorHostGroupDto.MhgId
		if id == dto.MhgId {
			return true, entryies[i].ID
		}
	}

	return false, -1
}

// 启动多个任务
func (task MonitorJob) Restart() {
	//停止 所有定时器

	//启动还没有停止的任务
	task.Start()

}
