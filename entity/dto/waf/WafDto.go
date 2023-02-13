package waf

import (
	"github.com/letheliu/hhjc-devops/entity/dto"
	"time"
)

const Waf_state_stop = "2002"
const Waf_state_start = "1001"

type WafDto struct {
	dto.PageDto
	WafId      string         `json:"wafId" sql:"-" `
	WafName    string         `json:"wafName" sql:"-" `
	Port       string         `json:"port" `
	HttpsPort  string         `json:"httpsPort" `
	HostIds    string         `json:"hostIds"`
	WafHosts   []*WafHostsDto `json:"wafHosts"`
	CreateTime time.Time      `json:"createTime" sql:"-"`
	StatusCd   string         `json:"statusCd" sql:"-"`
	State      string         `json:"state"`
}

type SlaveWafDataDto struct {
	ServerIpUrl string                `json:"serverIpUrl"`
	Waf         WafDto                `json:"waf"`
	Routes      []*WafRouteDto        `json:"routes"`
	Certs       []*WafHostnameCertDto `json:"certs"`
	Rules       []*WafRuleDataDto     `json:"rules"`
}
