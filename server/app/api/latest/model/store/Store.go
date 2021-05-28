package store

type Store interface {

	// Grab a download link for the latest express vpn .deb package stored in
	// the AWS S3 cache
	Get() (url string, err error)

	// Provide the direct download link for the express vpn package. This step
	// should happen post-scraping; the Store will not scrape the given URL - it
	// needs to be a direct download link with a file at the other end
	Put(downloadFromUrl string) (err error)
}
