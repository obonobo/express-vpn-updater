package cache

// A Cache provides access to the express vpn .deb package via a direct download
// link.
//
// The Cache implementation in this package uses an S3Store to store
// downloaded packages and a Scraper instance for refreshing the Cache
type Cache interface {

	// Retrieve the value stored in the Cache
	Get() (string, error)

	// Manually invoke a re-scraping of the expressvpn website and an updating
	// of the s3 cache
	Refresh() (string, error)

	// Manually invoke a re-scraping of the provided website and an updating of
	// the s3 cache
	RefreshFrom(url string) error
}
