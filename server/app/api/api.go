package api

import (
	healthyController "github.com/obonobo/express-vpn-updater/server/app/api/healthcheck/controller"
	latest "github.com/obonobo/express-vpn-updater/server/app/api/latest/controller"
	"github.com/obonobo/express-vpn-updater/server/app/config"
)

var (
	a      = &api{}
	logger = config.Get().Logger()
)

type api struct {
	lc *latest.Controller
	hc *healthyController.Controller
}

func (a *api) latestController() latest.Controller {
	if a.lc == nil {
		controller := latest.Default()
		a.lc = &controller
	}
	return *a.lc
}

func (a *api) healthcheckController() healthyController.Controller {
	if a.hc == nil {
		controller := healthyController.New()
		a.hc = &controller
	}
	return *a.hc
}
