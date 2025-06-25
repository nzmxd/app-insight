package analysis

import "app-insight/service"

type ApiGroup struct {
	AnalysisTaskApi
}

var (
	staticAnalysisTaskService = service.ServiceGroupApp.AnalysisServiceGroup.AppStaticAnalysisTaskService
)
