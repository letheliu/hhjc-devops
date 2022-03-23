package wafService

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/business/dao/wafDao"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/waf"
	"strconv"
)

type WafService struct {
	wafDao wafDao.WafDao
}

// get db link
// all db by this user
func (wafService *WafService) GetWafAll(WafDto waf.WafDto) ([]*waf.WafDto, error) {
	var (
		err        error
		WafDtos []*waf.WafDto
	)

	WafDtos, err = wafService.wafDao.GetWafs(WafDto)
	if err != nil {
		return nil, err
	}

	return WafDtos, nil

}

/**
查询 系统信息
*/
func (wafService *WafService) GetWafs(ctx iris.Context) result.ResultDto {
	var (
		err        error
		page       int64
		row        int64
		total      int64
		wafDto  = waf.WafDto{}
		wafDtos []*waf.WafDto
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

	total, err = wafService.wafDao.GetWafCount(wafDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	wafDtos, err = wafService.wafDao.GetWafs(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDtos, total, row)

}

/**
保存 系统信息
*/
func (wafService *WafService) SaveWafs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}
	wafDto.WafId = seq.Generator()
	//WafDto.Path = filepath.Join(curDest, fileHeader.Filename)

	err = wafService.wafDao.SaveWaf(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
修改 系统信息
*/
func (wafService *WafService) UpdateWafs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	//wafDto.Id = ctx.FormValue("id")

	//wafDto.Name = ctx.FormValue("name")

	err = wafService.wafDao.UpdateWaf(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}

/**
删除 系统信息
*/
func (wafService *WafService) DeleteWafs(ctx iris.Context) result.ResultDto {
	var (
		err       error
		wafDto waf.WafDto
	)
	if err = ctx.ReadJSON(&wafDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = wafService.wafDao.DeleteWaf(wafDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(wafDto)

}
