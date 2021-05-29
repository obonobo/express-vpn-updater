package queryparams

import (
	"strconv"
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/util"
	"github.com/stretchr/testify/assert"
)

var (
	inputs = [][]bool{
		{true, true},
		{true, false},
		{false, false},
		{false, true},
	}
)

func TestParseParams(t *testing.T) {
	runInputs(inputs, func(redirect, fresh bool) {
		request := createRequest(redirect, fresh)
		params := createQueryParams(redirect, fresh)
		parsed := ParseParams(request)
		assert.Equal(t, params, parsed, "parsed result must equal expected")
	})
}

func TestSave(t *testing.T) {
	runInputs(inputs, func(redirect, fresh bool) {
		cache := NewParamsCache()
		request := createRequest(redirect, fresh)
		params := createQueryParams(redirect, fresh)
		parsed := cache.Save(request).GetParams()
		assert.Equal(t, params, parsed)
	})
}

func TestParamsCache(t *testing.T) {
	cache := NewParamsCache()
	assert.NotNil(t, cache, "TODO: write a meaningful unit test")
}

func createRequest(redirect, fresh bool) util.Request {
	return util.Request{
		QueryStringParameters: map[string]string{
			"redirect": strconv.FormatBool(redirect),
			"fresh":    strconv.FormatBool(fresh),
		},
	}
}

func createQueryParams(redirect, fresh bool) *QueryParams {
	return &QueryParams{
		redirect: redirect,
		fresh:    fresh,
	}
}

func runInputs(inputs [][]bool, assertion func(bool, bool)) {
	for _, pair := range inputs {
		if len(pair) != 2 {
			continue
		}
		assertion(pair[0], pair[1])
	}
}
