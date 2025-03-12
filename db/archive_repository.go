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

func (r *ArchiveRepository) GetArchive(
	ctx context.Context,
	paginator models.Paginator,
	fg models.FilterGroup) (*[]models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	filters, err := fg.GenerateFilters()
	if err != nil {
		return nil, models.NewErrorFilter(err.Error())
	}

	var archiveList []models.AculeiImage

	findOptions := options.Find()
	findOptions.SetLimit(int64(paginator.Size))
	findOptions.SetSkip(int64(paginator.Size * paginator.Page))

	cursor, err := coll.Find(ctx, filters, findOptions)
	if err != nil {
		return nil, models.ErrorDatabaseFind
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var acueliImage models.AculeiImage

		if err := cursor.Decode(&acueliImage); err != nil {
			return nil, models.ErrorDatabaseImageDecoder
		}

		archiveList = append(archiveList, acueliImage)
	}

	if err := cursor.Err(); err != nil {
		return nil, models.ErrorDatabaseCursor
	}

	return &archiveList, nil
}

func (r *ArchiveRepository) GetArchiveCount(ctx context.Context, fg models.FilterGroup) (int, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	filters, _ := fg.GenerateFilters()

	count, err := coll.CountDocuments(ctx, filters)
	if err != nil {
		return 0, models.ErrorDatabaseCount
	}

	return int(count), nil
}

func (r *ArchiveRepository) GetArchiveImage(ctx context.Context, id string) (*models.AculeiImage, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)
	res := coll.FindOne(ctx, bson.D{{Key: "id", Value: id}})

	var img models.AculeiImage

	if err := res.Decode(&img); err != nil {
		return nil, models.ErrorDatabaseImageDecoder
	}

	return &img, nil
}
