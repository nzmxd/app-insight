package v1

import (
	"github.com/nzmxd/app-insight/api/v1/adslot"
	"github.com/nzmxd/app-insight/api/v1/analysis"
	"github.com/nzmxd/app-insight/api/v1/download"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	DownloadTaskApiGroup download.ApiGroup
	AdSlotApiGroup       adslot.ApiGroup
	AnalysisTaskApiGroup analysis.ApiGroup
}
