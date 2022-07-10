package repository

import (
	"context"
	"service-users/pkg/user/domain/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) GetPublicUser(uuid uuid.UUID) (models.UserPublic, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	var user models.UserPublic
	filter := bson.M{"_id": uuid}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return models.UserPublic{}, false, ErrUserNotFound
	}
	return user, true, nil
}
