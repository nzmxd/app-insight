package download

import (
	"app-insight/model/download"
	"app-insight/model/download/request"
	"github.com/gin-gonic/gin"
	"github.com/nzmxd/bserver/global"
	"github.com/nzmxd/bserver/model/common/response"
	"go.uber.org/zap"
)

type DownloadTaskApi struct{}

// AddDownloadTask 添加APP下载任务
// @Tags Download
// @Summary 添加APP下载任务
// @Accept application/json
// @Produce application/json
// @Param data body download.AppDownloadTask true "添加APP下载任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /download/addDownloadTask [post]
func (a *DownloadTaskApi) AddDownloadTask(c *gin.Context) {
	var appDownloadTask download.AppDownloadTask
	err := c.ShouldBindJSON(&appDownloadTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = downloadTaskService.CreateDownloadTask(&appDownloadTask)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// GetDownloadTaskById 根据ID查询下载任务
// @Tags Download
// @Summary 根据ID查询下载任务
// @Accept application/json
// @Produce application/json
// @Param id query uint true "根据ID查询下载任务"
// @Success 200 {object} response.Response{data=app.AppDownloadTask,msg=string} "查询成功"
// @Router /download/getDownloadTaskById [get]
func (a *DownloadTaskApi) GetDownloadTaskById(c *gin.Context) {
	id := c.Query("id")
	downloadTask, err := downloadTaskService.GetDownloadTask(id)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(downloadTask, c)

}

// GetDownloadStats 查询APP下载统计情况
// @Tags Download
// @Summary 查询APP下载统计情况
// @Accept application/json
// @Produce application/json
// @Param data query request.DownloadStatsSearch true "查询APP下载统计情况"
// @Success 200 {object} response.Response{data=app.DownloadStatsResult,msg=string} "查询成功"
// @Router /download/getDownloadStats [get]
func (a *DownloadTaskApi) GetDownloadStats(c *gin.Context) {
	var appDownloadTask request.DownloadStatsSearch
	err := c.ShouldBindQuery(&appDownloadTask)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	stats, err := downloadTaskService.GetDownloadStats(appDownloadTask)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(stats, c)
}

// GetDownloadTaskByAppId 根据AppId查询下载任务
// @Tags Download
// @Summary 根据AppId查询下载任务
// @Accept application/json
// @Produce application/json
// @Param appId query string true "根据AppId查询下载任务"
// @Success 200 {object} response.Response{data=download.AppDownloadTask,msg=string} "查询成功"
// @Router /download/getDownloadTaskByAppId [get]
func (a *DownloadTaskApi) GetDownloadTaskByAppId(c *gin.Context) {
	appId := c.Query("appId")
	downloadTask, err := downloadTaskService.GetDownloadTaskByAppId(appId)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(downloadTask, c)
}

// GetDownloadUrl 查询APP下载链接
// @Tags Download
// @Summary 查询APP下载链接
// @Accept application/json
// @Produce application/json
// @Param data query request.DownloadUrlSearch true "查询APP下载链接"
// @Success 200 {object} response.Response{data=app.DownloadStatsResult,msg=string} "查询成功"
// @Router /download/getDownloadUrl [get]
func (a *DownloadTaskApi) GetDownloadUrl(c *gin.Context) {
	var search request.DownloadUrlSearch
	err := c.ShouldBindQuery(&search)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	stats, err := downloadTaskService.GetDownloadUrl(search)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(stats, c)
}
