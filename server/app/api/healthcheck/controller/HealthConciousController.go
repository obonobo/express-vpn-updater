package controller

import "github.com/obonobo/express-vpn-updater/server/app/util"

const (
	HealthcheckMessage = "All good in the hood"
)

type healthConciousController struct{}

func New() Controller {
	return &healthConciousController{}
}

func (c *healthConciousController) Healthcheck(req util.Request) util.Response {
	return util.BasicMessage(HealthcheckMessage)
}
