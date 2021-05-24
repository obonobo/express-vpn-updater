package service

type Service interface {
	Latest() (string, error)
	UpdateCache() (url string, err error)
}
