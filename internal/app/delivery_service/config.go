package delivery_service

type DeliveryServiceConfig struct {
	BindPort string `toml:"bind_port"`
}

func NewDeliveryServiceConfig() *DeliveryServiceConfig {
	return &DeliveryServiceConfig{
		BindPort: ":8000",
	}
}
