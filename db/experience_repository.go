package db

import (
	"context"
	"fmt"

	_ "embed"

	"github.com/micheledinelli/aculei-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ExperienceRepository struct {
	mongo *Mongo
}

func NewExperienceRepository(mongo *Mongo) *ExperienceRepository {
	return &ExperienceRepository{
		mongo: mongo,
	}
}

func (r *ExperienceRepository) GetRandomExperienceImage(ctx context.Context) (*models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(experienceCollection)
	pipeline := bson.A{
		bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, models.ErrorDatabaseAggregate
	}
	defer cursor.Close(ctx)

	var img models.AculeiImage
	if cursor.Next(ctx) {
		if err := cursor.Decode(&img); err != nil {
			return nil, models.ErrorDatabaseImageDecoder
		}
		return &img, nil
	}

	return nil, fmt.Errorf("no image found")
}

func (r *ExperienceRepository) GetExperienceImage(ctx context.Context, id string) (*models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(experienceCollection)
	res := coll.FindOne(ctx, bson.D{{Key: "id", Value: id}})

	var img models.AculeiImage

	if err := res.Decode(&img); err != nil {
		return nil, models.ErrorDatabaseImageDecoder
	}

	return &img, nil
}
