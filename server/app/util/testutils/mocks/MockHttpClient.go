package mocks

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Get func(url string) (*http.Response, error)

type MockHttpClient struct {
	Urls []string
	get  Get
}

func NewMockHttpClient() *MockHttpClient {
	return &MockHttpClient{
		Urls: []string{},
		get: func(url string) (*http.Response, error) {
			return &http.Response{}, nil
		},
	}
}

func (mhc *MockHttpClient) WithGet(get Get) *MockHttpClient {
	mhc.get = get
	return mhc
}

func (mhc *MockHttpClient) Get(url string) (*http.Response, error) {
	mhc.Urls = append(mhc.Urls, url)
	return mhc.get(url)
}

// Assert that the given url was pinged (via a GET request) a specific number of
// times
func (mhc *MockHttpClient) AssertUrlPinged(
	t *testing.T, url string,
	numberOfTimes int,
	msgAndArgs ...interface{},
) (this *MockHttpClient) {
	urlWasPinged := 0
	for _, v := range mhc.Urls {
		if v == url {
			urlWasPinged++
		}
	}
	assert.Equal(t, urlWasPinged, numberOfTimes, msgAndArgs...)
	return mhc
}
