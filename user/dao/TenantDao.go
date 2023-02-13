package dao

import (
	"github.com/letheliu/hhjc-devops/common/db/sqlTemplate"
	"github.com/letheliu/hhjc-devops/common/objectConvert"
	"github.com/letheliu/hhjc-devops/entity/dto"
	"github.com/letheliu/hhjc-devops/entity/dto/tenant"
	"gorm.io/gorm"
)

const (
	query_tenant_count string = `
		select count(1) total from tenant t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if TenantName != '' then
					and t.tenant_name = #TenantName#
					$endif
					$if TenantType != '' then
					and t.tenant_type = #TenantType#
					$endif
					$if Phone != '' then
					and t.phone = #Phone#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
    	
	`
	query_tenant string = `
		select t.*,uu.username from tenant t
				left join u_user uu on t.tenant_id = uu.tenant_id and uu.user_role = '1001'
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if TenantName != '' then
					and t.tenant_name = #TenantName#
					$endif
					$if TenantType != '' then
					and t.tenant_type = #TenantType#
					$endif
					$if Phone != '' then
					and t.phone = #Phone#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Page != -1 then
						limit #Page#,#Row#
					$endif
	`

	insert_tenant string = `
insert into tenant(tenant_id, tenant_name, address, person_name, phone, remark) 
VALUES(#TenantId#, #TenantName#, #Address#, #PersonName#, #Phone#, #Remark#) 
`

	update_tenant string = `
	update tenant set
			$if TenantName != '' then
			 tenant_name = #TenantName#,
			$endif
			$if TenantType != '' then
			 tenant_type = #TenantType#,
			$endif
			$if Phone != '' then
			 phone = #Phone#,
			$endif
			$if State != '' then
			 state = #State#,
			$endif
			$if Address != '' then
			 address = #Address#,
			$endif
			$if PersonName != '' then
			 person_name = #PersonName#,
			$endif
			$if Remark != '' then
			 remark = #Remark#,
			$endif
			status_cd = '0'
			where status_cd = '0'
			$if TenantId != '' then
			and tenant_id = #TenantId#
			$endif
	`
	delete_tenant string = `
	update tenant set
			status_cd = '1'
			where status_cd = '0'
			and tenant_id = #TenantId#
	`
)

type TenantDao struct {
}

/*
*
查询用户
*/
func (*TenantDao) GetTenantCount(tenantDto tenant.TenantDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_tenant_count, objectConvert.Struct2Map(tenantDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/*
*
查询用户
*/
func (*TenantDao) GetTenants(tenantDto tenant.TenantDto) ([]*tenant.TenantDto, error) {
	var tenantDtos []*tenant.TenantDto
	sqlTemplate.SelectList(query_tenant, objectConvert.Struct2Map(tenantDto), func(db *gorm.DB) {
		db.Scan(&tenantDtos)
	}, false)

	return tenantDtos, nil
}

/*
*
保存服务sql
*/
func (*TenantDao) SaveTenant(tenantDto tenant.TenantDto) error {
	return sqlTemplate.Insert(insert_tenant, objectConvert.Struct2Map(tenantDto), false)
}

/*
*
修改服务sql
*/
func (*TenantDao) UpdateTenant(tenantDto tenant.TenantDto) error {
	return sqlTemplate.Update(update_tenant, objectConvert.Struct2Map(tenantDto), false)
}

/*
*
删除服务sql
*/
func (*TenantDao) DeleteTenant(tenantDto tenant.TenantDto) error {
	return sqlTemplate.Delete(delete_tenant, objectConvert.Struct2Map(tenantDto), false)
}
