package controller

import (
	"strconv"

	"github.com/obonobo/express-vpn-updater/scrape/service"
	"github.com/obonobo/express-vpn-updater/util"
)

const (
	RedirectQueryParamKey = "redirect"
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

func (c *ScraperController) ScrapeAndRespond(req util.Request) util.Response {
	params := req.QueryStringParameters
	wantsRedirect, ok := params[RedirectQueryParamKey]
	asBool, err := strconv.ParseBool(wantsRedirect)
	if ok && err == nil && !asBool {
		return c.GrabLatestLink()
	}
	return c.RedirectToLatest()
}

func (c *ScraperController) GrabLatestLink() util.Response {
	link, err := c.s.Link()
	if clientError, ok := util.Panic(err); !ok {
		return *clientError
	}
	return util.BasicMessage(link)
}

func (c *ScraperController) RedirectToLatest() util.Response {
	link, err := c.s.Link()
	if clientError, ok := util.Panic(err); !ok {
		return *clientError
	}
	return util.Redirect(link)
}
