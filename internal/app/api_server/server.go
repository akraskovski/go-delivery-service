package api_server

import (
	"github.com/akraskovski/go-delivery-service/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store, config *APIServerConfig) *server {
	server := server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}

	server.configureLogger(config)
	server.configureRouter()

	return &server
}

func (server *server) configureLogger(config *APIServerConfig) (err error) {
	level, err := logrus.ParseLevel(config.LogLevel)
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
