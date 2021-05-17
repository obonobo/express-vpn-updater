package argparse

import (
	"flag"

	"github.com/obonobo/express-vpn-updater/cli/config"
)

var (
	cfg config.Config = config.Defaults()
)

type Arguments struct {
	Url string
}

func Parse() Arguments {
	args := Arguments{Url: cfg.SERVICE_URL}
	url := flag.String(
		"url",
		cfg.SERVICE_URL,
		"Overrides the default download url for .deb package",
	)

	flag.Parse()
	if url != nil {
		args.Url = *url
	}
	return args
}
