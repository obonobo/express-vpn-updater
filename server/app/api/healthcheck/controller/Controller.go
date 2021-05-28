package controller

import "github.com/obonobo/express-vpn-updater/server/app/util"

type Controller interface {
	Healthcheck(req util.Request) util.Response
}
