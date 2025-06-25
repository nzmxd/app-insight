package adslot

import (
	"app-insight/model/adslot"
	"app-insight/model/adslot/request"
	"github.com/gin-gonic/gin"
	"github.com/nzmxd/bserver/global"
	"github.com/nzmxd/bserver/model/common/response"
	"go.uber.org/zap"
)

type AdSlotPushApi struct{}

// GetPushStats 查询广告位推送统计
// @Tags AdSlot
// @Summary 查询广告位推送统计
// @Accept application/json
// @Produce application/json
// @Param data query request.AdSlotPushLogSearch true "广告位推送统计查询条件"
// @Success 200 {object} response.Response{data=[]ad.AdUnitPushLogStatsResult,msg=string} "查询成功"
// @Router /adslot/getPushStats [get]
func (a *AdSlotPushApi) GetPushStats(c *gin.Context) {
	var search request.AdSlotPushLogSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	stats, err := adUnitPushLogService.GetAdUnitPushLogStats(search)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(stats, c)
}

// PushAdmob Admob广告位推送
// @Tags AdSlot
// @Summary Admob广告位推送
// @Accept application/json
// @Produce application/json
// @Param data body request.ScrapyAdmobCoreParamsRequest true "Admob广告位推送"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adslot/pushAdmob [post]
func (a *AdSlotPushApi) PushAdmob(c *gin.Context) {
	var params request.ScrapyAdmobCoreParamsRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	admobCoreParams, err := scrapyApplovinCoreParamsService.ConvertAdmob(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scrapyAdmobCoreParamsService.Create(admobCoreParams)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// PushApplovin Applovin广告位推送
// @Tags AdSlot
// @Summary Applovin广告位推送
// @Accept application/json
// @Produce application/json
// @Param data body request.ScrapyApplovinCoreParamsRequest true "Applovin广告位推送"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adslot/pushApplovin [post]
func (a *AdSlotPushApi) PushApplovin(c *gin.Context) {
	var params request.ScrapyApplovinCoreParamsRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	applovinCoreParams, err := scrapyApplovinCoreParamsService.ConvertApplovin(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scrapyApplovinCoreParamsService.Create(applovinCoreParams)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// PushCharboost Charboost广告位推送
// @Tags AdSlot
// @Summary Charboost广告位推送
// @Accept application/json
// @Produce application/json
// @Param data body adslot.ScrapyChartboostCoreParams true "Charboost广告位推送"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adslot/pushCharboost [post]
func (a *AdSlotPushApi) PushCharboost(c *gin.Context) {
	var params adslot.ScrapyChartboostCoreParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scrapyChartboostCoreParamsService.Create(&params)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// PushUnity Unity广告位推送
// @Tags AdSlot
// @Summary Unity广告位推送
// @Accept application/json
// @Produce application/json
// @Param data body adslot.ScrapyUnityadsCoreParams true "Unity广告位推送"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adslot/pushUnity [post]
func (a *AdSlotPushApi) PushUnity(c *gin.Context) {
	var params adslot.ScrapyUnityadsCoreParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scrapyUnityadsCoreParamsService.Create(&params)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// PushIronsource Ironsource广告位推送
// @Tags AdSlot
// @Summary Ironsource广告位推送
// @Accept application/json
// @Produce application/json
// @Param data body adslot.ScrapyIronsourceCoreParams true "Ironsource广告位推送"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adslot/pushIronsource [post]
func (a *AdSlotPushApi) PushIronsource(c *gin.Context) {
	var params adslot.ScrapyIronsourceCoreParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scrapyIronsourceCoreParamsService.Create(&params)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// PushVungle Vungle广告位推送
// @Tags AdSlot
// @Summary Vungle广告位推送
// @Accept application/json
// @Produce application/json
// @Param data body adslot.ScrapyVungleCoreParams true "Vungle广告位推送"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adslot/pushVungle [get]
func (a *AdSlotPushApi) PushVungle(c *gin.Context) {
	var params adslot.ScrapyVungleCoreParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scrapyVungleCoreParamsService.Create(&params)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
