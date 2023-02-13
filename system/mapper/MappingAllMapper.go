package mapper

import (
	"github.com/letheliu/hhjc-devops/common/db/mysql"
	"github.com/letheliu/hhjc-devops/common/db/sqlite"
	"github.com/letheliu/hhjc-devops/config"
	"github.com/letheliu/hhjc-devops/entity/dto/mapping"
)

type MappingAllMapper struct {
}

/*
*
查询用户
*/
func (*MappingAllMapper) GetMappings(mappingDto mapping.MappingDto) ([]*mapping.MappingDto, error) {
	var mappingDtos []*mapping.MappingDto
	dbSwatch := config.G_AppConfig.Db

	if Cache_mysql == dbSwatch {
		db := mysql.G_DB.Raw("select * from mapping t where t.status_cd = '0'")
		if err := db.Scan(&mappingDtos).Error; err != nil {
			return nil, err
		}
	}

	if Cache_sqlite == dbSwatch {
		db := sqlite.S_DB.Raw("select * from mapping t where t.status_cd = '0'")
		if err := db.Scan(&mappingDtos).Error; err != nil {
			return nil, err
		}
	}
	return mappingDtos, nil
}
