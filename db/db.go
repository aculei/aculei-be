package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "aculei"
const archiveCollection = "archive"
const experienceCollection = "experience"

type Mongo struct {
	Client *mongo.Client
	DbInfo DatabaseInfo
}

type DatabaseInfo struct {
	DatabaseName string
}

func InitDatabase(ctx context.Context, mongoUri string) (*Mongo, error) {
	opts := options.Client().ApplyURI(mongoUri).
		SetConnectTimeout(30 * time.Second).
		SetServerSelectionTimeout(30 * time.Second).
		SetSocketTimeout(30 * time.Second).
		SetMaxPoolSize(100).
		SetMinPoolSize(1)

	m := &Mongo{}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo: %w", err)
	}

	m.Client = client
	m.DbInfo.DatabaseName = dbName

	return m, nil
}

type DBRepository struct {
	Archive    ArchiveRepository
	Experience ExperienceRepository
	Filters    FiltersRepository
}

func (db *Mongo) InitRepositories() *DBRepository {
	archiveRepo := NewArchiveRepository(db)
	experienceRepo := NewExperienceRepository(db)
	filtersRepo := NewFiltersRepository(db)
	return &DBRepository{
		Archive:    *archiveRepo,
		Experience: *experienceRepo,
		Filters:    *filtersRepo,
	}
}
