package mocks

import "net/http"

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
