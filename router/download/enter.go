package download

import v1 "github.com/nzmxd/app-insight/api/v1"

type RouterGroup struct {
	DownloadTaskRouter
}

var (
	downloadTaskApi = v1.ApiGroupApp.DownloadTaskApiGroup.DownloadTaskApi
)
