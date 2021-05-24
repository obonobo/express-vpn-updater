package controller

import "github.com/obonobo/express-vpn-updater/server/app/util"

type Controller interface {
	Latest(req util.Request) util.Response
	UpdateCache(req util.Request, redirect bool) util.Response
	CachedResponse(req util.Request, params *QueryParams) util.Response
}
