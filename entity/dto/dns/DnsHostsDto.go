package dns

import (
	"github.com/letheliu/hhjc-devops/entity/dto"
	"time"
)

type DnsHostsDto struct {
	dto.PageDto
	DnsHostId  string    `json:"dnsHostId" sql:"-" `
	DnsId      string    `json:"dnsId" sql:"-" `
	HostId     string    `json:"hostId" sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
