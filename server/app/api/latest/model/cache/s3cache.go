package cache

import (
	"github.com/obonobo/express-vpn-updater/server/app/model/scraper"
	"github.com/obonobo/express-vpn-updater/server/app/model/store"
	"github.com/obonobo/express-vpn-updater/server/config"
)

type s3cache struct {
	store   *store.Store
	scraper *scraper.Scraper
}

func NewCache(bucket string, scrapingFromUrl string) Cache {
	return &s3cache{
		store:   store.NewStore(bucket),
		scraper: scraper.NewScraper(scrapingFromUrl),
	}
}

func NewDefaultCache() Cache {
	return NewCache(config.Get().Bucket, config.Get().Url)
}

func (c *s3cache) Get() (string, error) {
	if cached, err := c.store.Get(); err == nil {
		go c.Refresh()
		return cached, err
	}
	return c.Refresh()
}

func (c *s3cache) Refresh() (string, error) {
	if url, err := c.scraper.Scrape(); err == nil {
		go c.store.Put(url)
		return url, err
	} else {
		return "", err
	}
}

func (c *s3cache) RefreshFrom(url string) error {
	return c.store.Put(url)
}