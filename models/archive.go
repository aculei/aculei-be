package models

import (
	"errors"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AculeiImage struct {
	Id              string   `bson:"id" json:"id" example:"76288dfbf134376e0b6fae8d8ff87c26"`
	ImageName       string   `bson:"image_name" json:"image_name" example:"TF_ACULEI_25012021-203.jpg"`
	PredictedAnimal string   `bson:"predicted_animal" json:"predicted_animal" example:"fox"`
	TopPredictions  string   `bson:"top_predictions" json:"top_predictions" example:"[{'score': 0.9460213780403137, 'label': 'porcupine'}, {'score': 0.03565983474254608, 'label': 'wild boar'}, {'score': 0.012196173891425133, 'label': 'badger'}]"`
	MoonPhase       *string  `bson:"moon_phase,omitempty" json:"moon_phase" example:"Waning Gibbous"`
	Temperature     *float64 `bson:"temperature,omitempty" json:"temperature" example:"12.5"`
	Date            *string  `bson:"date,omitempty" json:"date" example:"2021-01-25T03:01:32+01:00"`
	Cam             *string  `bson:"cam,omitempty" json:"cam" example:"CAM7"`
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

	if topPredictions, ok := raw["top_predictions"]; ok {
		switch v := topPredictions.(type) {
		case string:
			a.TopPredictions = v
		default:
			return errors.New("top_predictions is not a string")
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
		case time.Time:
			t := v.Format(time.RFC3339)
			a.Date = &t
		case int64:
			t := time.Unix(v/1000, (v%1000)*int64(time.Millisecond))
			tStr := t.Format(time.RFC3339)
			a.Date = &tStr
		case float64:
			t := time.Unix(int64(v)/1000, (int64(v)%1000)*int64(time.Millisecond))
			tStr := t.Format(time.RFC3339)
			a.Date = &tStr
		case primitive.DateTime:
			t := v.Time()
			tStr := t.Format(time.RFC3339)
			a.Date = &tStr
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
