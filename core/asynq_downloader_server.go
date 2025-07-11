package core

import (
	"github.com/hibiken/asynq"
	aGlobal "github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/service"
	"github.com/nzmxd/bserver/global"
	"go.uber.org/zap"
)

func RunAsynqDownloaderServer() {
	if !global.CONFIG.System.UseRedis {
		return
	}
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: global.CONFIG.Redis.Addr},
		asynq.Config{
			Concurrency: aGlobal.Config.Downloader.Worker,
			Queues: map[string]int{
				"download": 1,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.Handle(aGlobal.DownloadTask, &service.ServiceGroupApp.DownloadServiceGroup.DownloadTaskService)
	if err := srv.Start(mux); err != nil {
		global.LOG.Panic("could not run server", zap.Error(err))
	}
}
