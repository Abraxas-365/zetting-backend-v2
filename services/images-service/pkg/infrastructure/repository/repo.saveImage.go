package repository

import (
	"context"
	"service-images/pkg/domain/events"
	"service-images/pkg/domain/models"
)

func (r *repository) PostImage(new models.Image) (events.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return events.Event{}, err
	}

	return events.NewEvent(events.ImagePosted{
		ImageName: new.Name,
		UserId:    new.Owner,
		ImageId:   new.ID,
		ImageEtag: new.EntityTag,
	}), nil

}
