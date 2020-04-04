package delivery_service

type DeliveryServiceConfig struct {
	BindPort string `toml:"bind_port"`
	LogLevel string `toml:"log_level"`
}

func NewDeliveryServiceConfig() *DeliveryServiceConfig {
	return &DeliveryServiceConfig{
		BindPort: ":8000",
		LogLevel: "debug",
	}
}
