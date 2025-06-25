package scheduler

import (
	"fmt"
	"github.com/nzmxd/app-insight/model/analysis"
	"github.com/nzmxd/bserver/global"
	"go.uber.org/zap"
	"time"
)

func EnqueueStaticAnalysis() error {
	var analysisQueueCount int64
	err := global.DB.Model(&analysis.AppAnalysisTask{}).Where("file_analysis_status = ?", analysis.StatusQueued).Count(&analysisQueueCount).Error
	if err != nil {
		global.LOG.Error("查询下载队列任务数失败", zap.Error(err))
		return err
	}
	if analysisQueueCount > 10 {
		global.LOG.Info(fmt.Sprintf("当前分析任务队列数为%d,无需添加新任务", analysisQueueCount))
		return nil
	}
	var analysisTasks []analysis.AppAnalysisTask
	err = global.DB.Where("file_analysis_status = ?", analysis.StatusPending).Limit(100).Find(&analysisTasks).Error
	if err != nil {
		global.LOG.Error("查询待分析任务失败", zap.Error(err))
		return err
	}
	if len(analysisTasks) == 0 {
		global.LOG.Info("当前没有待添加的分析任务")
		return nil
	}
	now := time.Now()
	var successIDs, failedIDs []int64
	for _, task := range analysisTasks {
		err = staticAnalysisTaskService.EnqueueStaticAnalysisTask(analysis.AppAnalysisTaskPayload{
			AppID:       task.AppID,
			VersionCode: task.VersionCode,
		})
		if err != nil {
			failedIDs = append(failedIDs, task.ID)
			global.LOG.Error("任务添加到分析队列失败", zap.Any("task", task), zap.Error(err))
			continue
		}
		successIDs = append(successIDs, task.ID)
	}

	if len(successIDs) > 0 {
		_ = global.DB.Model(&analysis.AppAnalysisTask{}).Where("id IN ?", successIDs).
			Updates(map[string]interface{}{
				"file_analysis_started_at": now,
				"file_analysis_status":     analysis.StatusQueued,
			}).Error
	}
	return nil
}
