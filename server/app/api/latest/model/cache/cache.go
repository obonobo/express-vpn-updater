package cache

type Cache interface {
	Get() (string, error)
	Refresh() (string, error)
	RefreshFrom(url string) error
}
