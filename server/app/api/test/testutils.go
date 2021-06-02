package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	latest "github.com/obonobo/express-vpn-updater/server/app/api/latest/controller"
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/model/cache"
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/model/scraper"
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/model/store"
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/service"
	"github.com/obonobo/express-vpn-updater/server/app/util"
	"github.com/obonobo/express-vpn-updater/server/app/util/testutils"
	"github.com/obonobo/express-vpn-updater/server/app/util/testutils/mocks"
)

const (
	scrapingUrl         = "http://scraping-url.com/scrape/me"
	bucket              = "handlers-test-bucket"
	controllerDotLatest = "Controller.Latest()"
	receivedError       = "received error: %v"
	once                = testutils.Once
)

var (
	mockListObjectsV2Output       = mocks.MockListObjectsV2Output()
	expectedLatestResponseHeaders = map[string]string{
		"Content-Type": "application/json",
	}
)

func AssertS3ClientDidTheRightStuff(t *testing.T, s3Client *mocks.MockS3Client) {
	s3Client.
		AssertListObjectsV2Input(t, func(input *s3.ListObjectsV2Input) {

		}).
		AssertBucketListed(t, "", "Bucket should have been listed here")
}

func AssertHttpClientDidTheRightStuff(t *testing.T, httpClient *mocks.MockHttpClient) {
	httpClient.
		AssertDoInput(t, func(r *http.Request) {

		}).
		AssertGetInput(t, func(s string) {

		}).
		AssertUrlPinged(t, "", 1, "The mock url should have been pinged here")
}

func AssertHeadersMatchExpected(t *testing.T, headers map[string]string) {
	testutils.AssertHeaders(
		t, expectedLatestResponseHeaders, headers,
		fmt.Sprintf(
			"Response headers for %s must match the expected value",
			controllerDotLatest))

}

func AssertBodyMatchesExpected(t *testing.T, body string) {
	testutils.AssertBody(
		t, map[string]interface{}{}, body,
		fmt.Sprintf(
			"Response body for %s must match the expected value ",
			controllerDotLatest))
}

// Creates a test http.Request to pass in to the Controller.Latest() endpoint.
func CreateMockRequest() util.Request {
	return util.Request{
		Body: "",
	}
}

// Initializes all the objects that are needed to respond to API requests on the
// /latest route. The final 2 objects in the chain: the S3Client, and the HTTP
// client are mocked out and references to the mock objects are returned.
func CreateMockedUpLatestController() (
	latest.Controller,
	*mocks.MockHttpClient,
	*mocks.MockS3Client,
) {
	var (
		httpClient       = mocks.NewMockHttpClient()
		s3Client         = mocks.NewMockS3Client()
		scraper          = scraper.New(httpClient, scrapingUrl)
		store            = store.New(s3Client, httpClient, bucket)
		cache            = cache.New(store, scraper)
		service          = service.New(cache)
		latestController = latest.New(service)
	)

	s3Client.
		WithListObjectsV2(func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
			return mockListObjectsV2Output, nil
		}).
		WithPutObject(func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			return nil, nil
		})

	httpClient.
		WithDo(func(r *http.Request) (*http.Response, error) {
			return nil, nil
		}).
		WithGet(func(url string) (*http.Response, error) {
			return nil, nil
		})

	return latestController, httpClient, s3Client
}
