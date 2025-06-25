package service

import (
	"app-insight/service/adslot"
	"app-insight/service/analysis"
	"app-insight/service/download"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	DownloadServiceGroup download.ServiceGroup
	AdSlotServiceGroup   adslot.ServiceGroup
	AnalysisServiceGroup analysis.ServiceGroup
}
