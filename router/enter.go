package router

import (
	"app-insight/router/adslot"
	"app-insight/router/analysis"
	"app-insight/router/download"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Download download.RouterGroup
	AdSlot   adslot.RouterGroup
	Analysis analysis.RouterGroup
}
