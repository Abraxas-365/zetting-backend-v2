package repository

import (
	"context"
	"notifications/pkg/notification/domain/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) FindByUserID(userId uuid.UUID) (models.NotificationsQuery, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var notifications models.NotificationsQuery
	filter := bson.M{"reciver": userId}
	if err := collection.FindOne(ctx, filter).Decode(&notifications); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.NotificationsQuery{}, false, nil
		}
		return models.NotificationsQuery{}, false, err
	}
	return notifications, true, nil
}
