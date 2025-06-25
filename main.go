package main

import (
	"flag"
	aCore "github.com/nzmxd/app-insight/core"
	"github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/initialize"
	"github.com/nzmxd/bserver/core"
)

func main() {
	var appConfig string
	flag.StringVar(&appConfig, "d", "", "app config file")
	initialize.InitRouters()
	core.InitializeSystem()
	core.NViper(appConfig, &global.Config)
	initialize.InitOthers()
	aCore.RunAsynqDownloaderServer()
	aCore.RunAsynqAnalyzerServer()
	core.RunServer()
}
