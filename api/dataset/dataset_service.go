package dataset

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/micheledinelli/aculei-be/models"
)

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

func (s *Service) GetDataset() (dataset *models.DatasetInfoResponse, err error) {
	rows, err := readCsv()
	if err != nil {
		return nil, models.ErrorInternalServerError
	}

	dataset = &models.DatasetInfoResponse{
		Badger:       rows[0].Badger,
		Buzzard:      rows[0].Buzzard,
		Cat:          rows[0].Cat,
		Cam1:         rows[0].Cam1,
		Cam2:         rows[0].Cam2,
		Cam3:         rows[0].Cam3,
		Cam4:         rows[0].Cam4,
		Cam5:         rows[0].Cam5,
		Cam6:         rows[0].Cam6,
		Cam7:         rows[0].Cam7,
		Deer:         rows[0].Deer,
		Fox:          rows[0].Fox,
		Hare:         rows[0].Hare,
		Heron:        rows[0].Heron,
		Horse:        rows[0].Horse,
		Mallard:      rows[0].Mallard,
		Marten:       rows[0].Marten,
		PhotosAutumn: rows[0].PhotosAutumn,
		PhotosWinter: rows[0].PhotosWinter,
		PhotosSpring: rows[0].PhotosSpring,
		PhotosSummer: rows[0].PhotosSummer,
		Porcupine:    rows[0].Porcupine,
		Squirrel:     rows[0].Squirrel,
		TotalRecords: rows[0].TotalRecords,
		WildBoar:     rows[0].WildBoar,
		Wolf:         rows[0].Wolf,
	}

	return dataset, nil
}

func readCsv() ([]*models.DatasetInfo, error) {
	csvFile, csvFileError := os.OpenFile("db/aculei_info.csv", os.O_RDWR, os.ModePerm)
	if csvFileError != nil {
		return nil, csvFileError
	}
	defer csvFile.Close()

	var csvRows []*models.DatasetInfo
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &csvRows); unmarshalError != nil {
		return nil, unmarshalError
	}

	return csvRows, nil
}
