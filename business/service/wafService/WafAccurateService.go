package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/letheliu/hhjc-devops/business/dao/wafDao"
	"github.com/letheliu/hhjc-devops/common/seq"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/entity/dto/waf"
	"strconv"
)

type WafAccurateService struct {
	wafDao             wafDao.WafAccurateDao
	wafRuleDao         wafDao.WafRuleDao
	wafHostnameCertDao wafDao.WafHostnameCertDao
}

// get db link
// all db by this user
func (wafService *WafAccurateService) GetWafAccurateAll(WafAccurateDto waf.WafAccurateDto) ([]*waf.WafAccurateDto, error) {
	var (
		err             error
		WafAccurateDtos []*waf.WafAccurateDto
	)

	WafAccurateDtos, err = wafService.wafDao.GetWafAccurates(WafAccurateDto)
	if err != nil {
		return nil, err
	}

	return WafAccurateDtos, nil

}

/*
*
查询 系统信息
*/
func (wafService *WafAccurateService) GetWafAccurates(ctx iris.Context) result.ResultDto {
	var (
		err     error
		page    int64
		row     int64
		total   int64
		wafDto  = waf.WafAccurateDto{}
		wafDtos []*waf.WafAccurateDto
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

	total, err = wafService.wafDao.GetWafAccurateCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafAccurates(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDtos, total, row)

}

/*
*
保存 系统信息
*/
func (wafService *WafAccurateService) SaveWafAccurates(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafAccurateDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.Id = seq.Generator()
	//WafAccurateDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWafAccurate(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	wafRuleDto := waf.WafRuleDto{
		RuleId:   seq.Generator(),
		GroupId:  wafDto.GroupId,
		RuleName: wafDto.Action,
		Scope:    wafDto.Scope,
		ObjId:    wafDto.Id,
		ObjType:  waf.Waf_obj_type_accurate,
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
func (wafService *WafAccurateService) UpdateWafAccurates(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafAccurateDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWafAccurate(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	qWafRuleDto := waf.WafRuleDto{
		ObjId:   wafDto.Id,
		ObjType: waf.Waf_obj_type_accurate,
	}
	qWafRuleDtos, _ := wafService.wafRuleDao.GetWafRules(qWafRuleDto)

	if qWafRuleDtos == nil || len(qWafRuleDtos) < 1 {
		return result.Success()
	}

	wafRuleDto := waf.WafRuleDto{
		RuleId:   qWafRuleDtos[0].RuleId,
		GroupId:  wafDto.GroupId,
		RuleName: wafDto.Action,
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
func (wafService *WafAccurateService) DeleteWafAccurates(ctx iris.Context) result.ResultDto {
	var (
		err    error
		wafDto waf.WafAccurateDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWafAccurate(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}
	qWafRuleDto := waf.WafRuleDto{
		ObjId:   wafDto.Id,
		ObjType: waf.Waf_obj_type_cc,
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
