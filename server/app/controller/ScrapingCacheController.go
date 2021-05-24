package controller

import (
	"github.com/obonobo/express-vpn-updater/server/app/service"
	"github.com/obonobo/express-vpn-updater/server/app/util"
	"github.com/obonobo/express-vpn-updater/server/config"
)

var logger = config.Get().Logger()

// var logger = log.Default()

type ScrapingCacheController struct {
	service service.Service
}

func New(servs service.Service) Controller {
	if servs == nil {
		return &ScrapingCacheController{service: service.Default()}
	} else {
		return &ScrapingCacheController{service: servs}
	}
}

func Default() Controller {
	return New(nil)
}

func (c *ScrapingCacheController) Latest(req util.Request) util.Response {
	logger.Println("Inside Controller.Latest...")
	params := ParseParams(req)
	logger.Println(params)
	if params.fresh {
		return c.UpdateCache(req, params.redirect)
	}
	return c.CachedResponse(req, params)
}

func (c *ScrapingCacheController) CachedResponse(req util.Request, params *QueryParams) util.Response {
	logger.Println("Inside Controller.CachedResponse")
	if params.redirect {
		return c.latestLinkResponse(util.Redirect)
	} else {
		return c.latestLinkResponse(util.BasicMessage)
	}
}

func (c *ScrapingCacheController) UpdateCache(req util.Request, redirect bool) util.Response {
	logger.Println("Inside Controller.UpdateCache")
	if res, err := c.service.UpdateCache(); err == nil {
		if redirect {
			return util.Redirect(res)
		}
		return util.BasicMessage(res)
	} else {
		return c.serviceUnavailable(err)
	}
}

func (c *ScrapingCacheController) latestLinkResponse(resp func(link string) util.Response) util.Response {
	logger.Println("Inside Controller.latestLinkResponse")
	link, err := c.service.Latest()
	if clientError, ok := util.Panic(err); !ok {
		return *clientError
	}
	return resp(link)
}

func (c *ScrapingCacheController) serviceUnavailable(err error) util.Response {
	return util.ServiceUnavailable(map[string]interface{}{
		"message": err,
	}, nil)
}
