package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/obonobo/express-vpn-updater/scrape/controller"
	"github.com/obonobo/express-vpn-updater/util"
)

func main() {
	lambda.Start(grabLatestLink)
}

func grabLatestLink(req util.Request) (util.Response, error) {
	c := controller.Default()
	response := c.ScrapeAndRespond(req)
	return response, nil
}
