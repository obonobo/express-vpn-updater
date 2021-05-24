package scraper

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/obonobo/express-vpn-updater/server/app/util"
)

const (
	userAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0"
)

type Scraper struct {
	sourceUrl    string
	scrappedLink string
}

func NewScraper(sourceUrl string) *Scraper {
	return &Scraper{sourceUrl: sourceUrl}
}

func (s *Scraper) Scrape() (string, error) {
	err := s.initLink()
	return s.scrappedLink, err
}

// Gets the source url of the scraper
func (s *Scraper) Url() string {
	return s.sourceUrl
}

func (s *Scraper) initLink() error {
	if s.scrappedLink != "" {
		return nil
	}
	page, err := s.getDownloadPage()
	if err != nil {
		return err
	}
	s.scrappedLink, err = s.extractDownloadLinkFromPage(page)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scraper) getDownloadPage() (*goquery.Document, error) {
	req, err := http.NewRequest("GET", s.Url(), nil)
	if err != nil {
		return &goquery.Document{}, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &goquery.Document{}, err
	}
	if util.ResponseIsBad(resp) {
		return &goquery.Document{}, NewBadWesiteError(resp)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return &goquery.Document{}, err
	}
	return doc, nil
}

func (s *Scraper) extractDownloadLinkFromPage(doc *goquery.Document) (string, error) {
	hrefPattern, err := regexp.Compile(`.*linux.*amd64.*\.deb`)
	if err != nil {
		return "", err
	}

	found := doc.
		Find("a:contains(\"Download\")").
		FilterFunction(func(v int, sel *goquery.Selection) bool {
			href, exists := sel.Attr("href")
			return exists && hrefPattern.Match([]byte(href))
		})

	href, exists := found.Attr("href")
	if !exists {
		return "", errors.New("could not find the Linux download link in this document")
	}

	return href, nil
}
