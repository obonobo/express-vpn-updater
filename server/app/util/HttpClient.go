package util

import "net/http"

type HttpClient interface {

	// Make a GET request to the given resource
	Get(url string) (resp *http.Response, err error)

	// Execute an HTTP request
	Do(*http.Request) (resp *http.Response, err error)
}
