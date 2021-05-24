package service

import "github.com/obonobo/express-vpn-updater/server/app/model/cache"

type scraper struct {
	cache cache.Cache
}

func New(cash cache.Cache) Service {
	if cash == nil {
		return &scraper{cache: cache.NewDefaultCache()}
	} else {
		return &scraper{cache: cash}
	}
}

func Default() Service {
	return New(nil)
}

func (s *scraper) Latest() (string, error) {
	return s.cache.Get()
}

func (s *scraper) UpdateCache() (string, error) {
	return s.cache.Refresh()
}
