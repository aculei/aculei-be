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

func (r *ArchiveRepository) GetArchiveList(ctx context.Context, paginator models.Paginator, filterGroup models.FilterGroup) (
	*[]models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	filters, err := filterGroup.GenerateFilters()
	if err != nil {
		return nil, err
	}

	var archiveList []models.AculeiImage

	findOptions := options.Find()
	findOptions.SetLimit(int64(paginator.Size))
	findOptions.SetSkip(int64(paginator.Size * paginator.Page))

	cursor, err := coll.Find(ctx, filters, findOptions)
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

func (r *ArchiveRepository) GetArchiveListCount(ctx context.Context, filterGroup models.FilterGroup) (int, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	filters, _ := filterGroup.GenerateFilters()

	count, err := coll.CountDocuments(ctx, filters)
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
