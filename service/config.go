package service

type Config struct {
	APIConfig struct {
		APISecret string `toml:"APISecret"`
		APIKey    string `toml:"APIKey"`
	} `toml:"APIConfig"`
}
