package initialize

import (
	"github.com/nzmxd/app-insight/core/analyzer"
	"github.com/nzmxd/app-insight/global"
)

func InitStaticAnalyzer() {
	switch global.Config.StaticAnalyzer.Use {
	case "jadx":
		global.Analyzer = &analyzer.JadxAnalyzer{Exec: global.Config.StaticAnalyzer.Exec, WorkDir: global.Config.StaticAnalyzer.Workdir}
	default:
		global.Analyzer = &analyzer.JadxAnalyzer{Exec: global.Config.StaticAnalyzer.Exec, WorkDir: global.Config.StaticAnalyzer.Workdir}
	}
}
