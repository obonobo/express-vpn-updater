package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/obonobo/express-vpn-updater/cli/config"
)

var (
	cfg config.Config = config.Defaults()
)

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{url: url}
}

func Default() *Client {
	return New(cfg.SERVICE_URL)
}

func (c *Client) DownloadLatest() (path string, err error) {
	url, filename, err := ScrapeDownloadLink()
	if err != nil {
		return "", err
	}

	got, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer got.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, got.Body)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func ScrapeDownloadLink() (url string, filename string, err error) {
	api := cfg.SERVICE_URL + "/latest?redirect=false"
	var data map[string]string

	got, err := http.Get(api)
	if err != nil {
		return "", "", err
	}
	defer got.Body.Close()

	body, err := ioutil.ReadAll(got.Body)
	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", err
	}
	if msg, ok := data["message"]; ok {
		return msg, path.Base(msg), nil
	}
	return "", "", err
}
