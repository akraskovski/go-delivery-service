package apiserver

import (
	"context"
	"github.com/akraskovski/go-delivery-service/internal/app/apiserver/config"
	"github.com/akraskovski/go-delivery-service/internal/app/store/repository/mongo"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func Start(cfg *config.Config) error {
	db, err := connectToDB(cfg)
	if err != nil {
		return err
	}

	store := mongo.New(db)

	server := newServer(store)
	server.logger.Info("Starting the API Server")
	return http.ListenAndServe(cfg.BindPort, server.router)
}

func connectToDB(cfg *config.Config) (*mongodriver.Database, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	client, err := mongodriver.Connect(ctx, options.Client().ApplyURI(cfg.DatabaseURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(cfg.DatabaseName), nil
}
