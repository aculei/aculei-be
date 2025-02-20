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

func (s *Service) GetArchiveList(ctx context.Context) (*[]models.AculeiImage, error) {
	return s.archiveRepo.GetArchiveList(ctx)
}

func (s *Service) GetArchiveImage(ctx context.Context, imageId string) (*models.AculeiImage, error) {
	return s.archiveRepo.GetArchiveImage(ctx, imageId)
}
