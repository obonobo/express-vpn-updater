package mocks

type MockScraper struct {
	Output              string
	Err                 error
	NumberOfScrapeCalls int
}

func (ms *MockScraper) Scrape() (string, error) {
	ms.NumberOfScrapeCalls++
	return ms.Output, ms.Err
}
