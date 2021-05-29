package controller

import (
	"strconv"

	"github.com/obonobo/express-vpn-updater/server/app/util"
)

const (
	REDIRECT_QUERY_PARAM_KEY = "redirect"
	FRESH_QUERY_PARAM_KEY    = "fresh"
)

type QueryParams struct {
	fresh    bool
	redirect bool
}

type ParamsCache struct {
	params *QueryParams
}

func NewParamsCache() *ParamsCache {
	logger.Println("Creating new ParamsCache")
	return &ParamsCache{}
}

func (p *ParamsCache) Save(req util.Request) *ParamsCache {
	logger.Println("Saving request: ", req)
	return p.
		parseRedirect(req).
		parseFresh(req)
}

func (p *ParamsCache) GetParams() *QueryParams {
	logger.Println("Inside ParamsCache.GetParams")
	if p.params == nil {
		logger.Println("I hit the if - ", p, p.params)
		p.params = &QueryParams{redirect: true}
	}
	return p.params
}

func (p *ParamsCache) setRedirect(value bool) {
	p.GetParams().redirect = value
}

func (p *ParamsCache) setFresh(value bool) {
	p.GetParams().fresh = value
}

func (p *ParamsCache) parseRedirect(req util.Request) *ParamsCache {
	return p.parseBoolQueryParam(req, REDIRECT_QUERY_PARAM_KEY, p.setRedirect)
}

func (p *ParamsCache) parseFresh(req util.Request) *ParamsCache {
	return p.parseBoolQueryParam(req, FRESH_QUERY_PARAM_KEY, p.setFresh)
}

func (p *ParamsCache) parseBoolQueryParam(req util.Request, key string, consume func(bool)) *ParamsCache {
	param, ok := req.QueryStringParameters[key]
	value, err := strconv.ParseBool(param)
	if ok && err == nil {
		consume(value)
	}
	return p
}

func ParseParams(req util.Request) *QueryParams {
	return NewParamsCache().
		Save(req).
		GetParams()
}
