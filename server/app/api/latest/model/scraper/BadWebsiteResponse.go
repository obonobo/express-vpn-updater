package scraper

import (
	"fmt"
	"net/http"
)

const BadWebsiteErrorMessage string = "got a bad response (status=%d) from expressvpn website"

type BadWebsiteError struct {
	msg string
}

func NewBadWesiteError(resp *http.Response) BadWebsiteError {
	return BadWebsiteError{
		msg: badWebsiteErrorMessage(resp.StatusCode),
	}
}

func (b BadWebsiteError) Error() string {
	return b.msg
}

func badWebsiteErrorMessage(statusCode int) string {
	return fmt.Sprintf(BadWebsiteErrorMessage, statusCode)
}
