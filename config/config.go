package config

var cfg *Config //nolint:gochecknoglobals

func init() { //nolint:gochecknoinits
	cfg = DefaultConfig()
}

type Config struct {
	App *AppConfig `mapstructure:"app"`
}

func GetConfig() *Config {
	return cfg
}

func DefaultConfig() *Config {
	return &Config{
		App: &AppConfig{},
	}
}
