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

func (r *FiltersRepository) GetFilters(ctx context.Context) (*[]models.Filter, error) {
	coll := r.mongo.Client.Database(dbName).Collection(archiveCollection)

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, models.ErrorDatabaseFind
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
		var img models.AculeiImage

		if err := cursor.Decode(&img); err != nil {
			return nil, models.ErrorDatabaseImageDecoder
		}

		if _, ok := animals[img.PredictedAnimal]; !ok {
			animals[img.PredictedAnimal] = img.PredictedAnimal
		}

		if img.Temperature != nil {
			if *img.Temperature > maxTemperature {
				maxTemperature = *img.Temperature
			}

			if *img.Temperature < minTemperature {
				minTemperature = *img.Temperature
			}
		}

		if img.Cam != nil {
			if _, ok := cameras[*img.Cam]; !ok {
				cameras[*img.Cam] = *img.Cam
			}
		}

		if img.MoonPhase != nil {
			if _, ok := moonPhases[*img.MoonPhase]; !ok {
				moonPhases[*img.MoonPhase] = *img.MoonPhase
			}
		}

		if img.Date != nil {
			date, err := time.Parse(time.RFC3339, *img.Date)
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
		return nil, models.ErrorDatabaseCursor
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
