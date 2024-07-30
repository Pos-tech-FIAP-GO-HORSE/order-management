package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(ctx context.Context, uri string, opts *options.ClientOptions) (*mongo.Client, error) {
	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return mongoClient, nil
}
