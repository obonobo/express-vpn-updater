package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCache struct {
	RefreshFromInputs []string
	GetOutputs        []string
	RefreshOutputs    []string

	get         func() (string, error)
	refresh     func() (string, error)
	refreshFrom func(string) error
}

func NewMockCache() *MockCache {
	return &MockCache{
		get:         func() (string, error) { return "", nil },
		refresh:     func() (string, error) { return "", nil },
		refreshFrom: func(s string) error { return nil },
	}
}

func (mc *MockCache) Get() (string, error) {
	got, err := mc.get()
	mc.GetOutputs = append(mc.GetOutputs, got)
	return got, err
}

func (mc *MockCache) Refresh() (string, error) {
	refreshed, err := mc.refresh()
	mc.RefreshOutputs = append(mc.RefreshOutputs, refreshed)
	return refreshed, err
}

func (mc *MockCache) RefreshFrom(url string) error {
	mc.RefreshFromInputs = append(mc.RefreshFromInputs, url)
	return mc.refreshFrom(url)
}

func (mc *MockCache) WithGet(get func() (string, error)) *MockCache {
	mc.get = get
	return mc
}

func (mc *MockCache) WithRefresh(refresh func() (string, error)) *MockCache {
	mc.refresh = refresh
	return mc
}

func (mc *MockCache) WithRefreshFrom(refreshFrom func(string) error) *MockCache {
	mc.refreshFrom = refreshFrom
	return mc
}

func (mc *MockCache) AssertGetWasCalled(t *testing.T, msgAndArgs ...interface{}) *MockCache {
	return mc.AssertGetWasCalledMultipleTimes(t, once, msgAndArgs...)
}

func (mc *MockCache) AssertGetWasCalledMultipleTimes(t *testing.T, numberOfTimes int, msgAndArgs ...interface{}) *MockCache {
	assert.Len(t, mc.GetOutputs, numberOfTimes, msgAndArgs...)
	return mc
}

func (mc *MockCache) AssertRefreshWasCalled(t *testing.T, msgAndArgs ...interface{}) *MockCache {
	return mc.AssertRefreshWasCalledMultipleTimes(t, once, msgAndArgs...)
}

func (mc *MockCache) AssertRefreshWasCalledMultipleTimes(t *testing.T, numberOfTimes int, msgAndArgs ...interface{}) *MockCache {
	assert.Len(t, mc.RefreshOutputs, numberOfTimes, msgAndArgs...)
	return mc
}

// Runs an assertion on the latest recorded output from the Cache.Get() method.
// If there is no recorded Cache.Get() calls, then the assertion fails
func (mc *MockCache) AssertGetOutput(t *testing.T, assertion func(string)) *MockCache {
	mc.AssertGetWasCalled(t)
	assertion(mc.GetOutputs[len(mc.GetOutputs)-1])
	return mc
}

// Runs an assertion on the latests recorded output from the Cache.Refresh()
// method. If there is no recorded Cache.Refresh() calls, then the assertion
// fails.
func (mc *MockCache) AssertRefreshOutput(t *testing.T, assertion func(string)) *MockCache {
	mc.AssertRefreshWasCalled(t)
	assertion(mc.RefreshOutputs[len(mc.RefreshOutputs)-1])
	return mc
}
