package businessDockerfile

import "github.com/zihao-boy/zihao/entity/dto"

const (
	ActionBuild = "build"
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
	Action string 	`json:"action"`
}

type ImagesVersion struct {
	ImagesId string `json:"imagesId"`
	VerId string `json:"verId"`
}
