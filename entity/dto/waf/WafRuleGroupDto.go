package waf

import (
	"github.com/letheliu/hhjc-devops/entity/dto"
	"time"
)

const (
	Waf_Rule_Group_State_T = "T" //start
	Waf_Rule_Group_State_F = "F" //stop
)

type WafRuleGroupDto struct {
	dto.PageDto
	GroupId    string    `json:"groupId" sql:"-" `
	GroupName  string    `json:"groupName" sql:"-" `
	State      string    `json:"state" `
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
}
