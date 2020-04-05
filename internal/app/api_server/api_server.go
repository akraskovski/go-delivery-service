package api_server

import (
	local_mongo "github.com/akraskovski/go-delivery-service/internal/app/store/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func Start(config *APIServerConfig) error {
	db, err := connectToDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	store := local_mongo.New(db)

	server := newServer(store, config)
	server.logger.Info("Starting the API Server")
	return http.ListenAndServe(config.BindPort, server.router)
}

func connectToDB(databaseUrl string) (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI(databaseUrl))
}
