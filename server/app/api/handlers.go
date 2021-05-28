package api

import "github.com/obonobo/express-vpn-updater/server/app/util"

func Healthcheck(req util.Request) (util.Response, error) {
	logger.LogApiCall("Healthcheck", a.healthcheckController())
	return a.healthcheckController().Healthcheck(req), nil
}

func Latest(req util.Request) (util.Response, error) {
	logger.LogApiCall("Latest", a.latestController())
	return a.latestController().Latest(req), nil
}
