package experience

import (
	"context"

	"github.com/micheledinelli/aculei-be/db"
	"github.com/micheledinelli/aculei-be/models"
)

type Service struct {
	configuration  models.Configuration
	mongo          *db.Mongo
	experienceRepo *db.ExperienceRepository
}

func NewService(
	configuration models.Configuration,
	mongo *db.Mongo,
	experienceRepo *db.ExperienceRepository,
) *Service {
	return &Service{
		configuration:  configuration,
		mongo:          mongo,
		experienceRepo: experienceRepo,
	}
}

func (s *Service) GetRandomExperienceImage(ctx context.Context) (*models.AculeiImage, error) {
	return s.experienceRepo.GetRandomExperienceImage(ctx)
}

func (s *Service) GetExperienceImage(ctx context.Context, imageId string) (*models.AculeiImage, error) {
	return s.experienceRepo.GetExperienceImage(ctx, imageId)
}
