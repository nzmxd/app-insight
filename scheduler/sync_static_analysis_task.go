package scheduler

import (
	"github.com/nzmxd/app-insight/model/analysis"
	"github.com/nzmxd/bserver/global"
	"go.uber.org/zap"
)

func SyncStaticAnalysisTask() error {
	var analysisPendingCount int64
	err := global.DB.Model(&analysis.AppAnalysisTask{}).Where("file_analysis_status = ?", analysis.StatusPending).Count(&analysisPendingCount).Error
	if err != nil {
		global.LOG.Error("查询待分析任务数失败", zap.Error(err))
		return err
	}
	if analysisPendingCount > 10 {
		global.LOG.Info("当前分析任务数大于指定值无需添加分析任务", zap.Int64("analysis_pending_count", analysisPendingCount))
		return nil
	}
	pendingTasks, err := analysisTaskService.GetAnalysisTaskFormDownloadTask(100)
	if err != nil {
		global.LOG.Error("获取分析任务失败", zap.Error(err))
		return err
	}
	global.LOG.Info("成功添加文件分析任务", zap.Int("count", len(pendingTasks)))
	return nil
}
