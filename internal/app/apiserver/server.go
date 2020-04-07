package apiserver

import (
	"encoding/json"
	"github.com/akraskovski/go-delivery-service/internal/app/model"
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
	server.router.Use(server.logRequest)
	server.router.HandleFunc("/orders", server.handleCreateOrder()).Methods("POST")
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

func (server *server) handleCreateOrder() http.HandlerFunc {
	type requestForm struct {
		Id             string    `json:"id"`
		Name           string    `json:"name"`
		DeliverAddress string    `json:"deliverAddress"`
		DeliverTime    time.Time `json:"deliverTime"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		req := &requestForm{}
		if err := json.NewDecoder(request.Body).Decode(req); err != nil {
			server.error(writer, request, http.StatusBadRequest, err)
			return
		}

		order := model.Order{
			Id:             req.Id,
			Name:           req.Name,
			DeliverAddress: req.DeliverAddress,
			DeliverTime:    req.DeliverTime,
		}
		if _, err := server.store.Order().Create(&order); err != nil {
			server.error(writer, request, http.StatusUnprocessableEntity, err)
			return
		}

		server.respond(writer, request, http.StatusCreated, order)
	}
}

func (server *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	server.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (server *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
