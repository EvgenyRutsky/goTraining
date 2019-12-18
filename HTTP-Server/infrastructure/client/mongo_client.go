package client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client interface {
	Open(ctx context.Context) (*mongo.Client, error)
}

type mongoClient struct {
	clientURI string
	dbname string
}

func NewClient() Client {
	return &mongoClient{
		clientURI: "mongodb://127.0.0.1:27017",
		dbname:    "dev",
	}
}

func (mc *mongoClient) Open(ctx context.Context) (*mongo.Client, error) {
	mclient, err := mongo.NewClient(options.Client().ApplyURI(mc.clientURI))
	if  err != nil {
		return nil, err
	}
	if err := mclient.Connect(ctx); err != nil {
		return nil, err
	}

	if err := mclient.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return mclient, nil
}

func Close(ctx context.Context, client *mongo.Client) error {
	if err := client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
