package initialize

import (
	aGlobal "app-insight/global"
	"github.com/hibiken/asynq"
	gloabl "github.com/nzmxd/bserver/global"
)

func InitAsynqClient() {
	if !gloabl.CONFIG.System.UseRedis {
		return
	}
	aGlobal.AsynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: gloabl.CONFIG.Redis.Addr})
}
