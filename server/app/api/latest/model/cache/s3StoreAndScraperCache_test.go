package cache

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/obonobo/express-vpn-updater/server/app/util/testutils/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	mockResponseFromStore       = "https://store.com/latest"
	mockResponseFromScraper     = "https://scraper.com/latest"
	refreshFromInputUrl         = mockResponseFromScraper
	methodShouldHaveBeenInvoked = "%s should have been invoked at least once"
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
	assert.GreaterOrEqual(
		t, store.NumberOfGetCalls, 1,
		fmt.Sprintf(methodShouldHaveBeenInvoked, "Store.Get()"))
	time.Sleep(time.Millisecond)
	assert.GreaterOrEqual(
		t, scraper.NumberOfScrapeCalls, 1,
		"Scraper should also have been invoked once within 1 ms of Store.Get()")
}

// Tests Cache.Refresh()
func TestRefresh(t *testing.T) {
	cache, store, scraper := createCacheWithMockedStoreAndScraper(false, false)
	resp, err := cache.Refresh()

	if err != nil {
		assert.Failf(t, "The mocks should not return errors here", "%v", err)
	}
	assert.Equal(
		t, mockResponseFromScraper, resp,
		"Scraper.Scrape() should have been called, and the mock response should have been returned")

	assert.GreaterOrEqual(
		t, scraper.NumberOfScrapeCalls, 1,
		fmt.Sprintf(methodShouldHaveBeenInvoked, "Scraper.Scrape()"))

	time.Sleep(time.Millisecond)
	assert.GreaterOrEqual(
		t, store.NumberOfPutCalls, 1,
		fmt.Sprintf(
			methodShouldHaveBeenInvoked+" withing 1 ms of Scraper.Scrape()",
			"Store.Get()"))
}

// Tests Cache.RefreshFrom()
func TestRefreshFrom(t *testing.T) {
	cache, store, _ := createCacheWithMockedStoreAndScraper(false, false)
	if err := cache.RefreshFrom(refreshFromInputUrl); err != nil {
		assert.Failf(t, "Mocked Store.Put() should not return an error", "%v", err)
	}
	assert.GreaterOrEqual(
		t, store.NumberOfPutCalls, 1,
		fmt.Sprintf(methodShouldHaveBeenInvoked, "Store.Put()"))
	assert.Equal(
		t, store.PutInputs[0], refreshFromInputUrl,
		"The url passed to Store.Put() should be the same as was passed to Cache.RefreshFrom()")
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
