package db

import (
	"context"
	"fmt"

	_ "embed"

	"github.com/micheledinelli/aculei-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ArchiveRepository struct {
	mongo *Mongo
}

func NewArchiveRepository(mongo *Mongo) *ArchiveRepository {
	return &ArchiveRepository{
		mongo: mongo,
	}
}

func (r *ArchiveRepository) GetArchiveList(ctx context.Context) (*[]models.Archive, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	var archiveList []models.Archive

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting archive list: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var res bson.M

		if err := cursor.Decode(&res); err != nil {
			return nil, fmt.Errorf("error decoding archive list: %w", err)
		}

		// for k, v := range res {
		// 	fmt.Printf("Key: %s, Value: %v\n", k, v)
		// }

		// Convert bson.M to models.Archive
		archive := models.Archive{}
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
