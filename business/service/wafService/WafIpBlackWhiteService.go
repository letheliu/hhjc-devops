package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/letheliu/hhjc-devops/business/dao/wafDao"
	"github.com/letheliu/hhjc-devops/common/seq"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/entity/dto/waf"
	"strconv"
)

type WafIpBlackWhiteService struct {
	wafDao             wafDao.WafIpBlackWhiteDao
	wafRuleDao         wafDao.WafRuleDao
	wafHostnameCertDao wafDao.WafHostnameCertDao
}

// get db link
// all db by this user
func (wafService *WafIpBlackWhiteService) GetWafIpBlackWhiteAll(WafIpBlackWhiteDto waf.WafIpBlackWhiteDto) ([]*waf.WafIpBlackWhiteDto, error) {
	var (
		err                 error
		WafIpBlackWhiteDtos []*waf.WafIpBlackWhiteDto
	)

	WafIpBlackWhiteDtos, err = wafService.wafDao.GetWafIpBlackWhites(WafIpBlackWhiteDto)
	if err != nil {
		return nil, err
	}

	return WafIpBlackWhiteDtos, nil

}

/*
*
查询 系统信息
*/
func (wafService *WafIpBlackWhiteService) GetWafIpBlackWhites(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		wafDto  = waf.WafIpBlackWhiteDto{}
		wafDtos []*waf.WafIpBlackWhiteDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	wafDto.Row = row * page

	wafDto.Page = (page - 1) * row

	total, err = wafService.wafDao.GetWafIpBlackWhiteCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafIpBlackWhites(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDtos, total, row)

}

/*
*
保存 系统信息
*/
func (wafService *WafIpBlackWhiteService) SaveWafIpBlackWhites(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafIpBlackWhiteDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.Id = seq.Generator()
	//WafIpBlackWhiteDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafIpBlackWhite(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	wafRuleDto := waf.WafRuleDto{
		RuleId:   seq.Generator(),
		GroupId:  wafDto.GroupId,
		RuleName: wafDto.Ip,
		Scope:    wafDto.Scope,
		ObjId:    wafDto.Id,
		ObjType:  waf.Waf_obj_type_ip,
		Seq:      wafDto.Seq,
		State:    wafDto.State,
	}
	err = wafService.wafRuleDao.SaveWafRule(wafRuleDto)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessData(wafDto)

}

/*
*
修改 系统信息
*/
func (wafService *WafIpBlackWhiteService) UpdateWafIpBlackWhites(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafIpBlackWhiteDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafIpBlackWhite(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	qWafRuleDto := waf.WafRuleDto{
		ObjId:   wafDto.Id,
		ObjType: waf.Waf_obj_type_ip,
	}
	qWafRuleDtos, _ := wafService.wafRuleDao.GetWafRules(qWafRuleDto)

	if qWafRuleDtos == nil || len(qWafRuleDtos) < 1 {
		return result.Success()
	}

	wafRuleDto := waf.WafRuleDto{
		RuleId:   qWafRuleDtos[0].RuleId,
		GroupId:  wafDto.GroupId,
		RuleName: wafDto.Ip,
		Scope:    wafDto.Scope,
		Seq:      wafDto.Seq,
		State:    wafDto.State,
	}
	err = wafService.wafRuleDao.UpdateWafRule(wafRuleDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/*
*
删除 系统信息
*/
func (wafService *WafIpBlackWhiteService) DeleteWafIpBlackWhites(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafIpBlackWhiteDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafIpBlackWhite(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	qWafRuleDto := waf.WafRuleDto{
		ObjId:   wafDto.Id,
		ObjType: waf.Waf_obj_type_ip,
	}
	qWafRuleDtos, _ := wafService.wafRuleDao.GetWafRules(qWafRuleDto)

	if qWafRuleDtos == nil || len(qWafRuleDtos) < 1 {
		return result.Success()
	}

	wafRuleDto := waf.WafRuleDto{
		RuleId: qWafRuleDtos[0].RuleId,
	}

	err = wafService.wafRuleDao.DeleteWafRule(wafRuleDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
