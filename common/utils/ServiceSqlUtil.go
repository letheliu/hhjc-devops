package utils

import (
	"github.com/letheliu/hhjc-devops/common/cache/local"
	"github.com/letheliu/hhjc-devops/common/cache/redis"
	"github.com/letheliu/hhjc-devops/common/encrypt"
	"github.com/letheliu/hhjc-devops/config"
	"github.com/letheliu/hhjc-devops/entity/dto/serviceSql"
)

const (
	Cache_redis = "redis"
	Cache_local = "local"
)

func GetServiceSql(sqlCode string) serviceSql.ServiceSqlDto {
	cacheSwatch := config.G_AppConfig.Cache
	var (
		serviceSqlDto serviceSql.ServiceSqlDto
	)
	if Cache_redis == cacheSwatch {
		serviceSqlDto, _ = redis.G_Redis.GetServiceSql(sqlCode)
	}

	if Cache_local == cacheSwatch {
		serviceSqlDto, _ = local.G_Local.GetServiceSql(sqlCode)
	}
	serviceSqlDto.SqlText = encrypt.Decode(serviceSqlDto.SqlText)
	return serviceSqlDto
}
