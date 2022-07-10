package repository

import (
	"context"
	"service-users/pkg/user/domain/events"
	"service-users/pkg/user/domain/models"
)

func (r *repository) Create(new models.User) (events.Event, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return events.Event{}, false, err
	}

	return events.NewEvent(events.UserCreated{
		UserId:   new.ID,
		UserName: new.Name,
	}), true, nil
}
