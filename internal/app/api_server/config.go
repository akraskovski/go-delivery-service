package api_server

type APIServerConfig struct {
	BindPort    string `toml:"bind_port"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewAPIServerConfig() *APIServerConfig {
	return &APIServerConfig{
		BindPort: ":8000",
		LogLevel: "debug",
	}
}
