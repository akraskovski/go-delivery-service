package api_server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type server struct {
	config *APIServerConfig
	router *mux.Router
	logger *logrus.Logger
}

func newServer(config *APIServerConfig) (*server, error) {
	server := server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

	if err := server.configureLogger(); err != nil {
		return nil, err
	}

	server.configureRouter()

	return &server, nil
}

func (server *server) configureLogger() (err error) {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return
	}

	server.logger.SetLevel(level)
	return nil
}

func (server *server) configureRouter() {
	server.router.HandleFunc("/hello", handleHello())
}

func handleHello() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "Welcome to the club, buddy\n")
	}
}
