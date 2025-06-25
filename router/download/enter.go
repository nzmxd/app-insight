package download

import v1 "app-insight/api/v1"

type RouterGroup struct {
	DownloadTaskRouter
}

var (
	downloadTaskApi = v1.ApiGroupApp.DownloadTaskApiGroup.DownloadTaskApi
)
