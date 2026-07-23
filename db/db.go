package db

import (
	"context"

	"github.com/WesleySouzaDev/cv-manager/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(databaseName string) (*mongo.Database, error) {
	env := config.GetEnv()
	clientOptions := options.Client().ApplyURI(env.DBURL)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(databaseName)
	return db, nil
}
