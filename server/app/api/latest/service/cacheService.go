package service

import c "github.com/obonobo/express-vpn-updater/server/app/api/latest/model/cache"

type cacheService struct {
	c.Cache
}

func New(cache c.Cache) Service {
	if cache == nil {
		return &cacheService{c.NewDefaultCache()}
	} else {
		return &cacheService{cache}
	}
}

func Default() Service {
	return New(nil)
}

func (s *cacheService) Latest() (string, error) {
	return s.Cache.Get()
}

func (s *cacheService) UpdateCache() (string, error) {
	return s.Cache.Refresh()
}
