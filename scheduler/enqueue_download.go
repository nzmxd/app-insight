package scheduler

import (
	"app-insight/model/download"
	"fmt"
	"github.com/nzmxd/bserver/global"
	"go.uber.org/zap"
	"time"
)

// EnqueueDownloadTask 定时从AppDownloadTask向下载队列中添加下载任务
func EnqueueDownloadTask() error {
	err := downloadTaskService.UpdateTimeoutTasks()
	if err != nil {
		global.LOG.Error("更新超时任务失败", zap.Error(err))
	}
	var downloadQueueCount int64
	err = global.DB.Model(&download.AppDownloadTask{}).Where("status = ?", download.StatusQueued).Count(&downloadQueueCount).Error
	if err != nil {
		global.LOG.Error("查询下载队列任务数失败", zap.Error(err))
		return err
	}
	if downloadQueueCount > 10 {
		global.LOG.Info(fmt.Sprintf("当前下载任务队列数为%d,无需添加新任务", downloadQueueCount))
		return nil
	}
	var downloadTasks []download.AppDownloadTask
	err = global.DB.Where("status = ?", download.StatusPending).Limit(100).Find(&downloadTasks).Error
	if err != nil {
		global.LOG.Error("查询待下载任务失败", zap.Error(err))
		return err
	}
	if len(downloadTasks) == 0 {
		global.LOG.Info("当前没有待添加的下载任务")
		return nil
	}
	now := time.Now()
	var successIDs, failedIDs []int64
	for _, task := range downloadTasks {
		err = downloadTaskService.EnqueueDownloadTask(task)
		if err != nil {
			failedIDs = append(failedIDs, task.ID)
			global.LOG.Error("任务添加到下载队列失败", zap.Any("task", task), zap.Error(err))
			continue
		}
		successIDs = append(successIDs, task.ID)
	}

	if len(successIDs) > 0 {
		_ = global.DB.Model(&download.AppDownloadTask{}).Where("id IN ?", successIDs).
			Updates(map[string]interface{}{
				"started_at": now,
				"status":     download.StatusQueued,
			}).Error
	}
	return nil
}
