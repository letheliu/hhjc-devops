package waf

import (
	"github.com/letheliu/hhjc-devops/entity/dto"
	"time"
)

const Scheme_http = "http"
const Scheme_https = "https"

type WafRouteDto struct {
	dto.PageDto
	WafHostnameCertDto
	RouteId    string    `json:"routeId" sql:"-" `
	WafId      string    `json:"wafId" sql:"-" `
	Hostname   string    `json:"hostname" `
	Ip         string    `json:"ip" `
	Port       string    `json:"port" `
	Scheme     string    `json:"scheme"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
