package router

import (
	"back-end/router/expand"
	"back-end/router/system"
)

type AppRouterGroup struct {
	System system.SystemRouteGroup
	Expand expand.AppExpandRouterGroup
}

var AppRoute = new(AppRouterGroup)
