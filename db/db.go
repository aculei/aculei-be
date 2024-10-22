package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/micheledinelli/aculei-be/utils"
)

type DB struct {
	pool *pgxpool.Pool
}

type DBRepository struct {
	Utils   DBRepositoryUtils
	Dataset DatasetRepository
}

type DBRepositoryUtils struct {
	CreateTransaction func(context.Context) (pgx.Tx, error)
}

func (db *DB) InitRepositories() *DBRepository {
	datasetRepository := NewDatasetRepository(db)
	return &DBRepository{
		Dataset: *datasetRepository,
		Utils: DBRepositoryUtils{
			CreateTransaction: db.CreateTransaction,
		},
	}
}

func NewDb(ctx context.Context, configuration models.Configuration) (*DB, error) {
	dbPool, err := utils.New(ctx, utils.Configuration{
		User:             configuration.DB.User,
		Pass:             configuration.DB.Pass,
		Host:             configuration.DB.Host,
		Port:             configuration.DB.Port,
		Name:             configuration.DB.Name,
		MaxConns:         0,
		TraceServiceName: "aculei-postgres",
	})

	if err != nil {
		return nil, err
	}

	return &DB{
		pool: dbPool.Pool,
	}, nil
}

func (db *DB) CreateTransaction(ctx context.Context) (pgx.Tx, error) {
	return db.pool.Begin(ctx)
}
