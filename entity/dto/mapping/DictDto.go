package mapping

import "github.com/letheliu/hhjc-devops/entity/dto"

type DictDto struct {
	dto.PageDto

	Id           string `json:"id"`
	StatusCd     string `json:"statusCd" sql:"-"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreateTime   string `json:"createTime" sql:"-"`
	TableName    string `json:"tableName" sql:"-"`
	TableColumns string `json:"tableColumns" sql:"-"`
}
