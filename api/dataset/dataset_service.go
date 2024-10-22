package dataset

import (
	"github.com/micheledinelli/aculei-be/db"
	"github.com/micheledinelli/aculei-be/models"
)

type Service struct {
	configuration models.Configuration
	repsositories *db.DBRepository
}

func NewService(
	configuration models.Configuration,
	repository *db.DBRepository,
) *Service {
	return &Service{
		configuration: configuration,
		repsositories: repository,
	}
}

func (s *Service) GetDatasetInfo() (dataset *models.DatasetInfo, err error) {
	// return s.repsositories.Dataset.GetDatasetInfo()
	return &models.DatasetInfo{}, nil
}
