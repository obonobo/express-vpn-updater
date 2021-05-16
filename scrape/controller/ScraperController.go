package controller

import (
	"github.com/obonobo/express-vpn-updater/scrape/service"
	"github.com/obonobo/express-vpn-updater/util"
)

type ScraperController struct {
	s service.ScraperService
}

func New(s service.ScraperService) *ScraperController {
	myService := s
	if s == nil {
		myService = service.Default()
	}
	return &ScraperController{
		s: myService,
	}
}

func Default() *ScraperController {
	return New(nil)
}

func (c *ScraperController) DownloadLatest() util.Response {
	link, err := c.s.Link()
	if clientError, ok := util.Panic(err); !ok {
		return *clientError
	}
	return util.BasicMessage(link)
}
