package controller

import (
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/controller/queryparams"
	"github.com/obonobo/express-vpn-updater/server/app/util"
)

// A controller that provides 3 simple APIs for accessing the cache and
// webscraping infrastructure of the application
type Controller interface {

	// Retrieve the latest expressvpn .deb package from cache (and request a
	// subsequent updating of the cache)
	Latest(req util.Request) util.Response

	// Retrieve the latest expressvpn .deb package and update the package stored
	// in cache
	UpdateCache(req util.Request, redirect bool) util.Response

	// Retrieve the latest expressvpn .deb package from cache
	CachedResponse(req util.Request, params *queryparams.QueryParams) util.Response
}
