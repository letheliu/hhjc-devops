package businessImages

import "github.com/letheliu/hhjc-devops/entity/dto"

type ImagesPoolsDto struct {
	dto.PageDto
	AppId              string               `json:"appId" sql:"-"`
	Name               string               `json:"name"`
	Version            string               `json:"version"`
	Compose            string               `json:"compose"`
	AppShell           string               `json:"appShell"`
	PublisherId        string               `json:"publisherId"`
	ZihaoAppImagesDtos []ZihaoAppImagesDtos `json:"zihaoAppImagesDtos"`
}
