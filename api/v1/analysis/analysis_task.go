package analysis

import (
	"github.com/gin-gonic/gin"
	"github.com/nzmxd/app-insight/model/analysis/request"
	"github.com/nzmxd/bserver/global"
	"github.com/nzmxd/bserver/model/common/response"
	"go.uber.org/zap"
)

type AnalysisTaskApi struct{}

// GetStaticStats 查询静态分析结果
// @Tags Analysis
// @Summary 查询静态分析结果
// @Accept application/json
// @Produce application/json
// @Param data query request.StaticAnalysisStatsSearch true "查询静态分析结果"
// @Success 200 {object} response.Response{msg=string} "查询成功"
// @Router /analysis/getStaticStats [get]
func (s *AnalysisTaskApi) GetStaticStats(c *gin.Context) {
	var search request.StaticAnalysisStatsSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	stats, err := staticAnalysisTaskService.GetStaticAnalysisStats(search)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(stats, c)
}
