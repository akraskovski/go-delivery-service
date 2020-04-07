package config

type Config struct {
	BindPort     string `toml:"bind_port"`
	LogLevel     string `toml:"log_level"`
	DatabaseURL  string `toml:"database_url"`
	DatabaseName string `toml:"database_name"`
}

func New() *Config {
	return &Config{
		BindPort: ":8000",
		LogLevel: "debug",
	}
}
