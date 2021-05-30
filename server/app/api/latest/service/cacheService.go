package service

import "github.com/obonobo/express-vpn-updater/server/app/api/latest/model/cache"

type cacheService struct {
	cache cache.Cache
}

func New(cash cache.Cache) Service {
	if cash == nil {
		return &cacheService{cache: cache.NewDefaultCache()}
	} else {
		return &cacheService{cache: cash}
	}
}

func Default() Service {
	return New(nil)
}

func (s *cacheService) Latest() (string, error) {
	return s.cache.Get()
}

func (s *cacheService) UpdateCache() (string, error) {
	return s.cache.Refresh()
}
