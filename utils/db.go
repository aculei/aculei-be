package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

type Configuration struct {
	User             string  `yaml:"user"`
	Pass             string  `yaml:"password"`
	Host             string  `yaml:"host"`
	Port             string  `yaml:"port"`
	Name             string  `yaml:"name"`
	MaxConns         int     `yaml:"max_conns"`
	MaxConnIdleTime  *string `yaml:"max_conn_idle_time"`
	MaxConnLifetime  *string `yaml:"max_conn_lifetime"`
	TraceServiceName string  `yaml:"trace_service_name"`
}

func New(ctx context.Context, conf Configuration) (*DB, error) {
	pgConfig, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s",
			conf.User,
			conf.Pass,
			conf.Host,
			conf.Port,
			conf.Name,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing db connection string: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgConfig)
	if err != nil {
		return nil, fmt.Errorf("error initializing connection: %w", err)
	}

	return &DB{
		pool,
	}, nil
}
