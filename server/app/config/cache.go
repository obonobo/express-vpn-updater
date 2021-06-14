package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	configPath string = "config.json"
)

var (
	cachedConfig  *Config
	defaultConfig Config = Config{
		MaxFileSize: 50000000,
		Bucket:      "prod.express-vpn-updater.ca",
		Url:         "https://www.expressvpn.com/latest-1#linux",
	}
)

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
		cachedConfig = cfg
	}
}

func parse(data []byte) (*Config, error) {
	cfg := Config{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return &cfg, nil
}
