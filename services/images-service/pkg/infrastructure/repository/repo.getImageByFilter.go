package repository

import (
	"context"
	"service-images/pkg/domain/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) GetImagesFilter(filter models.ImageFilter, page int) (models.ImageList, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	filterQuery := make(map[string]interface{})

	if len(filter.Tag) != 0 {
		filterQuery["tag"] = filter.Tag
	}

	if len(filter.Format) != 0 {
		filterQuery["metadata.format"] = filter.Format
	}

	if len(filter.Owner) != 0 {
		id, err := uuid.Parse(filter.Owner)
		if err != nil {
			return nil, err
		}
		filterQuery["owner"] = id
	}

	if len(filter.Name) != 0 {

		filterQuery["name"] = filter.Name
	}

	images := models.ImageList{}
	options := options.Find()
	options.SetLimit(20)
	options.SetSkip(20 * (int64(page) - 1))

	cur, err := collection.Find(ctx, filterQuery, options)
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &images); err != nil {
		return nil, err
	}
	return images, nil
}
