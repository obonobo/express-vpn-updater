package service

import "github.com/obonobo/express-vpn-updater/scrape/model"

const defaultUrl = "https://www.expressvpn.com/latest-1#linux"

type ScraperService interface {
	Link() (string, error)
}

type scraper struct {
	rootUrl      string
	downloadLink string
}

func New(url string) ScraperService {
	return &scraper{
		downloadLink: "",
		rootUrl:      url,
	}
}

func Default() ScraperService {
	return New(defaultUrl)
}

func (s *scraper) Link() (string, error) {
	err := s.initLink()
	if err != nil {
		return s.getFromCacheInstead()
	}
	s.uploadToCache(s.downloadLink)
	return s.downloadLink, nil
}

func (s *scraper) getFromCacheInstead() (string, error) {
	return "", nil
}

func (s *scraper) uploadToCache(url string) error {
	return nil
}

func (s *scraper) initLink() error {
	if s.downloadLink == "" {
		link, err := scrape(s.rootUrl)
		if err != nil {
			return err
		}
		s.downloadLink = link
	}
	return nil
}

func scrape(url string) (string, error) {
	page, err := model.GetDownloadPage(url)
	if err != nil {
		return "", err
	}
	link, err := model.ExtractLinuxDownloadLinkFromPage(page)
	if err != nil {
		return "", err
	}
	return link, nil
}
