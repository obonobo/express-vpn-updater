package api

import "github.com/obonobo/express-vpn-updater/server/app/util"

func Healthcheck(request util.Request) (util.Response, error) {
	logger.LogApiCall("Healthcheck", a.healthcheckController())
	resp := logger.LogRequestAndResponse(request, a.healthcheckController().Healthcheck)
	return resp, nil
}

func Latest(req util.Request) (util.Response, error) {
	logger.LogApiCall("Latest", a.latestController())
	resp := logger.LogRequestAndResponse(req, a.latestController().Latest)
	return resp, nil
}
