package store

import (
	"fmt"
	"net/http"

	"github.com/obonobo/express-vpn-updater/server/config"
)

var maxFileSize int64 = config.Get().MaxFileSize

type FileTooBigError struct {
	FileSize int64
}

func NewFileTooBigError(res *http.Response) FileTooBigError {
	return FileTooBigError{FileSize: res.ContentLength}
}

func (e FileTooBigError) Error() string {
	return fmt.Sprintf(
		"attempted too big of a file - filesize=%d, max=%d",
		e.FileSize,
		maxFileSize,
	)
}

func exceedsMaxFileSize(res *http.Response) bool {
	return res.ContentLength > maxFileSize
}
