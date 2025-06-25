package initialize

import (
	"app-insight/core/analyzer"
	"app-insight/global"
)

func InitStaticAnalyzer() {
	switch global.Config.StaticAnalyzer.Use {
	case "jadx":
		global.Analyzer = &analyzer.JadxAnalyzer{Exec: global.Config.StaticAnalyzer.Exec, WorkDir: global.Config.StaticAnalyzer.Workdir}
	default:
		global.Analyzer = &analyzer.JadxAnalyzer{Exec: global.Config.StaticAnalyzer.Exec, WorkDir: global.Config.StaticAnalyzer.Workdir}
	}
}
