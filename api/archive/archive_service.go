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

func (s *Service) GetArchiveList(ctx context.Context, paginator models.Paginator, filterGroup models.FilterGroup) (
	*[]models.AculeiImage, error) {
	return s.archiveRepo.GetArchiveList(ctx, paginator, filterGroup)
}

func (s *Service) GetArchiveListCount(ctx context.Context, filterGroup models.FilterGroup) (int, error) {
	return s.archiveRepo.GetArchiveListCount(ctx, filterGroup)
}

func (s *Service) GetArchiveImage(ctx context.Context, imageId string) (*models.AculeiImage, error) {
	return s.archiveRepo.GetArchiveImage(ctx, imageId)
}
