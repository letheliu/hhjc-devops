package logTraceDao

import (
	"github.com/letheliu/hhjc-devops/common/db/sqlTemplate"
	"github.com/letheliu/hhjc-devops/common/objectConvert"
	"github.com/letheliu/hhjc-devops/entity/dto"
	"github.com/letheliu/hhjc-devops/entity/dto/log"
	"gorm.io/gorm"
)

const (
	query_logTrace_count string = `
	select count(1) total
	from log_trace t
	where t.status_cd = '0'
	$if TraceId != '' then
	and t.trace_id = #TraceId#
	$endif
	$if Name != '' then
	and t.name = #Name#
	$endif
	$if ParentId != '' then
	and t.parent_id = #ParentId#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_logTrace string = `
				select t.*,lt.service_name parent_service_name
				from log_trace t
				left join log_trace lt on t.parent_id = lt.id and lt.status_cd = '0'
				where t.status_cd = '0'
				$if TraceId != '' then
				and t.trace_id = #TraceId#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if ParentId != '' then
				and t.parent_id = #ParentId#
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_logTrace string = `
	insert into log_trace(id, name, parent_id, trace_id,timestamp,duration,service_name,ip,port)
VALUES(#Id#,#Name#,#ParentId#,#TraceId#,#Timestamp#,#Duration#,#ServiceName#,#Ip#,#Port#)
`

	update_logTrace string = `
	update log_trace set
		$if Name != '' then
		name = #Name#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_logTrace string = `
	update log_trace  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type LogTraceDao struct {
}

/*
*
查询用户
*/
func (*LogTraceDao) GetLogTraceCount(logTraceDto log.LogTraceDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_logTrace_count, objectConvert.Struct2Map(logTraceDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/*
*
查询用户
*/
func (*LogTraceDao) GetLogTraces(logTraceDto log.LogTraceDto) ([]*log.LogTraceDto, error) {
	var logTraceDtos []*log.LogTraceDto
	sqlTemplate.SelectList(query_logTrace, objectConvert.Struct2Map(logTraceDto), func(db *gorm.DB) {
		db.Scan(&logTraceDtos)
	}, false)

	return logTraceDtos, nil
}

/*
*
保存服务sql
*/
func (*LogTraceDao) SaveLogTrace(logTraceDto log.LogTraceDto) error {
	return sqlTemplate.Insert(insert_logTrace, objectConvert.Struct2Map(logTraceDto), false)
}

/*
*
修改服务sql
*/
func (*LogTraceDao) UpdateLogTrace(logTraceDto log.LogTraceDto) error {
	return sqlTemplate.Update(update_logTrace, objectConvert.Struct2Map(logTraceDto), false)
}

/*
*
删除服务sql
*/
func (*LogTraceDao) DeleteLogTrace(logTraceDto log.LogTraceDto) error {
	return sqlTemplate.Delete(delete_logTrace, objectConvert.Struct2Map(logTraceDto), false)
}
