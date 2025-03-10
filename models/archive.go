package models

import (
	"errors"
	"math"

	"go.mongodb.org/mongo-driver/bson"
)

type AculeiImage struct {
	Id              string   `bson:"id" json:"id"`
	ImageName       string   `bson:"image_name" json:"image_name"`
	PredictedAnimal string   `bson:"predicted_animal" json:"predicted_animal"`
	MoonPhase       *string  `bson:"moon_phase,omitempty" json:"moon_phase"`
	Temperature     *float64 `bson:"temperature,omitempty" json:"temperature"`
	Date            *string  `bson:"date,omitempty" json:"date"`
	Cam             *string  `bson:"cam,omitempty" json:"cam"`
}

func (a *AculeiImage) UnmarshalBSON(data []byte) error {
	var raw bson.M
	if err := bson.Unmarshal(data, &raw); err != nil {
		return err
	}

	if id, ok := raw["id"]; ok {
		switch v := id.(type) {
		case string:
			a.Id = v
		default:
			return errors.New("id is not a string")
		}
	}

	if imageName, ok := raw["image_name"]; ok {
		switch v := imageName.(type) {
		case string:
			a.ImageName = v
		default:
			return errors.New("image_name is not a string")
		}
	}

	if predictedAnimal, ok := raw["predicted_animal"]; ok {
		switch v := predictedAnimal.(type) {
		case string:
			a.PredictedAnimal = v
		default:
			return errors.New("predicted_animal is not a string")
		}
	}

	if moonPhase, ok := raw["moon_phase"]; ok {
		switch v := moonPhase.(type) {
		case string:
			a.MoonPhase = &v
		default:
			a.MoonPhase = nil
		}
	}

	if temp, ok := raw["temperature"]; ok {
		switch v := temp.(type) {
		case float64:
			if math.IsNaN(v) {
				a.Temperature = nil
			} else {
				a.Temperature = &v
			}
		case int32:
			t := float64(v)
			a.Temperature = &t
		case int64:
			t := float64(v)
			a.Temperature = &t
		default:
			a.Temperature = nil
		}
	}

	if date, ok := raw["date"]; ok {
		switch v := date.(type) {
		case string:
			a.Date = &v
		default:
			a.Date = nil
		}
	}

	if cam, ok := raw["cam"]; ok {
		switch v := cam.(type) {
		case string:
			a.Cam = &v
		default:
			a.Cam = nil
		}
	}

	return nil
}
