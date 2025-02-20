package db

import (
	"context"
	"fmt"

	_ "embed"

	"github.com/micheledinelli/aculei-be/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArchiveRepository struct {
	mongo *Mongo
}

func NewArchiveRepository(mongo *Mongo) *ArchiveRepository {
	return &ArchiveRepository{
		mongo: mongo,
	}
}

func (r *ArchiveRepository) GetArchiveList(
	ctx context.Context,
	paginator models.Paginator) (*[]models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	var archiveList []models.AculeiImage

	findOptions := options.Find()
	findOptions.SetLimit(int64(paginator.Size))
	findOptions.SetSkip(int64(paginator.Size * paginator.Page))

	cursor, err := coll.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error getting archive list: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var res bson.M

		if err := cursor.Decode(&res); err != nil {
			return nil, fmt.Errorf("error decoding archive list: %w", err)
		}

		archive := models.AculeiImage{}
		archive.Id = res["id"].(string)
		archive.Cam = res["cam"].(string)
		archive.PredictedAnimal = res["predicted_animal"].(string)
		archive.ImageName = res["image_name"].(string)

		archiveList = append(archiveList, archive)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating archive list: %w", err)
	}

	return &archiveList, nil
}

func (r *ArchiveRepository) GetArchiveListCount(ctx context.Context) (int, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, fmt.Errorf("error counting archive list: %w", err)
	}

	return int(count), nil
}

func (r *ArchiveRepository) GetArchiveImage(ctx context.Context, imageId string) (*models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)
	res := coll.FindOne(ctx, bson.D{{Key: "id", Value: imageId}})

	var aculeiImage models.AculeiImage

	if err := res.Decode(&aculeiImage); err != nil {
		return nil, fmt.Errorf("error decoding archive image: %w", err)
	}

	return &aculeiImage, nil
}
