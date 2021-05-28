package controller

import (
	"github.com/obonobo/express-vpn-updater/server/app/util"
)

const (
	HealthcheckMessage = "All good in the hood"
)

var a *api = &api{}

type api struct {
	c *Controller
}

func (a *api) controller() Controller {
	logger.Println("Grabbing controller")
	if a.c == nil {
		controller := Default()
		a.c = &controller
	}
	return *a.c
}

// func Healthcheck(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	buf, _ := json.Marshal(map[string]interface{}{
// 		"message": "All good in the hood",
// 	})

// 	return events.APIGatewayProxyResponse{
// 		IsBase64Encoded: false,
// 		StatusCode:      200,
// 		Body:            string(buf),
// 	}, nil
// }

func Healthcheck(req util.Request) (util.Response, error) {
	return util.BasicMessage(HealthcheckMessage), nil
}

func Latest(req util.Request) (util.Response, error) {
	logger.Println("Calling API: Latest")
	logger.Println(a.controller())
	return a.controller().Latest(req), nil
}
