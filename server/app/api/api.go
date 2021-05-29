package api

import (
	health "github.com/obonobo/express-vpn-updater/server/app/api/healthcheck/controller"
	latest "github.com/obonobo/express-vpn-updater/server/app/api/latest/controller"
	"github.com/obonobo/express-vpn-updater/server/app/config"
)

var (
	a      = &api{}
	logger = config.Get().Logger()
)

type api struct {
	lc *latest.Controller
	hc *health.Controller
}

func (a *api) latestController() latest.Controller {
	if a.lc == nil {
		controller := latest.Default()
		a.lc = &controller
	}
	return *a.lc
}

func (a *api) healthcheckController() health.Controller {
	if a.hc == nil {
		controller := health.New()
		a.hc = &controller
	}
	return *a.hc
}
