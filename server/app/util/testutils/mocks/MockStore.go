package mocks

type MockStore struct {
	Output           string
	Err              error
	PutInputs        []string
	NumberOfGetCalls int
	NumberOfPutCalls int
}

func (ms *MockStore) Get() (string, error) {
	ms.NumberOfGetCalls++
	return ms.Output, ms.Err
}

func (ms *MockStore) Put(downloadFromUrl string) error {
	ms.NumberOfPutCalls++
	return ms.Err
}
