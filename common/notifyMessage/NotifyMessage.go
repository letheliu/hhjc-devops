package notifyMessage

import (
	"github.com/letheliu/hhjc-devops/common/cache/factory"
	"github.com/letheliu/hhjc-devops/common/httpReq"
	"github.com/letheliu/hhjc-devops/entity/dto/tenant"
	"github.com/letheliu/hhjc-devops/user/service"
)

const DINGDING = "DINGDING"
const WECHAT = "WECHAT"

func SendMsg(tenantId, message string) {
	mapping, err := factory.GetMapping("SEND_WAY")

	if err != nil {
		sendToDingDing(tenantId, message)
		return
	}

	if WECHAT == mapping.Value {
		sendToCompanyWechat(tenantId, message)
	} else {
		sendToDingDing(tenantId, message)
	}
}

func sendToCompanyWechat(tenantId, message string) (string, error) {

	//eventDto.TenantId

	var tenantSettingService service.TenantSettingService
	var tenantSettingDto = tenant.TenantSettingDto{
		TenantId: tenantId,
		SpecCd:   "300302",
	}
	tenantSettingDtos, err := tenantSettingService.GetTenantSettingAll(tenantSettingDto)

	if err != nil || len(tenantSettingDtos) < 1 {
		return "", err
	}
	//根据告警类型告警相应平台
	var url string = tenantSettingDtos[0].Value
	// 1、构建需要的参数
	context := map[string]string{
		"content": "[华恒DevOps平台]" + message,
	}
	data := map[string]interface{}{
		"msgtype": "text",
		"text":    context,
	}
	resp, err := httpReq.SendRequest(url, data, nil, "POST")
	return string(resp), err
}

func sendToDingDing(tenantId, message string) (string, error) {

	//eventDto.TenantId

	var tenantSettingService service.TenantSettingService
	var tenantSettingDto = tenant.TenantSettingDto{
		TenantId: tenantId,
		SpecCd:   "300301",
	}
	tenantSettingDtos, err := tenantSettingService.GetTenantSettingAll(tenantSettingDto)

	if err != nil || len(tenantSettingDtos) < 1 {
		return "", err
	}
	//根据告警类型告警相应平台
	var url string = tenantSettingDtos[0].Value
	// 1、构建需要的参数
	context := map[string]string{
		"content": "[华恒DevOps平台]" + message,
	}
	data := map[string]interface{}{
		"msgtype": "text",
		"text":    context,
	}
	resp, err := httpReq.SendRequest(url, data, nil, "POST")
	return string(resp), err
}
