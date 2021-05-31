package cache

import (
	"errors"
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/util/testutils/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	mockResponseFromStore   = "https://store.com/latest"
	mockResponseFromScraper = "https://scraper.com/latest"
)

var (
	errMockStore   = errors.New("bad store")
	errMockScraper = errors.New("bad scraper")
)

// Tests Cache.Get()
func TestGet(t *testing.T) {
	cache, store, scraper := createCacheWithMockedStoreAndScraper(false, false)
	resp, err := cache.Get()

	if err != nil {
		assert.Failf(t, "MockStore should not return an error", "%v", err)
	}
	assert.Equal(
		t, mockResponseFromStore, resp,
		"The Cache should respond with the mocked download link")
	assert.Equal(
		t, 1, store.NumberOfGetCalls,
		"Store.Get() should have been invoked once")
	assert.Equal(
		t, 1, scraper.NumberOfScrapeCalls,
		"Scraper should also have been invoked once")
}

// Tests Cache.Refresh()
func TestRefreshA(t *testing.T) {

}

// Tests Cache.RefreshFrom()
func TestRefreshFrom(t *testing.T) {

}

func createCacheWithMockedStoreAndScraper(
	withStoreError, withScraperError bool,
) (Cache, *mocks.MockStore, *mocks.MockScraper) {
	var storeError, scraperError error
	if withStoreError {
		storeError = errMockStore
	}
	if withScraperError {
		scraperError = errMockScraper
	}
	store := &mocks.MockStore{Output: mockResponseFromStore, Err: storeError}
	scraper := &mocks.MockScraper{Output: mockResponseFromScraper, Err: scraperError}
	return New(store, scraper), store, scraper
}
