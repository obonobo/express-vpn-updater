package service

import (
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/util/testutils/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	unexpectedErrorMessage = "service should not throw an error"
)

var (
	cacheGetOutput     = "Get"
	cacheRefreshOutput = "Refresh"
)

// Tests Service.Latest()
func TestLatest(t *testing.T) {
	service, mockCache := createServiceWithMock()
	resp, err := service.Latest()
	if err != nil {
		assert.Failf(t, unexpectedErrorMessage, "%v", err)
	}
	assert.Equal(t, cacheGetOutput, resp)
	mockCache.AssertGetOutput(t, func(output string) {
		assert.Equal(t, cacheGetOutput, output)
	})
}

// Tests Service.UpdateCache()
func TestUpdateCache(t *testing.T) {
	service, mockCache := createServiceWithMock()
	resp, err := service.UpdateCache()
	if err != nil {
		assert.Failf(t, unexpectedErrorMessage, "%v", err)
	}
	assert.Equal(t, cacheRefreshOutput, resp)
	mockCache.AssertRefreshOutput(t, func(output string) {
		assert.Equal(t, cacheRefreshOutput, output)
	})
}

func createServiceWithMock() (Service, *mocks.MockCache) {
	cache := createMockCache()
	return New(cache), cache
}

func createMockCache() *mocks.MockCache {
	return mocks.
		NewMockCache().
		WithGet(func() (string, error) { return cacheGetOutput, nil }).
		WithRefresh(func() (string, error) { return cacheRefreshOutput, nil }).
		WithRefreshFrom(func(s string) error { return nil })
}
