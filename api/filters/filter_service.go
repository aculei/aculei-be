package filters

import (
	"context"

	"github.com/micheledinelli/aculei-be/db"
	"github.com/micheledinelli/aculei-be/models"
)

type Service struct {
	configuration models.Configuration
	mongo         *db.Mongo
	filtersRepo   *db.FiltersRepository
}

func NewService(
	configuration models.Configuration,
	mongo *db.Mongo,
	filtersRepo *db.FiltersRepository,
) *Service {
	return &Service{
		configuration: configuration,
		mongo:         mongo,
		filtersRepo:   filtersRepo,
	}
}

func (s *Service) GetFilters(ctx context.Context) (*[]models.Filter, error) {
	return s.filtersRepo.GetFilters(ctx)
}
