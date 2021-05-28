package service

// Accesses the model
type Service interface {

	// Grab a download link for the latest express vpn package
	Latest() (downloadLink string, err error)

	// Scraps the download link, saves in cache, and sends to you
	UpdateCache() (downloadLink string, err error)
}
