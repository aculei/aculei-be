package db

import (
	"context"

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
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var acueliImage models.AculeiImage

		if err := cursor.Decode(&acueliImage); err != nil {
			return nil, err
		}

		archiveList = append(archiveList, acueliImage)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &archiveList, nil
}

func (r *ArchiveRepository) GetArchiveListCount(ctx context.Context) (int, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *ArchiveRepository) GetArchiveImage(ctx context.Context, imageId string) (*models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)
	res := coll.FindOne(ctx, bson.D{{Key: "id", Value: imageId}})

	var aculeiImage models.AculeiImage

	if err := res.Decode(&aculeiImage); err != nil {
		return nil, err
	}

	return &aculeiImage, nil
}
