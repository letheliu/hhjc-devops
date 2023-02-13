package log

import "github.com/letheliu/hhjc-devops/entity/dto"

// trace dto
type LogTraceParamDto struct {
	dto.PageDto
	Id         string `json:"id"`
	SpanId     string `json:"spanId" sql:"-"`
	ReqHeader  string `json:"reqHeader" sql:"-"`
	ReqParam   string `json:"reqParam" sql:"-"`
	ResHeader  string `json:"resHeader" sql:"-"`
	ResParam   string `json:"resParam" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
