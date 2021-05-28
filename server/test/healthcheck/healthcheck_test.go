package healthcheck

import (
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/controller"
	"github.com/obonobo/express-vpn-updater/server/app/util"
	"github.com/obonobo/express-vpn-updater/server/test/testutils"
	"github.com/stretchr/testify/assert"
)

const (
	receivedError      = "received error: %v"
	errorInResponse    = "controller.Healthcheck may not return an error"
	badResponseBody    = "The body of the healthcheck must be present and must be a valid JSON formatted string"
	badResponseHeaders = "got bad headers in the health check response"
)

var (
	desiredHeaders = util.DefaultHeaders()
	desiredBody    = map[string]interface{}{"message": controller.HealthcheckMessage}
)

func TestHealthcheck(t *testing.T) {
	response := requestAndAssert(t)
	assertHeaders(t, response)
	assertBody(t, response)
}

func requestAndAssert(t *testing.T) util.Response {
	response, err := controller.Healthcheck(util.Request{})
	if err != nil {
		assert.Failf(t, errorInResponse, receivedError, err)
	}
	return response
}

func assertHeaders(t *testing.T, response util.Response) {
	testutils.AssertHeaders(t, desiredHeaders, response.Headers)
}

func assertBody(t *testing.T, response util.Response) {
	testutils.AssertBody(t, desiredBody, response.Body)
}
