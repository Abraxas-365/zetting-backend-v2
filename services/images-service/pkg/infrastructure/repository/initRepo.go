package repository

import (
	"context"
	"errors"
	"service-images/pkg/domain/ports"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ErrConflict     = errors.New("Allready exist")
	ErrNotFound     = errors.New("Not found")
	ErrUserNotFound = errors.New("User not found")
)

type repository struct {
	client     *mongo.Client
	database   string
	timeout    time.Duration
	collection string
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int, collection string) (ports.Repository, error) {
	repo := &repository{
		timeout:    time.Duration(mongoTimeout) * time.Second,
		database:   mongoDB,
		collection: collection,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.New(err.Error() + " repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}