package analysis

import "github.com/gin-gonic/gin"

type AnalysisTaskRouter struct{}

func (d *AnalysisTaskRouter) InitAnalysisTaskRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("analysis")
	appRouter.GET("getStaticStats", analysisTaskApi.GetStaticStats)
}
