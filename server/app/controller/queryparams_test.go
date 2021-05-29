package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamsCache(t *testing.T) {
	cache := NewParamsCache()
	assert.NotNil(t, cache, "TODO: write a meaningful unit test")
}

func TestParamsParsing(t *testing.T) {
	assert.True(t, true, "TODO: write a unit test")
}
