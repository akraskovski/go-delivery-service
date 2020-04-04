package delivery_service

type APIServer struct {
	config *DeliveryServiceConfig
}

func NewAPIServer(config *DeliveryServiceConfig) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (server *APIServer) Start() error {
	return nil
}
