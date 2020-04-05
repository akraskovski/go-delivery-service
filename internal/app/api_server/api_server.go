package api_server

import (
	"net/http"
)

func Start(config *APIServerConfig) error {
	server, err := newServer(config)
	if err != nil {
		return err
	}

	server.logger.Info("Starting the API Server")
	return http.ListenAndServe(server.config.BindPort, server.router)
}
