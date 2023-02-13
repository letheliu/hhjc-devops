package host

import "github.com/letheliu/hhjc-devops/entity/dto"

/*
*
主机组实体
*/
type HostGroupDto struct {
	dto.PageDto
	GroupId     string `json:"groupId" sql:"-"`
	Name        string `json:"name" `
	Description string `json:"description" `
	TenantId    string `json:"tenantId" sql:"-"`
	CreateTime  string `json:"createTime" sql:"-"`
	StatusCd    string `json:"statusCd" sql:"-"`
}
