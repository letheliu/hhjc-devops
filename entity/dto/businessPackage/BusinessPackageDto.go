package businessPackage

import "github.com/letheliu/hhjc-devops/entity/dto"

type BusinessPackageDto struct {
	dto.PageDto
	Id           string `json:"id" sql:"-"`
	Name         string `json:"name" sql:"-"`
	Varsion      string `json:"varsion" sql:"-"`
	Path         string `json:"path" sql:"-"`
	BasePath     string `json:"basePath" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	Username     string `json:"username"`
}
