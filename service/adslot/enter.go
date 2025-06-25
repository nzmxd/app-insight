package adslot

import (
	"errors"
	"github.com/nzmxd/app-insight/model/adslot"
)

type ServiceGroup struct {
	AdSlotPushService
	AdSlotPushLogService
	ScrapyAdmobCoreParamsService      CoreParamService[*adslot.ScrapyAdmobCoreParams]
	ScrapyApplovinCoreParamsService   CoreParamService[*adslot.ScrapyApplovinCoreParams]
	ScrapyChartboostCoreParamsService CoreParamService[*adslot.ScrapyChartboostCoreParams]
	ScrapyUnityadsCoreParamsService   CoreParamService[*adslot.ScrapyUnityadsCoreParams]
	ScrapyIronsourceCoreParamsService CoreParamService[*adslot.ScrapyIronsourceCoreParams]
	ScrapyVungleCoreParamsService     CoreParamService[*adslot.ScrapyVungleCoreParams]
	//ScrapyPangleCoreParamsService     CoreParamService[*adslot.ScrapyPangleCoreParams]
	//ScrapyYandexCoreParamsService     CoreParamService[*adslot.ScrapyYandexCoreParams]
}

var (
	DuplicateAdUnitErr     = errors.New("重复的广告位")   // 这个错误表示在线上广告位系统里面有但是在本地推送日志里面没有
	DuplicatePushAdUnitErr = errors.New("重复推送的广告位") // 这个错误表示在本地推送日志里面有
	InvalidAdUnitErr       = errors.New("无效的广告位")
	adUnitPushLogService   AdSlotPushLogService
)
