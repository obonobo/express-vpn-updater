package mocks

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Get func(url string) (*http.Response, error)
type Do func(*http.Request) (*http.Response, error)

type MockHttpClient struct {
	GetInputs []string
	DoInputs  []*http.Request

	get Get
	do  Do
}

func NewMockHttpClient() *MockHttpClient {
	return &MockHttpClient{
		GetInputs: []string{},
		get: func(url string) (*http.Response, error) {
			return &http.Response{}, nil
		},
	}
}

func (mhc *MockHttpClient) Get(url string) (*http.Response, error) {
	mhc.GetInputs = append(mhc.GetInputs, url)
	return mhc.get(url)
}

func (mhc *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	mhc.DoInputs = append(mhc.DoInputs, req)
	return mhc.do(req)
}

func (mhc *MockHttpClient) WithGet(get Get) *MockHttpClient {
	mhc.get = get
	return mhc
}

func (mhc *MockHttpClient) WithDo(do Do) *MockHttpClient {
	mhc.do = do
	return mhc
}

func (mhc *MockHttpClient) AssertGetWasCalled(t *testing.T, msgAndArgs ...interface{}) *MockHttpClient {
	return mhc.AssertGetWasCalledMultipleTimes(t, once, msgAndArgs...)
}

func (mhc *MockHttpClient) AssertGetWasCalledMultipleTimes(
	t *testing.T,
	numberOfTimes int,
	msgAndArgs ...interface{},
) *MockHttpClient {
	assert.Len(t, mhc.GetInputs, numberOfTimes, msgAndArgs...)
	return mhc
}

func (mhc *MockHttpClient) AssertDoWasCalled(t *testing.T, msgAndArgs ...interface{}) *MockHttpClient {
	return mhc.AssertDoWasCalledMultipleTimes(t, once, msgAndArgs...)
}

func (mhc *MockHttpClient) AssertDoWasCalledMultipleTimes(
	t *testing.T,
	numberOfTimes int,
	msgAndArgs ...interface{},
) *MockHttpClient {
	assert.Len(t, mhc.DoInputs, numberOfTimes, msgAndArgs...)
	return mhc
}

// Assert that the given url was pinged (via a GET request) a specific number of
// times
func (mhc *MockHttpClient) AssertUrlPinged(
	t *testing.T, url string,
	numberOfTimes int,
	msgAndArgs ...interface{},
) (this *MockHttpClient) {
	urlWasPinged := 0
	for _, v := range mhc.GetInputs {
		if v == url {
			urlWasPinged++
		}
	}
	assert.Equal(t, urlWasPinged, numberOfTimes, msgAndArgs...)
	return mhc
}

// Run an assertion on the last recorded input that was passed to the
// HttpClient.Do() function of the MockHttpClient. If the HttpClient.Do()
// function was never called and there are no recorded inputs, then the
// assertion fails immediately.
func (mhc *MockHttpClient) AssertDoInput(t *testing.T, assertion func(*http.Request)) *MockHttpClient {
	mhc.AssertDoWasCalled(t, "HttpClient.Do() must be called at least once")
	assertion(mhc.DoInputs[len(mhc.DoInputs)-1])
	return mhc
}

// Run an assertion on the last recorded input that was passed to the
// HttpClient.Get() function of the MockHttpClient. If the HttpClient.Get()
// function was never called and there are no recorded inputs, then the
// assertion fails immediately.
func (mhc *MockHttpClient) AssertGetInput(
	t *testing.T,
	assertion func(string),
) *MockHttpClient {
	mhc.AssertGetWasCalled(t, "HttpClient.Get() must be called at least once")
	assertion(mhc.GetInputs[len(mhc.GetInputs)-1])
	return mhc
}
