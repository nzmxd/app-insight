package scheduler

import (
	"app-insight/service"
)

var (
	downloadTaskService       = service.ServiceGroupApp.DownloadServiceGroup.DownloadTaskService
	analysisTaskService       = service.ServiceGroupApp.AnalysisServiceGroup.AppAnalysisTaskService
	staticAnalysisTaskService = service.ServiceGroupApp.AnalysisServiceGroup.AppStaticAnalysisTaskService
)
