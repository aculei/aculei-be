package db

import (
	"context"

	_ "embed"

	"github.com/micheledinelli/aculei-be/models"
)

type DatasetRepository struct {
	db *DB
}

func NewDatasetRepository(db *DB) *DatasetRepository {
	return &DatasetRepository{
		db: db,
	}
}

//go:embed sql/dataset/list_info.sql
// var listInfo string

func (r *DatasetRepository) GetDatasetInfo(ctx context.Context, target string, webhookId string) (*models.DatasetInfo, error) {
	datasetInfo := &models.DatasetInfo{}
	// var id, eventId, url, crn, targetType, targetId, status, eventTopic, signingKey string
	// var createdAt, updatedAt time.Time
	// var verifiedAt, deletedAt pgtype.Timestamp
	// var onFailRetry bool
	// var err error

	// err := r.db.pool.QueryRow(ctx, listInfo, nil).Scan(

	// )

	// if err != nil {
	// 	return nil, err
	// }

	// err := rows.Scan(&id, &url, &crn, &targetType, &targetId, &status, &signingKey, &onFailRetry, &createdAt, &updatedAt, &deletedAt, &verifiedAt, &eventId, &eventTopic)

	// if err != nil {
	// 	return nil, err
	// }

	return datasetInfo, nil
}
