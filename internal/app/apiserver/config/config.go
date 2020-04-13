package config

type Config struct {
	BindPort     string `toml:"bind_port" envconfig:"BIND_PORT"`
	LogLevel     string `toml:"log_level" envconfig:"LOG_LEVEL"`
	DatabaseURL  string `toml:"database_url" envconfig:"DATABASE_URL"`
	DatabaseName string `toml:"database_name" envconfig:"DATABASE_NAME"`
}

func New() *Config {
	return &Config{
		BindPort: ":8000",
		LogLevel: "debug",
	}
}
