package mapper

import (
	"github.com/letheliu/hhjc-devops/common/db/mysql"
	"github.com/letheliu/hhjc-devops/common/db/sqlite"
	"github.com/letheliu/hhjc-devops/config"
	"github.com/letheliu/hhjc-devops/entity/dto/serviceSql"
)

const (
	Cache_sqlite = "sqlite"
	Cache_mysql  = "local"
)

type ServiceSqlAllMapper struct {
}

/*
*
查询用户
*/
func (*ServiceSqlAllMapper) GetServiceSqls(serviceSqlDto serviceSql.ServiceSqlDto) ([]*serviceSql.ServiceSqlDto, error) {
	var serviceSqlDtos []*serviceSql.ServiceSqlDto
	dbSwatch := config.G_AppConfig.Db

	if Cache_mysql == dbSwatch {
		db := mysql.G_DB.Raw("select * from service_sql")
		if err := db.Scan(&serviceSqlDtos).Error; err != nil {
			return nil, err
		}
	}

	if Cache_sqlite == dbSwatch {
		db := sqlite.S_DB.Raw("select * from service_sql")
		if err := db.Scan(&serviceSqlDtos).Error; err != nil {
			return nil, err
		}
	}
	return serviceSqlDtos, nil
}
