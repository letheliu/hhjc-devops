package dns

import (
	"github.com/letheliu/hhjc-devops/entity/dto"
	"time"
)

const (
	Type_A = "A"
)

type DnsMapDto struct {
	dto.PageDto
	DnsMapId   string    `json:"dnsMapId"`
	Host       string    `json:"host"`
	Type       string    `json:"type"`
	Value      string    `json:"value"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
