package core

import (
	"github.com/hibiken/asynq"
	aGlobal "github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/service"
	"github.com/nzmxd/bserver/global"
	"go.uber.org/zap"
)

func RunAsynqAnalyzerServer() {
	if !global.CONFIG.System.UseRedis {
		return
	}
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: global.CONFIG.Redis.Addr},
		asynq.Config{
			Concurrency: aGlobal.Config.StaticAnalyzer.Worker,
			Queues: map[string]int{
				"analysis": 1,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.Handle(aGlobal.StaticAnalysisTask, &service.ServiceGroupApp.AnalysisServiceGroup.AppStaticAnalysisTaskService)
	if err := srv.Start(mux); err != nil {
		global.LOG.Panic("could not run server", zap.Error(err))
	}
}
