package analysis

import "github.com/nzmxd/app-insight/service"

type ApiGroup struct {
	AnalysisTaskApi
}

var (
	staticAnalysisTaskService = service.ServiceGroupApp.AnalysisServiceGroup.AppStaticAnalysisTaskService
)
