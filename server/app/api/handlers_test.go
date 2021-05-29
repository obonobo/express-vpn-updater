package api

import (
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/api/healthcheck/controller"
	"github.com/stretchr/testify/assert"
)

const (
	receivedError = "received error: %v"
)

func TestHealthcheck(t *testing.T) {
	controller.AssertHealthTest(t, Healthcheck)
}

func TestLatest(t *testing.T) {
	assert.True(t, true, "TODO: implement handler test")
}
