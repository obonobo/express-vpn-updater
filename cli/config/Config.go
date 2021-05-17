package config

type Config struct {
	SERVICE_URL string
}

func (cfg *Config) Copy() Config {
	return Config{
		SERVICE_URL: cfg.SERVICE_URL,
	}
}
