package apiserver

import (
	"encoding/json"
	"github.com/akraskovski/go-delivery-service/internal/app/model"
	"net/http"
	"time"
)

func (server *server) initOrderController() {
	server.router.HandleFunc("/orders", server.handleCreateOrder()).Methods("POST")
}

func (server *server) handleCreateOrder() http.HandlerFunc {
	type requestForm struct {
		Name           string    `json:"name"`
		DeliverAddress string    `json:"deliverAddress"`
		DeliverTime    time.Time `json:"deliverTime"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		req := &requestForm{}
		if err := json.NewDecoder(request.Body).Decode(req); err != nil {
			server.error(writer, http.StatusBadRequest, err)
			return
		}

		order := model.Order{
			Name:           req.Name,
			DeliverAddress: req.DeliverAddress,
			DeliverTime:    req.DeliverTime,
		}
		id, err := server.store.Order().Create(&order)
		if err != nil {
			server.error(writer, http.StatusUnprocessableEntity, err)
			return
		}

		server.respond(writer, http.StatusCreated, id)
	}
}
