package apiserver

import (
	"encoding/json"
	"github.com/akraskovski/go-delivery-service/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	ctxKeyRequestID = iota
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	server := server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}
	server.configureRouter()

	return &server
}

func (server *server) configureRouter() {
	server.router.Use(server.logRequest)
	server.initOrderController()
}

func (server *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := server.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (server *server) error(w http.ResponseWriter, code int, err error) {
	server.respond(w, code, map[string]string{"error": err.Error()})
}

func (server *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
