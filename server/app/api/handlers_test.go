package api

import (
	"testing"

	healthcheck "github.com/obonobo/express-vpn-updater/server/app/api/healthcheck/controller"
	"github.com/obonobo/express-vpn-updater/server/app/api/test"
)

// Tests: GET /healthcheck
func TestHealthcheck(t *testing.T) {
	healthcheck.AssertHealthTest(t, Healthcheck)
}

// Tests: GET /latest
func TestLatest(t *testing.T) {

	// !!! REMOVE
	t.Skip()
	// !!! REMOVE

	latestController, httpClient, s3Client := test.CreateMockedUpLatestController()
	req := test.CreateMockRequest()
	resp := latestController.Latest(req)

	test.AssertHeadersMatchExpected(t, resp.Headers)
	test.AssertBodyMatchesExpected(t, resp.Body)
	test.AssertHttpClientDidTheRightStuff(t, httpClient)
	test.AssertS3ClientDidTheRightStuff(t, s3Client)
}
