package initialize

import (
	"github.com/nzmxd/app-insight/scheduler"
	"github.com/nzmxd/bserver/core"
	"github.com/nzmxd/bserver/global"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Timer() {
	core.RegisterCronTask(
		`EnqueueDownloadTask`,
		"@every 2m",
		scheduler.EnqueueDownloadTask,
		"定时向任务队列中添加下载任务",
		cron.WithSeconds(),
	)

	core.RegisterCronTask(
		`SyncDownloadTask`,
		"@every 5m",
		scheduler.SyncDownloadTask,
		"同步ASO数据库中的安卓包名",
		cron.WithSeconds(),
	)

	core.RegisterCronTask(
		`EnqueueStaticAnalysis`,
		"@every 3m",
		scheduler.EnqueueStaticAnalysis,
		"定时向任务队列中添加分析任务",
		cron.WithSeconds(),
	)
	core.RegisterCronTask(
		`SyncStaticAnalysisTask`,
		"@every 5m",
		scheduler.SyncStaticAnalysisTask,
		"定时从下载任务表中同步下载完成的任务",
		cron.WithSeconds(),
	)

}

func ExecOnce() {
	err := scheduler.SyncDownloadTask()
	if err != nil {
		global.LOG.Error("exec SyncDownloadTask err", zap.Error(err))
	}
	err = scheduler.EnqueueDownloadTask()
	if err != nil {
		global.LOG.Error("exec EnqueueDownloadTask err", zap.Error(err))
	}
	err = scheduler.SyncStaticAnalysisTask()
	if err != nil {
		global.LOG.Error("exec SyncStaticAnalysisTask err", zap.Error(err))
	}
	err = scheduler.EnqueueStaticAnalysis()
	if err != nil {
		global.LOG.Error("exec EnqueueStaticAnalysis err", zap.Error(err))
	}
}
