package delivery_service

import "github.com/sirupsen/logrus"

type APIServer struct {
	config *DeliveryServiceConfig
	logger *logrus.Logger
}

func NewAPIServer(config *DeliveryServiceConfig) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}

	server.logger.Info("Starting the API Server")

	return nil
}

func (server *APIServer) configureLogger() (err error) {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return
	}

	server.logger.SetLevel(level)
	return nil
}
