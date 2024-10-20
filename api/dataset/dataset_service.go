package dataset

import "github.com/micheledinelli/aculei-be/models"

type Service struct {
	configuration models.Configuration
}

func NewService(
	configuration models.Configuration,
) *Service {
	return &Service{
		configuration: configuration,
	}
}

func (s *Service) GetDataset() (dataset *models.Dataset, err error) {
	dataset = &models.Dataset{
		Fox: "test",
	}

	return dataset, nil
}
