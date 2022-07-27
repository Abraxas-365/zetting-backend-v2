package repository

import (
	"connection-service/pkg/connection/domain/models"
	"context"
)

func (r *repository) Create(new models.ConnectionInput) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return false, err
	}
	return true, nil
}
