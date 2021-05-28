package controller

import (
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/util"
)

func TestHealthcheck(t *testing.T) {
	AssertHealthTest(t, func(req util.Request) (util.Response, error) {
		return New().Healthcheck(util.Request{}), nil
	})
}
