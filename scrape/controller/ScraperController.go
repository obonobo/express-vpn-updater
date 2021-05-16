package controller

import (
	"strconv"

	"github.com/obonobo/express-vpn-updater/scrape/service"
	"github.com/obonobo/express-vpn-updater/util"
)

const (
	REDIRECT_QUERY_PARAM_KEY = "redirect"
)

type QueryParams struct {
	redirect bool
}

type paramsCache struct {
	params *QueryParams
}

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
	if parse(req).redirect {
		return c.RedirectToLatest()
	}
	return c.GrabLatestLink()
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

func (p *paramsCache) save(req util.Request) *paramsCache {
	if p.params == nil {
		p.params = &QueryParams{redirect: true}
	}
	requestParams := req.QueryStringParameters
	wantsRedirect, ok := requestParams[REDIRECT_QUERY_PARAM_KEY]
	redirectNeeded, err := strconv.ParseBool(wantsRedirect)
	if ok && err == nil {
		p.params.redirect = redirectNeeded
	}
	return p
}

func parse(req util.Request) QueryParams {
	return *(&paramsCache{}).save(req).params
}
