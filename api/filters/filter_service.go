package filters

import (
	"context"
	"encoding/json"
	"time"

	"github.com/micheledinelli/aculei-be/db"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Service struct {
	configuration models.Configuration
	mongo         *db.Mongo
	filtersRepo   *db.FiltersRepository
	rdb           *redis.Client
}

func NewService(
	configuration models.Configuration,
	mongo *db.Mongo,
	filtersRepo *db.FiltersRepository,
	rdb *redis.Client,
) *Service {
	return &Service{
		configuration: configuration,
		mongo:         mongo,
		filtersRepo:   filtersRepo,
		rdb:           rdb,
	}
}

func (s *Service) GetFilters(ctx context.Context) (*[]models.Filter, error) {
	var filters []models.Filter

	cacheKey := "filters"

	cache, err := s.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil && len(cache) > 0 {
		log.Info().Msgf("cache hit for key %s", cacheKey)
		if jsonErr := json.Unmarshal(cache, &filters); jsonErr == nil {
			return &filters, nil
		}
	}

	filtersPtr, err := s.filtersRepo.GetFilters(ctx)
	if err != nil {
		return nil, err
	}
	filters = *filtersPtr

	data, jsonErr := json.Marshal(filters)
	if jsonErr != nil {
		return &filters, nil
	}

	log.Info().Msgf("cache miss for key %s, setting it", cacheKey)
	s.rdb.Set(ctx, cacheKey, data, time.Hour)
	return &filters, nil
}
