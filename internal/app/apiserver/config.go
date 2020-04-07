package apiserver

type APIServerConfig struct {
	BindPort     string `toml:"bind_port"`
	LogLevel     string `toml:"log_level"`
	DatabaseURL  string `toml:"database_url"`
	DatabaseName string `toml:"database_name"`
}

func NewAPIServerConfig() *APIServerConfig {
	return &APIServerConfig{
		BindPort: ":8000",
		LogLevel: "debug",
	}
}
