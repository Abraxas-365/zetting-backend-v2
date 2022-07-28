package repository

import (
	"context"
	"notifications/pkg/notification/domain/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) FindByUserID(userId uuid.UUID, page int) (models.NotificationsQuery, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	notificationsQuery := models.NotificationsQuery{}

	optionsLimit := bson.D{{Key: "$limit", Value: 20}}
	optionsSkip := bson.D{{Key: "$skip", Value: page - 1}}
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "reciver", Value: userId}}}}
	lookupEmitter := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "NotificationUsers"}, {Key: "localField", Value: "emitter"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "emitter"}}}}
	lookupResiver := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "NotificationUsers"}, {Key: "localField", Value: "reciver"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "reciver"}}}}
	unwindEmitter := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$emitter"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	unwindResiver := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$reciver"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{optionsLimit, optionsSkip, matchStage,
		lookupEmitter, lookupResiver,
		unwindEmitter, unwindResiver})
	if err = cur.All(ctx, &notificationsQuery); err != nil {
		return nil, false, err
	}

	return notificationsQuery, true, nil
}
