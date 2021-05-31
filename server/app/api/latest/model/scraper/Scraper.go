package scraper

// A Scraper scrapes a data source and finds the direct download link for the
// express vpn .deb package
type Scraper interface {

	// Scrapes the data source and returns the download link for the .deb package
	Scrape() (directDownloadLink string, err error)
}
