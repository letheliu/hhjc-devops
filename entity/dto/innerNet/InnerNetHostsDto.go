package innerNet

import (
	"github.com/letheliu/hhjc-devops/entity/dto"
	"time"
)

type InnerNetHostsDto struct {
	dto.PageDto
	InnerNetHostId string    `json:"innerNetHostId" sql:"-" `
	InnerNetId     string    `json:"innerNetId" sql:"-" `
	HostId         string    `json:"hostId" sql:"-"`
	CreateTime     time.Time `json:"createTime" sql:"-"`
	StatusCd       string    `json:"statusCd" sql:"-"`
}
