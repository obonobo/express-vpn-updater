package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	configPath string = "config.json"
)

var (
	cachedConfig  *Config
	defaultConfig Config = Config{
		MaxFileSize: 50000000,
		Bucket:      "express-vpn-deb-cache",
		Url:         "https://www.expressvpn.com/latest-1#linux",
	}

	loggerOut    = os.Stderr
	loggerPrefix = "<><><> "
	loggerFlags  = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
)

type Config struct {
	MaxFileSize int64
	Bucket      string
	Url         string
	logger      *log.Logger
}

func (c Config) Logger() *log.Logger {
	if c.logger == nil {
		c.logger = NewLogger()
	}
	return c.logger
}

func NewLogger() *log.Logger {
	return log.New(loggerOut, loggerPrefix, loggerFlags)
}

func Get() Config {
	if cachedConfig == nil {
		initConfig()
	}
	return *cachedConfig
}

func initConfig() {
	configFileData, err := ioutil.ReadFile(configPath)
	if err != nil {
		cachedConfig = &defaultConfig
	} else if cfg, err := parse(configFileData); err != nil {
		cachedConfig = &cfg
	}
}

func parse(data []byte) (Config, error) {
	cfg := Config{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
