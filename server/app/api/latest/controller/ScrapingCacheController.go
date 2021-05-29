package controller

import (
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/controller/queryparams"
	"github.com/obonobo/express-vpn-updater/server/app/api/latest/service"
	"github.com/obonobo/express-vpn-updater/server/app/config"
	"github.com/obonobo/express-vpn-updater/server/app/util"
)

var logger = config.Get().Logger()

// A Controller that allows you to access the scraping and caching functionality
type ScrapingCacheController struct {
	service service.Service
}

// Create a new ScrapingCacheController connected to the given service
func New(servs service.Service) Controller {
	if servs == nil {
		return &ScrapingCacheController{service: service.Default()}
	} else {
		return &ScrapingCacheController{service: servs}
	}
}

// Create a new ScrapingCacheController that grabs a default Service instance
func Default() Controller {
	return New(nil)
}

//
func (c *ScrapingCacheController) Latest(req util.Request) util.Response {
	logger.Println("Inside Controller.Latest...")
	params := queryparams.ParseParams(req)
	logger.Println(params)
	if params.Fresh {
		return c.UpdateCache(req, params.Redirect)
	}
	return c.CachedResponse(req, params)
}

func (c *ScrapingCacheController) CachedResponse(
	req util.Request,
	params *queryparams.QueryParams,
) util.Response {
	logger.Println("Inside Controller.CachedResponse")
	if params.Redirect {
		return c.latestLinkResponse(util.Redirect)
	} else {
		return c.latestLinkResponse(util.BasicMessage)
	}
}

func (c *ScrapingCacheController) UpdateCache(
	req util.Request,
	redirect bool,
) util.Response {
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

func (c *ScrapingCacheController) latestLinkResponse(
	resp func(link string) util.Response,
) util.Response {
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
