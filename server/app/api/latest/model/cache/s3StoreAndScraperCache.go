package cache

import (
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/model/scraper"
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/model/store"
	"github.com/obonobo/express-vpn-updater/server/app/config"
)

type s3StoreAndScraperCache struct {
	store.Store
	scraper.Scraper
}

func New(store store.Store, scraper scraper.Scraper) Cache {
	return &s3StoreAndScraperCache{store, scraper}
}

func NewCache(bucket string, scrapingFromUrl string) Cache {
	return New(
		store.NewStoreWithBucket(bucket),
		scraper.NewScraper(scrapingFromUrl),
	)
}

func NewDefaultCache() Cache {
	return NewCache(config.Get().Bucket, config.Get().Url)
}

func (c *s3StoreAndScraperCache) Get() (string, error) {
	if cached, err := c.Store.Get(); err == nil {
		go c.Refresh()
		return cached, err
	}
	return c.Refresh()
}

func (c *s3StoreAndScraperCache) Refresh() (string, error) {
	if url, err := c.Scraper.Scrape(); err == nil {
		go c.Store.Put(url)
		return url, err
	} else {
		return "", err
	}
}

func (c *s3StoreAndScraperCache) RefreshFrom(url string) error {
	return c.Store.Put(url)
}
