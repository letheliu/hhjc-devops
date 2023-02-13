package businessDockerfile

import "github.com/letheliu/hhjc-devops/entity/dto"

const (
	ActionBuild      = "build"
	ActionBuildStart = "buildAndStart"
)

type BusinessDockerfileDto struct {
	dto.PageDto
	ImagesVersion
	Id           string `json:"id" sql:"-"`
	Name         string `json:"name" sql:"-"`
	Version      string `json:"version" sql:"-"`
	Dockerfile   string `json:"dockerfile" sql:"-"`
	CreateUserId string `json:"createUserId" sql:"-"`
	CreateTime   string `json:"createTime" sql:"-"`
	StatusCd     string `json:"statusCd" sql:"-"`
	TenantId     string `json:"tenantId" sql:"-"`
	Username     string `json:"username" sql:"-"`
	LogPath      string `json:"logPath"`
	Action       string `json:"action"`
	AvgIds       string `json:"avgIds"`
}

type BusinessDockerfileCommonDto struct {
	Name         string `json:"name" sql:"-"`
	ShellContext string `json:"shellContext" sql:"-"`
	DeployType   string `json:"deployType" sql:"-"`
	Path         string `json:"path" sql:"-"`
}

type ImagesVersion struct {
	ImagesId string `json:"imagesId"`
	VerId    string `json:"verId"`
}
