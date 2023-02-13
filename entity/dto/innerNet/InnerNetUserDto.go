package innerNet

import (
	"github.com/letheliu/hhjc-devops/entity/dto"
	"time"
)

type InnerNetUserDto struct {
	dto.PageDto
	UserId     string    `json:"userId" sql:"-" `
	Username   string    `json:"username"  `
	Password   string    `json:"password" `
	Tel        string    `json:"tel"`
	Ip         string    `json:"ip"`
	LoginTime  time.Time `json:"loginTime" sql:"-"`
	CreateTime time.Time `json:"createTime" sql:"-"`
	StatusCd   string    `json:"statusCd" sql:"-"`
	Token      string    `json:"token"`
}
