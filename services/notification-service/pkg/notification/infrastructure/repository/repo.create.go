package repository

import (
	"context"
	"notifications/pkg/notification/domain/models"
)

func (r *repository) Create(new models.NotificationInput) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return false, err
	}
	return true, nil
}
