package v1

import (
	"app-insight/api/v1/adslot"
	"app-insight/api/v1/analysis"
	"app-insight/api/v1/download"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	DownloadTaskApiGroup download.ApiGroup
	AdSlotApiGroup       adslot.ApiGroup
	AnalysisTaskApiGroup analysis.ApiGroup
}
