package initialize

import (
	"github.com/nzmxd/app-insight/router"
	"github.com/nzmxd/bserver/initialize"
)

func InitRouters() {
	initialize.AddRouters(
		router.RouterGroupApp.Download.InitDownloadTaskRouter,
		router.RouterGroupApp.AdSlot.InitAdSlotRouter,
		router.RouterGroupApp.Analysis.InitAnalysisTaskRouter,
	)
}
