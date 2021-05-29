package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertHeaders(t *testing.T, expected map[string]string, actual map[string]string, msgAndArgs ...interface{}) {
	assert.Equal(t, expected, actual, msgAndArgs...)
}
