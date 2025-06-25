package service

import (
	"github.com/nzmxd/app-insight/service/adslot"
	"github.com/nzmxd/app-insight/service/analysis"
	"github.com/nzmxd/app-insight/service/download"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	DownloadServiceGroup download.ServiceGroup
	AdSlotServiceGroup   adslot.ServiceGroup
	AnalysisServiceGroup analysis.ServiceGroup
}
