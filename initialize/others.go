package initialize

import (
	aGlobal "github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/service"
	"github.com/nzmxd/bserver/global"
	"go.uber.org/zap"
)

var (
	appSdkMatchRule = service.ServiceGroupApp.AnalysisServiceGroup.AppSdkMatchRuleService
)

func InitOthers() {
	Timer()
	InitDb()
	InitDownloader()
	InitStaticAnalyzer()
	InitAsynqClient()
	InitMatchRules()
	go ExecOnce()
}

func InitDb() {
	aGlobal.SpRawDB = global.MustGetGlobalDBByDBName("sp_raw")
	aGlobal.AppRankDB = global.MustGetGlobalDBByDBName("apprank")
	aGlobal.AppRankOnlineDB = global.MustGetGlobalDBByDBName("appranko")

}

func InitMatchRules() {
	rules, err := appSdkMatchRule.GetAllAppSdkMatchRule()
	if err != nil {
		global.LOG.Error("匹配规则加载失败", zap.Error(err))
		return
	}
	aGlobal.Analyzer.LoadRules(rules)
	global.LOG.Info("成功加载匹配规则", zap.Int("rules", len(rules)))
}
