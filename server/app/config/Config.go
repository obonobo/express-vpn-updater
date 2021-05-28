package config

import "github.com/obonobo/express-vpn-updater/server/app/util/logging"

type Config struct {
	MaxFileSize int64
	Bucket      string
	Url         string
	logger      *logging.Logger
}

func (c Config) Logger() *logging.Logger {
	if c.logger == nil {
		c.logger = logging.New()
	}
	return c.logger
}
