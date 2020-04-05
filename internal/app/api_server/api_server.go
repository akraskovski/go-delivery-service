package api_server

import (
	"github.com/akraskovski/go-delivery-service/internal/app/store/repository/mongo"
	"net/http"
)

func Start(config *APIServerConfig) error {
	db, err := connectToDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	store := mongo.New(db)

	server := newServer(store, config)
	server.logger.Info("Starting the API Server")
	return http.ListenAndServe(config.BindPort, server.router)
}

//todo: implement function
func connectToDB(databaseUrl string) (*interface{}, error) {
	return nil, nil
}
