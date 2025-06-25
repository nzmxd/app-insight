package download

import "app-insight/service"

type ApiGroup struct {
	DownloadTaskApi
}

var (
	downloadTaskService = service.ServiceGroupApp.DownloadServiceGroup.DownloadTaskService
)
