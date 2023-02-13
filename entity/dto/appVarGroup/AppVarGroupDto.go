package appVarGroup

import "github.com/letheliu/hhjc-devops/entity/dto"

const (
	Avg_Type_public string = "001"
)

type AppVarGroupDto struct {
	dto.PageDto
	AvgId      string `json:"avgId" sql:"-"`
	AvgName    string `json:"avgName" sql:"-"`
	AvgType    string `json:"avgType" sql:"-"`
	TenantId   string `json:"tenantId" sql:"-"`
	AvgDesc    string `json:"avgDesc" sql:"-"`
	CreateTime string `json:"createTime" sql:"-"`
	StatusCd   string `json:"statusCd" sql:"-"`
}
