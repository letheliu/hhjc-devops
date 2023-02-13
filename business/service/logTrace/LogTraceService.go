package logTraceService

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/letheliu/hhjc-devops/business/dao/logTraceAnnotationsDao"
	"github.com/letheliu/hhjc-devops/business/dao/logTraceDao"
	"github.com/letheliu/hhjc-devops/business/dao/logTraceDbDao"
	"github.com/letheliu/hhjc-devops/business/dao/logTraceParamDao"
	"github.com/letheliu/hhjc-devops/common/seq"
	"github.com/letheliu/hhjc-devops/common/utils"
	"github.com/letheliu/hhjc-devops/entity/dto/log"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"strconv"
)

type LogTraceService struct {
	logTraceDao            logTraceDao.LogTraceDao
	logTraceParamDao       logTraceParamDao.LogTraceParamDao
	logTraceDbDao          logTraceDbDao.LogTraceDbDao
	logTraceAnnotationsDao logTraceAnnotationsDao.LogTraceAnnotationsDao
}

// get db link
// all db by this user
func (logTraceService *LogTraceService) GetLogTraceAll(LogTraceDto log.LogTraceDto) ([]*log.LogTraceDto, error) {
	var (
		err          error
		LogTraceDtos []*log.LogTraceDto
	)

	LogTraceDtos, err = logTraceService.logTraceDao.GetLogTraces(LogTraceDto)
	if err != nil {
		return nil, err
	}

	return LogTraceDtos, nil

}

/*
*
查询 系统信息
*/
func (logTraceService *LogTraceService) GetLogTraces(ctx iris.Context) result.ResultDto {
	var (
		err          error
		page         int64
		row          int64
		total        int64
		logTraceDto  = log.LogTraceDto{}
		logTraceDtos []*log.LogTraceDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceDto.Row = row * page

	logTraceDto.Page = (page - 1) * row

	logTraceDto.Name = ctx.URLParam("name")

	logTraceDto.TraceId = ctx.URLParam("traceId")

	total, err = logTraceService.logTraceDao.GetLogTraceCount(logTraceDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceDtos, err = logTraceService.logTraceDao.GetLogTraces(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDtos, total, row)

}

/*
*
查询 系统信息
*/
func (logTraceService *LogTraceService) GetLogTraceDetail(ctx iris.Context) result.ResultDto {
	var (
		err          error
		page         int64
		row          int64
		total        int64
		logTraceDto  = log.LogTraceDto{}
		logTraceDtos []*log.LogTraceDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceDto.Row = row * page

	logTraceDto.Page = (page - 1) * row

	logTraceDto.Name = ctx.URLParam("name")

	logTraceDto.TraceId = ctx.URLParam("traceId")

	total, err = logTraceService.logTraceDao.GetLogTraceCount(logTraceDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceDtos, err = logTraceService.logTraceDao.GetLogTraces(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	logTraceDtos = logTraceService.addAnn(logTraceDtos)

	return result.SuccessData(logTraceDtos, total, row)

}

/*
*
保存 系统信息
*/
func (logTraceService *LogTraceService) SaveLogTraces(param string) result.ResultDto {
	var (
		err                    error
		logTraceDto            log.LogTraceDto
		logTraceDataDto        log.LogTraceDataDto
		logTraceAnnotationsDto log.LogTraceAnnotationsDto
		crTimestame            int64
		csTimestame            int64
	)
	json.Unmarshal([]byte(param), &logTraceDataDto)

	//object convert
	json.Unmarshal([]byte(param), &logTraceDto)

	//logTraceDto.Id = seq.Generator()
	//LogTraceDto.Path = filepath.Join(curDest, fileHeader.Filename)

	if logTraceDataDto.Annotations == nil || len(logTraceDataDto.Annotations) < 1 {
		return result.Error("未包含Annotations")
	}

	logTraceDto.ServiceName = logTraceDataDto.Annotations[0].Endpoint.ServiceName
	logTraceDto.Ip = logTraceDataDto.Annotations[0].Endpoint.Ip
	logTraceDto.Port = logTraceDataDto.Annotations[0].Endpoint.Port
	logTraceDto.Duration = 0
	//compute Duration cr - cs
	if len(logTraceDataDto.Annotations) > 0 {

		for _, annotation := range logTraceDataDto.Annotations {
			if annotation.Value == "cr" {
				crTimestame = annotation.Timestamp
			}

			if annotation.Value == "cs" {
				csTimestame = annotation.Timestamp
			}
		}

		if crTimestame != 0 && csTimestame != 0 {
			logTraceDto.Duration = crTimestame - csTimestame
		}
	}

	err = logTraceService.logTraceDao.SaveLogTrace(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	for _, annotation := range logTraceDataDto.Annotations {
		logTraceAnnotationsDto = log.LogTraceAnnotationsDto{
			Id:          seq.Generator(),
			SpanId:      logTraceDto.Id,
			ServiceName: annotation.Endpoint.ServiceName,
			Ip:          annotation.Endpoint.Ip,
			Port:        annotation.Endpoint.Port,
			Value:       annotation.Value,
			Timestamp:   annotation.Timestamp,
		}
		err = logTraceService.logTraceAnnotationsDao.SaveLogTraceAnnotations(logTraceAnnotationsDto)
		if err != nil {
			return result.Error(err.Error())
		}
	}

	//save log

	logTraceParamDto := logTraceDataDto.Param

	if logTraceParamDto != nil && !utils.IsEmpty(logTraceParamDto.ReqParam) {
		logTraceParamDto.Id = seq.Generator()
		logTraceParamDto.SpanId = logTraceDto.Id
		logTraceService.logTraceParamDao.SaveLogTraceParam(*logTraceParamDto)
	}

	logTraceDbDtos := logTraceDataDto.Dbs

	if logTraceDbDtos != nil && len(logTraceDbDtos) > 0 {
		for _, logTraceDbDto := range logTraceDbDtos {
			logTraceDbDto.Id = seq.Generator()
			logTraceDbDto.SpanId = logTraceDto.Id
			logTraceService.logTraceDbDao.SaveLogTraceDb(*logTraceDbDto)
		}
	}

	return result.SuccessData(logTraceDto)
}

/*
*
修改 系统信息
*/
func (logTraceService *LogTraceService) UpdateLogTraces(ctx iris.Context) result.ResultDto {
	var (
		err         error
		logTraceDto log.LogTraceDto
	)
	if err = ctx.ReadJSON(&logTraceDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = logTraceService.logTraceDao.UpdateLogTrace(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDto)

}

/*
*
删除 系统信息
*/
func (logTraceService *LogTraceService) DeleteLogTraces(ctx iris.Context) result.ResultDto {
	var (
		err         error
		logTraceDto log.LogTraceDto
	)
	if err = ctx.ReadJSON(&logTraceDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = logTraceService.logTraceDao.DeleteLogTrace(logTraceDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDto)

}

func (logTraceService *LogTraceService) addAnn(dtos []*log.LogTraceDto) []*log.LogTraceDto {

	for _, trace := range dtos {

		anno := log.LogTraceAnnotationsDto{
			SpanId: trace.Id,
		}

		annos, err := logTraceService.logTraceAnnotationsDao.GetLogTraceAnnotationss(anno)
		if err != nil {
			continue
		}
		trace.Annotations = annos
	}

	return dtos

}

func (logTraceService *LogTraceService) GetLogTraceParam(ctx iris.Context) interface{} {
	var (
		err               error
		page              int64
		row               int64
		total             int64
		logTraceParamDto  = log.LogTraceParamDto{}
		logTraceParamDtos []*log.LogTraceParamDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceParamDto.Row = row * page

	logTraceParamDto.Page = (page - 1) * row

	logTraceParamDto.SpanId = ctx.URLParam("spanId")

	total, err = logTraceService.logTraceParamDao.GetLogTraceParamCount(logTraceParamDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceParamDtos, err = logTraceService.logTraceParamDao.GetLogTraceParams(logTraceParamDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceParamDtos, total, row)
}

func (logTraceService *LogTraceService) GetLogTraceDb(ctx iris.Context) interface{} {
	var (
		err            error
		page           int64
		row            int64
		total          int64
		logTraceDbDto  = log.LogTraceDbDto{}
		logTraceDbDtos []*log.LogTraceDbDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	logTraceDbDto.Row = row * page

	logTraceDbDto.Page = (page - 1) * row

	logTraceDbDto.SpanId = ctx.URLParam("spanId")
	logTraceDbDto.ServiceName = ctx.URLParam("serviceName")
	logTraceDbDto.TraceId = ctx.URLParam("traceId")

	total, err = logTraceService.logTraceDbDao.GetLogTraceDbCount(logTraceDbDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	logTraceDbDtos, err = logTraceService.logTraceDbDao.GetLogTraceDbs(logTraceDbDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(logTraceDbDtos, total, row)
}
