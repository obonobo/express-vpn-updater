package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/obonobo/express-vpn-updater/server/app/controller"
)

func main() {
	lambda.Start(controller.Healthcheck)
}
