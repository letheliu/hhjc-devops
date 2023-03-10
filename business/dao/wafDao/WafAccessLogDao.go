package wafDao

import (
	"github.com/letheliu/hhjc-devops/common/db/sqlTemplate"
	"github.com/letheliu/hhjc-devops/common/objectConvert"
	"github.com/letheliu/hhjc-devops/entity/dto"
	"github.com/letheliu/hhjc-devops/entity/dto/waf"
	"gorm.io/gorm"
)

const (
	query_wafAccessLog_count string = `
	select count(1) total
from waf_access_log t
					where t.status_cd = '0'
					$if RequestId != '' then
					and t.request_id = #RequestId#
					$endif
					$if WafIp != '' then
					and t.waf_ip = #WafIp#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if XRealIp != '' then
					and t.x_real_ip = #XRealIp#
					$endif
					$if Scheme != '' then
					and t.scheme = #Scheme#
					$endif
					$if ResponseCode != '' then
					and t.response_code = #ResponseCode#
					$endif
					$if Method != '' then
					and t.method = #Method#
					$endif
					$if HttpHost != '' then
					and t.http_host = #HttpHost#
					$endif
					$if UpstreamAddr != '' then
					and t.upstream_addr = #UpstreamAddr#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif

	`
	query_wafAccessLog string = `
select t.*
from waf_access_log t
					where t.status_cd = '0'
					$if RequestId != '' then
					and t.request_id = #RequestId#
					$endif
					$if WafIp != '' then
					and t.waf_ip = #WafIp#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if XRealIp != '' then
					and t.x_real_ip = #XRealIp#
					$endif
					$if Scheme != '' then
					and t.scheme = #Scheme#
					$endif
					$if ResponseCode != '' then
					and t.response_code = #ResponseCode#
					$endif
					$if Method != '' then
					and t.method = #Method#
					$endif
					$if HttpHost != '' then
					and t.http_host = #HttpHost#
					$endif
					$if UpstreamAddr != '' then
					and t.upstream_addr = #UpstreamAddr#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	query_wafAccessLogMap string = `
select t.x_real_ip,t.waf_ip,count(1) total from waf_access_log t
where t.create_time > #StartTime#
group by t.x_real_ip,t.waf_ip
	`
	query_wafAccessLogTop5 string = `
select t.x_real_ip,t.waf_ip,count(1) total from waf_access_log t
where t.create_time > #StartTime#
group by t.x_real_ip,t.waf_ip
order by total desc
limit 5
	`

	query_wafAccessLogIntercept string = `
select t.x_real_ip,t.waf_ip,state,t.create_time ,td.name state_name
from waf_access_log t
left join t_dict td on td.table_name = 'waf_access_log'  and td.table_columns = 'state' and td.status_cd = t.state
where  t.state !='default'
order by t.create_time desc
limit 5
	`

	insert_wafAccessLog string = `
	insert into waf_access_log(request_id, waf_ip, host_id,x_real_ip,scheme,response_code,method,http_host,upstream_addr,url,request_length,response_length,state,message)
VALUES(#RequestId#,#WafIp#,#HostId#,#XRealIp#,#Scheme#,#ResponseCode#,#Method#,#HttpHost#,#UpstreamAddr#,#Url#,#RequestLength#,#ResponseLength#,#State#,#Message#)
`

	update_wafAccessLog string = `
	update waf_access_log set
		$if State != '' then
		state = #State#,
		$endif
		$if Message != '' then
		message = #Message#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if RequestId != '' then
		and request_id = #RequestId#
		$endif
	`
	delete_wafAccessLog string = `
	update waf_access_log  set
                          status_cd = '1'
                          where status_cd = '0'
		$if RequestId != '' then
		and request_id = #RequestId#
		$endif
	`
)

type WafAccessLogDao struct {
}

/*
*
????????????
*/
func (*WafAccessLogDao) GetWafAccessLogCount(wafAccessLogDto waf.WafAccessLogDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_wafAccessLog_count, objectConvert.Struct2Map(wafAccessLogDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/*
*
????????????
*/
func (*WafAccessLogDao) GetWafAccessLogs(wafAccessLogDto waf.WafAccessLogDto) ([]*waf.WafAccessLogDto, error) {
	var wafAccessLogDtos []*waf.WafAccessLogDto
	sqlTemplate.SelectList(query_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), func(db *gorm.DB) {
		db.Scan(&wafAccessLogDtos)
	}, false)

	return wafAccessLogDtos, nil
}

/*
*
????????????
*/
func (*WafAccessLogDao) GetWafAccessLogMap(wafAccessLogDto waf.WafAccessLogDto) ([]*waf.WafAccessLogMapDto, error) {
	var wafAccessLogDtos []*waf.WafAccessLogMapDto
	sqlTemplate.SelectList(query_wafAccessLogMap, objectConvert.Struct2Map(wafAccessLogDto), func(db *gorm.DB) {
		db.Scan(&wafAccessLogDtos)
	}, false)

	return wafAccessLogDtos, nil
}

/*
*
????????????
*/
func (*WafAccessLogDao) GetWafAccessLogTop5(wafAccessLogDto waf.WafAccessLogDto) ([]*waf.WafAccessLogMapDto, error) {
	var wafAccessLogDtos []*waf.WafAccessLogMapDto
	sqlTemplate.SelectList(query_wafAccessLogTop5, objectConvert.Struct2Map(wafAccessLogDto), func(db *gorm.DB) {
		db.Scan(&wafAccessLogDtos)
	}, false)

	return wafAccessLogDtos, nil
}

/*
*
????????????
*/
func (*WafAccessLogDao) GetWafAccessLogIntercept(wafAccessLogDto waf.WafAccessLogDto) ([]*waf.WafAccessLogDto, error) {
	var wafAccessLogDtos []*waf.WafAccessLogDto
	sqlTemplate.SelectList(query_wafAccessLogIntercept, objectConvert.Struct2Map(wafAccessLogDto), func(db *gorm.DB) {
		db.Scan(&wafAccessLogDtos)
	}, false)

	return wafAccessLogDtos, nil
}

/*
*
????????????sql
*/
func (*WafAccessLogDao) SaveWafAccessLog(wafAccessLogDto waf.WafAccessLogDto) error {
	return sqlTemplate.Insert(insert_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), false)
}

/*
*
????????????sql
*/
func (*WafAccessLogDao) UpdateWafAccessLog(wafAccessLogDto waf.WafAccessLogDto) error {
	return sqlTemplate.Update(update_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), false)
}

/*
*
????????????sql
*/
func (*WafAccessLogDao) DeleteWafAccessLog(wafAccessLogDto waf.WafAccessLogDto) error {
	return sqlTemplate.Delete(delete_wafAccessLog, objectConvert.Struct2Map(wafAccessLogDto), false)
}
