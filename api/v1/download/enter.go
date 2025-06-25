package download

import "github.com/nzmxd/app-insight/service"

type ApiGroup struct {
	DownloadTaskApi
}

var (
	downloadTaskService = service.ServiceGroupApp.DownloadServiceGroup.DownloadTaskService
)
