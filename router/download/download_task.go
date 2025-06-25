package download

import "github.com/gin-gonic/gin"

type DownloadTaskRouter struct{}

func (d *DownloadTaskRouter) InitDownloadTaskRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("download")
	appRouter.POST("addDownloadTask", downloadTaskApi.AddDownloadTask)
	appRouter.GET("getDownloadTaskById", downloadTaskApi.GetDownloadTaskById)
	appRouter.GET("getDownloadTaskByAppId", downloadTaskApi.GetDownloadTaskByAppId)
	appRouter.GET("getDownloadStats", downloadTaskApi.GetDownloadStats)
	appRouter.GET("getDownloadUrl", downloadTaskApi.GetDownloadUrl)
}
