package mocks

type stringOrErrorFunction func() (string, error)
type refreshFromFunction func(string) error

type MockCache struct {
	RefreshFromInputs []string
	GetOutputs        []string
	RefreshOutputs    []string

	get         stringOrErrorFunction
	refresh     stringOrErrorFunction
	refreshFrom refreshFromFunction
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

func (mc *MockCache) WithGet(get stringOrErrorFunction) *MockCache {
	mc.get = get
	return mc
}

func (mc *MockCache) WithRefresh(refresh stringOrErrorFunction) *MockCache {
	mc.refresh = refresh
	return mc
}

func (mc *MockCache) WithRefreshFrom(refreshFrom refreshFromFunction) *MockCache {
	mc.refreshFrom = refreshFrom
	return mc
}
