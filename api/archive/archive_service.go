package archive

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/micheledinelli/aculei-be/db"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Service struct {
	configuration models.Configuration
	mongo         *db.Mongo
	archiveRepo   *db.ArchiveRepository
	rdb           *redis.Client
}

func NewService(
	configuration models.Configuration,
	mongo *db.Mongo,
	archiveRepo *db.ArchiveRepository,
	rdb *redis.Client,
) *Service {
	return &Service{
		configuration: configuration,
		mongo:         mongo,
		archiveRepo:   archiveRepo,
		rdb:           rdb,
	}
}

func (s *Service) GetArchive(ctx context.Context, p models.Paginator, fg models.FilterGroup) (*[]models.AculeiImage, error) {
	var archive []models.AculeiImage
	cacheKey := getCacheKey(p, fg)

	cachedData, err := s.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil && len(cachedData) > 0 {
		log.Info().Msgf("cache hit for key %s", cacheKey)
		if jsonErr := json.Unmarshal(cachedData, &archive); jsonErr == nil {
			return &archive, nil
		}
	}

	archivePtr, err := s.archiveRepo.GetArchive(ctx, p, fg)
	if err != nil {
		return nil, err
	}
	archive = *archivePtr

	data, jsonErr := json.Marshal(archive)
	if jsonErr != nil {
		return &archive, nil
	}

	log.Info().Msgf("cache miss for key %s, setting it", cacheKey)
	s.rdb.Set(ctx, cacheKey, data, time.Minute)
	return &archive, nil
}

func (s *Service) GetArchiveCount(ctx context.Context, fg models.FilterGroup) (int, error) {
	return s.archiveRepo.GetArchiveCount(ctx, fg)
}

func (s *Service) GetArchiveImage(ctx context.Context, id string) (*models.AculeiImage, error) {
	return s.archiveRepo.GetArchiveImage(ctx, id)
}

func getCacheKey(p models.Paginator, fg models.FilterGroup) string {
	s := "archive"
	if fg.Animals != nil {
		s = fmt.Sprintf(s+"-%v", *fg.Animals)
	}
	if fg.Dates != nil {
		s = fmt.Sprintf(s+"-%v", *fg.Dates)
	}
	if fg.MoonPhases != nil {
		s = fmt.Sprintf(s+"-%v", *fg.MoonPhases)
	}
	if fg.Temperatures != nil {
		s = fmt.Sprintf(s+"-%v", *fg.Temperatures)
	}
	s = fmt.Sprintf(s+"-%d-%d-%d", p.Page, p.Size, p.SortBy)
	return s
}
