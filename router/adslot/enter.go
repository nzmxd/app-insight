package adslot

import v1 "github.com/nzmxd/app-insight/api/v1"

type RouterGroup struct {
	AdSlotPushRouter
}

var (
	adSlotPushApi = v1.ApiGroupApp.AdSlotApiGroup.AdSlotPushApi
)
