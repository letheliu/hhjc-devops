package dao

import (
	"github.com/letheliu/hhjc-devops/common/db/sqlTemplate"
	"github.com/letheliu/hhjc-devops/common/objectConvert"
	"github.com/letheliu/hhjc-devops/entity/dto"
	"github.com/letheliu/hhjc-devops/entity/dto/businessDockerfile"
	"gorm.io/gorm"
)

const (
	query_businessDockerfile_count string = `
	select count(1) total
	from business_dockerfile t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name like '%' || #Name# || '%'
	$endif
	$if Version != '' then
	and t.version like '%' || #Version# || '%'
	$endif
	$if CreateUserId != '' then
	and t.create_user_id = #CreateUserId#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_businessDockerfile string = `
				select t.*,uu.username
				from business_dockerfile t
				left join u_user uu on t.create_user_id = uu.user_id and uu.status_cd = '0'
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name like '%' || #Name# || '%'
				$endif
				$if Version != '' then
				and t.version like '%' || #Version# || '%'
				$endif
			
				$if CreateUserId != '' then
				and t.create_user_id = #CreateUserId#
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_businessDockerfile string = `
	insert into business_dockerfile(id, name, version, dockerfile, create_user_id,tenant_id,log_path)
VALUES(#Id#,#Name#,#Version#,#Dockerfile#,#CreateUserId#,#TenantId#,#LogPath#)
`

	update_businessDockerfile string = `
	update business_dockerfile set
		$if Name != '' then
		name = #Name#,
		$endif
		$if Dockerfile != '' then
		dockerfile = #Dockerfile#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_businessDockerfile string = `
	update business_dockerfile  set
                          status_cd = '1'
                          where status_cd = '0'
						  and id = #Id#
	`
)

type BusinessDockerfileDao struct {
}

/*
*
查询用户
*/
func (*BusinessDockerfileDao) GetBusinessDockerfileCount(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_businessDockerfile_count, objectConvert.Struct2Map(businessDockerfileDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/*
*
查询用户
*/
func (*BusinessDockerfileDao) GetBusinessDockerfiles(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) ([]*businessDockerfile.BusinessDockerfileDto, error) {
	var businessDockerfileDtos []*businessDockerfile.BusinessDockerfileDto
	sqlTemplate.SelectList(query_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), func(db *gorm.DB) {
		db.Scan(&businessDockerfileDtos)
	}, false)

	return businessDockerfileDtos, nil
}

/*
*
保存服务sql
*/
func (*BusinessDockerfileDao) SaveBusinessDockerfile(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) error {
	return sqlTemplate.Insert(insert_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), false)
}

/*
*
修改服务sql
*/
func (*BusinessDockerfileDao) UpdateBusinessDockerfile(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) error {
	return sqlTemplate.Update(update_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), false)
}

/*
*
删除服务sql
*/
func (*BusinessDockerfileDao) DeleteBusinessDockerfile(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) error {
	return sqlTemplate.Delete(delete_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), false)
}
