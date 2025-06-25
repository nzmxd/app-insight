package main

import (
	aCore "app-insight/core"
	"app-insight/global"
	"app-insight/initialize"
	"flag"
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
