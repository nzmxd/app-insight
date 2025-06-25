package adslot

import "app-insight/service"

type ApiGroup struct {
	AdSlotPushApi
}

var (
	adUnitPushService                 = service.ServiceGroupApp.AdSlotServiceGroup.AdSlotPushService
	adUnitPushLogService              = service.ServiceGroupApp.AdSlotServiceGroup.AdSlotPushLogService
	scrapyAdmobCoreParamsService      = service.ServiceGroupApp.AdSlotServiceGroup.ScrapyAdmobCoreParamsService
	scrapyApplovinCoreParamsService   = service.ServiceGroupApp.AdSlotServiceGroup.ScrapyApplovinCoreParamsService
	scrapyChartboostCoreParamsService = service.ServiceGroupApp.AdSlotServiceGroup.ScrapyChartboostCoreParamsService
	scrapyUnityadsCoreParamsService   = service.ServiceGroupApp.AdSlotServiceGroup.ScrapyUnityadsCoreParamsService
	scrapyIronsourceCoreParamsService = service.ServiceGroupApp.AdSlotServiceGroup.ScrapyIronsourceCoreParamsService
	scrapyVungleCoreParamsService     = service.ServiceGroupApp.AdSlotServiceGroup.ScrapyVungleCoreParamsService
)
