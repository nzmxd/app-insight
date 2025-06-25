package analysis

import v1 "github.com/nzmxd/app-insight/api/v1"

type RouterGroup struct {
	AnalysisTaskRouter
}

var (
	analysisTaskApi = v1.ApiGroupApp.AnalysisTaskApiGroup.AnalysisTaskApi
)
