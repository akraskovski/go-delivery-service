package apiserver

import (
	"context"
	"github.com/akraskovski/go-delivery-service/internal/app/store/repository/mongo"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func Start(config *APIServerConfig) error {
	db, err := connectToDB(config)
	if err != nil {
		return err
	}

	store := mongo.New(db)

	server := newServer(store, config)
	server.logger.Info("Starting the API Server")
	return http.ListenAndServe(config.BindPort, server.router)
}

func connectToDB(config *APIServerConfig) (*mongodriver.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongodriver.Connect(ctx, options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(config.DatabaseName), nil
}
