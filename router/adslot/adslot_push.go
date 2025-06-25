package adslot

import "github.com/gin-gonic/gin"

type AdSlotPushRouter struct{}

func (d *AdSlotPushRouter) InitAdSlotRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("adslot")
	appRouter.GET("getPushStats", adSlotPushApi.GetPushStats)
	appRouter.POST("pushApplovin", adSlotPushApi.PushApplovin)
	appRouter.POST("pushAdmob", adSlotPushApi.PushAdmob)
}
