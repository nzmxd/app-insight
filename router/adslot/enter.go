package adslot

import v1 "app-insight/api/v1"

type RouterGroup struct {
	AdSlotPushRouter
}

var (
	adSlotPushApi = v1.ApiGroupApp.AdSlotApiGroup.AdSlotPushApi
)
