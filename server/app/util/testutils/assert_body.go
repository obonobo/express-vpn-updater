package testutils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	jsonUnmarshalError = "failed to unmarshal response body"
)

func AssertBody(t *testing.T, expected map[string]interface{}, actual string, msgAndArgs ...interface{}) {
	body, err := parseBody(actual)
	if err != nil {
		assert.Fail(t, jsonUnmarshalError, msgAndArgs...)
	}
	assert.Equal(t, expected, body, msgAndArgs...)
}

func parseBody(body string) (map[string]interface{}, error) {
	var bod map[string]interface{}
	err := json.Unmarshal([]byte(body), &bod)
	return bod, err
}
