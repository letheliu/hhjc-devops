package firewallRuleDao

import (
	"github.com/letheliu/hhjc-devops/common/db/sqlTemplate"
	"github.com/letheliu/hhjc-devops/common/objectConvert"
	"github.com/letheliu/hhjc-devops/entity/dto"
	"github.com/letheliu/hhjc-devops/entity/dto/firewall"
	"gorm.io/gorm"
)

const (
	query_firewallRule_count string = `
	select count(1) total
from firewall_rule t
					where t.status_cd = '0'
					$if RuleId != '' then
					and t.rule_id = #RuleId#
					$endif
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if Inout != '' then
					and t.in_out = #Inout#
					$endif
					$if AllowLimit != '' then
					and t.allow_limit = #AllowLimit#
					$endif

	`
	query_firewallRule string = `
select t.*,t.in_out inout
from firewall_rule t
					where t.status_cd = '0'
					$if RuleId != '' then
					and t.rule_id = #RuleId#
					$endif
					$if GroupId != '' then
					and t.group_id = #GroupId#
					$endif
					$if Inout != '' then
					and t.in_out = #Inout#
					$endif
					$if AllowLimit != '' then
					and t.allow_limit = #AllowLimit#
					$endif
					order by t.seq
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	query_FirewallRulesByHost string = `
select t.*,t.in_out inout
from firewall_rule t
left join firewall_rule_group frg on t.group_id = frg.group_id and frg.status_cd = '0'
left join host_firewall_group hfg on frg.group_id = hfg.group_id and hfg.status_cd = '0'
where t.status_cd = '0'
and hfg.host_id = #HostId#
order by t.seq
`

	insert_firewallRule string = `
	insert into firewall_rule(rule_id, group_id, in_out,allow_limit,seq,protocol,src_obj,dst_obj,remark)
VALUES(#RuleId#,#GroupId#,#Inout#,#AllowLimit#,#Seq#,#Protocol#,#SrcObj#,#DstObj#,#Remark#)
`

	update_firewallRule string = `
	update firewall_rule set
					$if Seq != '' then
					 seq = #Seq#,
					$endif
					$if Protocol != '' then
					 protocol = #Protocol#,
					$endif
					$if AllowLimit != '' then
					 allow_limit = #AllowLimit#,
					$endif
					$if SrcObj != '' then
					 src_obj = #SrcObj#,
					$endif
					$if DstObj != '' then
					 dst_obj = #DstObj#,
					$endif
					$if Remark != '' then
					 remark = #Remark#,
					$endif
		status_cd = '0'
		where status_cd = '0'
		$if RuleId != '' then
		and rule_id = #RuleId#
		$endif
	`
	delete_firewallRule string = `
	update firewall_rule  set
                          status_cd = '1'
                          where status_cd = '0'
		$if RuleId != '' then
		and rule_id = #RuleId#
		$endif
	`
)

type FirewallRuleDao struct {
}

/*
*
查询用户
*/
func (*FirewallRuleDao) GetFirewallRuleCount(firewallRuleDto firewall.FirewallRuleDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_firewallRule_count, objectConvert.Struct2Map(firewallRuleDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/*
*
查询用户
*/
func (*FirewallRuleDao) GetFirewallRules(firewallRuleDto firewall.FirewallRuleDto) ([]*firewall.FirewallRuleDto, error) {
	var firewallRuleDtos []*firewall.FirewallRuleDto
	sqlTemplate.SelectList(query_firewallRule, objectConvert.Struct2Map(firewallRuleDto), func(db *gorm.DB) {
		db.Scan(&firewallRuleDtos)
	}, false)

	return firewallRuleDtos, nil
}

/*
*
保存服务sql
*/
func (*FirewallRuleDao) SaveFirewallRule(firewallRuleDto firewall.FirewallRuleDto) error {
	return sqlTemplate.Insert(insert_firewallRule, objectConvert.Struct2Map(firewallRuleDto), false)
}

/*
*
修改服务sql
*/
func (*FirewallRuleDao) UpdateFirewallRule(firewallRuleDto firewall.FirewallRuleDto) error {
	return sqlTemplate.Update(update_firewallRule, objectConvert.Struct2Map(firewallRuleDto), false)
}

/*
*
删除服务sql
*/
func (*FirewallRuleDao) DeleteFirewallRule(firewallRuleDto firewall.FirewallRuleDto) error {
	return sqlTemplate.Delete(delete_firewallRule, objectConvert.Struct2Map(firewallRuleDto), false)
}

func (d *FirewallRuleDao) GetFirewallRulesByHost(groupDto firewall.HostFirewallGroupDto) ([]*firewall.FirewallRuleDto, error) {
	var firewallRuleDtos []*firewall.FirewallRuleDto
	sqlTemplate.SelectList(query_FirewallRulesByHost, objectConvert.Struct2Map(groupDto), func(db *gorm.DB) {
		db.Scan(&firewallRuleDtos)
	}, false)

	return firewallRuleDtos, nil
}
