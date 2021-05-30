package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/obonobo/express-vpn-updater/server/app/api"
)

func main() {
	lambda.Start(api.Latest)
}
