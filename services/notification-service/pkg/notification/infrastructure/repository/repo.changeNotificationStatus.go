package repository

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) ChangeNotificationStatus(notificationId uuid.UUID, isSeen bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	filter := bson.M{"_id": notificationId}
	updateQuery := bson.M{
		"$set": bson.M{"status": isSeen},
	}
	_, err := collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err

	}
	return nil
}
