package scraper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/obonobo/express-vpn-updater/server/app/util/testutils/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	mockSourceUrl       = "http://some-fake-url.com/download"
	mockDownloadLink    = "https://www.expressvpn.works/clients/linux/expressvpn_3.8.0.4-1_amd64.deb"
	mockBadDownloadLink = "https://www.expressvpn.works/clients/windows/expressvpn_3.8.0.4-1_amd64.exe"
)

var errMockBadResponse = errors.New("bad response")
var mockDownloadPage = fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
  <body>
    <a
      class="btn app-url primary"
      data-ga-event='{"category":"Update App - Linux","action":"Linux - Ubuntu 64-bit","label":"ExpressVPN 3.8.0"}'
      href="%s"
    >
      Download
    </a>
    <a href="%s">
      Download
    </a>
    <a href="%s">
      Some other stuff
    </a>
  </body>
</html>
`, mockDownloadLink, mockBadDownloadLink, mockDownloadLink)

// Tests Scraper.Scrape()
func TestScrape(t *testing.T) {
	scraper, mockHttpClient := createScraperWithMockedHttpClient(false)
	resp, err := scraper.Scrape()
	if err != nil {
		assert.Failf(t, "Scraper.Scrape() is not supposed to return an error", "%v", err)
	}
	assert.Equal(t, mockDownloadLink, resp, "Expecting a specific url to be scraped")
	mockHttpClient.AssertDoInput(t, func(r *http.Request) {
		assert.Equal(
			t, mockSourceUrl, r.URL.String(),
			"The request should be made to the mocked source url")
	})
}

// Tests Scraper.Scrape() where the HttpClient returned an error
func TestScrapeWithError(t *testing.T) {
	scraper, mockHttpClient := createScraperWithMockedHttpClient(true)
	_, err := scraper.Scrape()
	if assert.Error(t, err) {
		assert.Equal(t, errMockBadResponse, err)
	}
	mockHttpClient.AssertDoWasCalled(t, "HttpClient.Do() should have been called exactly once")
}

func createScraperWithMockedHttpClient(withError bool) (Scraper, *mocks.MockHttpClient) {
	client := createMockHttpClient(withError)
	return New(client, mockSourceUrl), client
}

func createMockHttpClient(withError bool) *mocks.MockHttpClient {
	return mocks.
		NewMockHttpClient().
		WithDo(func(*http.Request) (*http.Response, error) {
			var err error
			if withError {
				err = errMockBadResponse
			}
			return &http.Response{
				Status:        "200 OK",
				Body:          io.NopCloser(bytes.NewBufferString(mockDownloadPage)),
				ContentLength: int64(len(mockDownloadPage)),
				Header: http.Header{
					http.CanonicalHeaderKey("Content-Type"): []string{"text/html; charset=UTF-8"},
				},
			}, err
		})
}
