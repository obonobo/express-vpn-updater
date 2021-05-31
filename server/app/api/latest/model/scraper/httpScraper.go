package scraper

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/obonobo/express-vpn-updater/server/app/config"
	"github.com/obonobo/express-vpn-updater/server/app/util"
)

const (
	userAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0"
)

type httpScraper struct {
	client       util.HttpClient
	sourceUrl    string
	scrappedLink string
}

func New(client util.HttpClient, sourceUrl string) *httpScraper {
	return &httpScraper{sourceUrl: sourceUrl, client: client}
}

func NewScraper(sourceUrl string) *httpScraper {
	return New(http.DefaultClient, sourceUrl)
}

func Default() *httpScraper {
	return New(http.DefaultClient, config.Get().Url)
}

func (s *httpScraper) Scrape() (string, error) {
	err := s.initLink()
	return s.scrappedLink, err
}

// Gets the source url of the scraper
func (s *httpScraper) Url() string {
	return s.sourceUrl
}

func (s *httpScraper) initLink() error {
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

func (s *httpScraper) getDownloadPage() (*goquery.Document, error) {
	resp, err := s.requestDownloadPage()
	if err != nil {
		return &goquery.Document{}, err
	}
	return s.extractGoQueryDocumentFromPage(resp)
}

func (s *httpScraper) extractGoQueryDocumentFromPage(resp *http.Response) (*goquery.Document, error) {
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return &goquery.Document{}, err
	}
	return doc, nil
}

func (s *httpScraper) requestDownloadPage() (*http.Response, error) {
	req, err := s.createRequestForDownloadPage()
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	} else if util.ResponseIsBad(resp) {
		return nil, NewBadWesiteError(resp)
	}
	return resp, nil
}

func (s *httpScraper) createRequestForDownloadPage() (*http.Request, error) {
	req, err := http.NewRequest("GET", s.Url(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	return req, nil
}

func (s *httpScraper) extractDownloadLinkFromPage(doc *goquery.Document) (string, error) {
	hrefPattern, err := regexp.Compile(`.*linux.*amd64.*\.deb`)
	if err != nil {
		return "", err
	}
	return s.scrapeDownloadLink(doc, hrefPattern)
}

func (s *httpScraper) scrapeDownloadLink(doc *goquery.Document, hrefPattern *regexp.Regexp) (string, error) {
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
