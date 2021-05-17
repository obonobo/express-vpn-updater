package config

var (
	defaults *Config = &Config{
		SERVICE_URL: "https://qa5jmmkml3.execute-api.us-east-1.amazonaws.com/dev",
	}
)

func Defaults() Config {
	return defaults.Copy()
}
