package model

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

const (
	userAgent  = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0"
)

func GetDownloadPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &goquery.Document{}, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &goquery.Document{}, err
	}
	if resp.StatusCode != 200 {
		return &goquery.Document{}, errors.New("got a bad response from expressvpn website")
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return &goquery.Document{}, err
	}
	return doc, nil
}

func ExtractLinuxDownloadLinkFromPage(doc *goquery.Document) (string, error) {
	hrefPattern := regexp.MustCompile(`.*linux.*amd64.*\.deb`)

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
