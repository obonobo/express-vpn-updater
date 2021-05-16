package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/obonobo/express-vpn-updater/util"
)

func main() {
	lambda.Start(Handler)
}

func Handler(req util.Request) (util.Response, error) {
	return util.BasicMessage("All good in the hood"), nil
}
