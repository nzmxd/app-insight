package router

import (
	"github.com/nzmxd/app-insight/router/adslot"
	"github.com/nzmxd/app-insight/router/analysis"
	"github.com/nzmxd/app-insight/router/download"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Download download.RouterGroup
	AdSlot   adslot.RouterGroup
	Analysis analysis.RouterGroup
}
