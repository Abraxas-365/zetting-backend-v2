package repository

import (
	"context"
	"service-users/pkg/user/domain/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) GetPublicUser(uuid uuid.UUID) (models.UserPublic, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var user models.UserPublic
	filter := bson.M{"_id": uuid}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.UserPublic{}, false, nil
		}
		return models.UserPublic{}, false, err
	}
	return user, true, nil
}
