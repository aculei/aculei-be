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
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)
	pipeline := bson.A{
		bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error retrieving random image: %w", err)
	}
	defer cursor.Close(ctx)

	var aculeiImage models.AculeiImage
	if cursor.Next(ctx) {
		if err := cursor.Decode(&aculeiImage); err != nil {
			return nil, fmt.Errorf("error decoding archive image: %w", err)
		}
		return &aculeiImage, nil
	}

	return nil, fmt.Errorf("no image found")
}
