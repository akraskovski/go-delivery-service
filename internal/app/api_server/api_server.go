package api_server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *APIServerConfig
	logger *logrus.Logger
	router *mux.Router
}

func NewAPIServer(config *APIServerConfig) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}

	server.logger.Info("Starting the API Server")
	server.configureRouter()

	return http.ListenAndServe(server.config.BindPort, server.router)
}

func (server *APIServer) configureLogger() (err error) {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return
	}

	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", handleHello())
}

func handleHello() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "Welcome to the club, buddy\n")
	}
}
