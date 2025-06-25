package initialize

import (
	"github.com/hibiken/asynq"
	aGlobal "github.com/nzmxd/app-insight/global"
	gloabl "github.com/nzmxd/bserver/global"
)

func InitAsynqClient() {
	if !gloabl.CONFIG.System.UseRedis {
		return
	}
	aGlobal.AsynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: gloabl.CONFIG.Redis.Addr})
}
