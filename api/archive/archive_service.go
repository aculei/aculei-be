package archive

import (
	"context"

	"github.com/micheledinelli/aculei-be/db"
	"github.com/micheledinelli/aculei-be/models"
)

type Service struct {
	configuration models.Configuration
	mongo         *db.Mongo
	archiveRepo   *db.ArchiveRepository
}

func NewService(
	configuration models.Configuration,
	mongo *db.Mongo,
	archiveRepo *db.ArchiveRepository,
) *Service {
	return &Service{
		configuration: configuration,
		mongo:         mongo,
		archiveRepo:   archiveRepo,
	}
}

func (s *Service) GetArchive(ctx context.Context, p models.Paginator, fg models.FilterGroup) (*[]models.AculeiImage, error) {
	return s.archiveRepo.GetArchive(ctx, p, fg)
}

func (s *Service) GetArchiveCount(ctx context.Context, fg models.FilterGroup) (int, error) {
	return s.archiveRepo.GetArchiveCount(ctx, fg)
}

func (s *Service) GetArchiveImage(ctx context.Context, id string) (*models.AculeiImage, error) {
	return s.archiveRepo.GetArchiveImage(ctx, id)
}
