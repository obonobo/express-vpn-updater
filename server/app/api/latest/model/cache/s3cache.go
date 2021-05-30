package cache

import (
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/model/scraper"
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/model/store"
	"github.com/obonobo/express-vpn-updater/server/app/config"
)

type s3cache struct {
	store.Store
	*scraper.Scraper
}

func NewCache(bucket string, scrapingFromUrl string) Cache {
	return &s3cache{
		Store:   store.NewStoreWithBucket(bucket),
		Scraper: scraper.NewScraper(scrapingFromUrl),
	}
}

func NewDefaultCache() Cache {
	return NewCache(config.Get().Bucket, config.Get().Url)
}

func (c *s3cache) Get() (string, error) {
	if cached, err := c.Store.Get(); err == nil {
		go c.Refresh()
		return cached, err
	}
	return c.Refresh()
}

func (c *s3cache) Refresh() (string, error) {
	if url, err := c.Scraper.Scrape(); err == nil {
		go c.Store.Put(url)
		return url, err
	} else {
		return "", err
	}
}

func (c *s3cache) RefreshFrom(url string) error {
	return c.Store.Put(url)
}
