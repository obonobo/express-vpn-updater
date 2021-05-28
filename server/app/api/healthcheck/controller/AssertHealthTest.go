package controller

import (
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/util"
	"github.com/obonobo/express-vpn-updater/server/app/util/testutils"
	"github.com/stretchr/testify/assert"
)

// Generic assertion for any request handlers that are supposed to act like the
// /healthcheck route.
func AssertHealthTest(t *testing.T, requestConsumer func(req util.Request) (util.Response, error)) {
	const (
		errorInResponse    = "controller.Healthcheck may not return an error"
		badResponseBody    = "The body of the healthcheck must be present and must be a valid JSON formatted string"
		badResponseHeaders = "got bad headers in the health check response"
		receivedError      = "received error: %v"
	)

	var (
		desiredHeaders = util.DefaultHeaders()
		desiredBody    = map[string]interface{}{"message": HealthcheckMessage}
	)

	response, err := requestConsumer(util.Request{})
	if err != nil {
		assert.Failf(t, errorInResponse, receivedError, err)
	}
	testutils.AssertHeaders(t, desiredHeaders, response.Headers)
	testutils.AssertBody(t, desiredBody, response.Body)
}
