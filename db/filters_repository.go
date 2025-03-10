package db

import (
	"context"
	"fmt"
	"time"

	_ "embed"

	"github.com/micheledinelli/aculei-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

type FiltersRepository struct {
	mongo *Mongo
}

func NewFiltersRepository(mongo *Mongo) *FiltersRepository {
	return &FiltersRepository{
		mongo: mongo,
	}
}

func (r *FiltersRepository) GetAvailableFilters(ctx context.Context) (*[]models.Filter, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting archive list: %w", err)
	}
	defer cursor.Close(ctx)

	var maxTemperature float64
	var minTemperature float64
	minDate := time.Now()
	maxDate := time.Time{}
	animals := make(map[string]string)
	cameras := make(map[string]string)
	moonPhases := make(map[string]string)

	for cursor.Next(ctx) {
		var aculeiImage models.AculeiImage

		if err := cursor.Decode(&aculeiImage); err != nil {
			return nil, fmt.Errorf("error decoding archive list: %w", err)
		}

		if _, ok := animals[aculeiImage.PredictedAnimal]; !ok {
			animals[aculeiImage.PredictedAnimal] = aculeiImage.PredictedAnimal
		}

		if aculeiImage.Temperature != nil {
			if *aculeiImage.Temperature > maxTemperature {
				maxTemperature = *aculeiImage.Temperature
			}

			if *aculeiImage.Temperature < minTemperature {
				minTemperature = *aculeiImage.Temperature
			}
		}

		if aculeiImage.Cam != nil {
			if _, ok := cameras[*aculeiImage.Cam]; !ok {
				cameras[*aculeiImage.Cam] = *aculeiImage.Cam
			}
		}

		if aculeiImage.MoonPhase != nil {
			if _, ok := moonPhases[*aculeiImage.MoonPhase]; !ok {
				moonPhases[*aculeiImage.MoonPhase] = *aculeiImage.MoonPhase
			}
		}

		if aculeiImage.Date != nil {
			date, err := time.Parse("2006-01-02 15:04:05", *aculeiImage.Date)
			if err != nil {
				return nil, fmt.Errorf("error parsing date: %w", err)
			}

			if date.After(maxDate) {
				maxDate = date
			}

			if date.Before(minDate) {
				minDate = date
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating archive list: %w", err)
	}

	animalsAvailable := make([]interface{}, 0, len(animals))
	for _, animal := range animals {
		animalsAvailable = append(animalsAvailable, animal)
	}

	moonPhasesAvailable := make([]interface{}, 0, len(moonPhases))
	for _, moonPhase := range moonPhases {
		moonPhasesAvailable = append(moonPhasesAvailable, moonPhase)
	}

	camerasAvailable := make([]interface{}, 0, len(cameras))
	for _, camera := range cameras {
		camerasAvailable = append(camerasAvailable, camera)
	}

	filters := []models.Filter{
		{
			Name:   "temperatures",
			Values: nil,
			From:   minTemperature,
			To:     maxTemperature,
		},
		{
			Name:   "animals",
			Values: &animalsAvailable,
		},
		{
			Name:   "moon_phases",
			Values: &moonPhasesAvailable,
		},
		{
			Name:   "cameras",
			Values: &camerasAvailable,
		},
		{
			Name:   "dates",
			Values: nil,
			From:   minDate.Format("2006-01-02"),
			To:     maxDate.Format("2006-01-02"),
		},
	}

	return &filters, nil
}
