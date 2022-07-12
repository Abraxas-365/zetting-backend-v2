package repository

import (
	"context"
	"service-images/pkg/domain/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) GetImagesByTag(tag string) (models.ImageList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"tag": tag}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var listImages models.ImageList
	if err = cur.All(ctx, &listImages); err != nil {
		return nil, err
	}
	return listImages, nil
}
