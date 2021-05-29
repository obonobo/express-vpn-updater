package cache

type Cache interface {

	// Retrieve the value stored in the Cache
	Get() (string, error)

	// Manually invoke a re-scraping of the expressvpn website and an updating of
	// the s3 cache
	Refresh() (string, error)

	// Manually invoke a re-scraping of the provided website and an updating of
	// the s3 cache
	RefreshFrom(url string) error
}
