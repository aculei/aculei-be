package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Filter struct {
	Name   string         `json:"name"`
	Values *[]interface{} `json:"values,omitempty"`
	From   interface{}    `json:"from,omitempty"`
	To     interface{}    `json:"to,omitempty"`
}

type FilterGroup struct {
	Animals      *[]string  `json:"animals"`
	MoonPhases   *[]string  `json:"moon_phases"`
	Temperatures *[]float64 `json:"temperatures"`
	Dates        *[]string  `json:"dates"`
}

func BuildFilterGroup(ctx *gin.Context) (*FilterGroup, error) {
	f := FilterGroup{}

	animals := ctx.QueryArray("animals")
	moonPhases := ctx.QueryArray("moon_phases")
	temperatures := ctx.QueryArray("temperatures")
	dates := ctx.QueryArray("dates")

	if len(animals) > 0 {
		f.Animals = &animals
	}

	if len(moonPhases) > 0 {
		f.MoonPhases = &moonPhases
	}

	temperaturesConverted := make([]float64, len(temperatures))
	if len(temperatures) > 0 {
		for i, t := range temperatures {
			conv, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return nil, ErrorInvalidTemperatureValues
			}
			temperaturesConverted[i] = conv
		}
		f.Temperatures = &temperaturesConverted
	}

	if len(dates) > 0 {
		if len(dates) > 2 {
			return nil, ErrorTooManyDates
		}
		f.Dates = &dates
	}

	return &f, nil
}

func (f *FilterGroup) GenerateFilters() (bson.D, error) {
	filter := bson.D{}

	if f.Animals != nil {
		filter = append(filter, bson.E{Key: "predicted_animal", Value: bson.D{{Key: "$in", Value: *f.Animals}}})
	}

	if f.MoonPhases != nil {
		filter = append(filter, bson.E{Key: "moon_phase", Value: bson.D{{Key: "$in", Value: *f.MoonPhases}}})
	}

	if f.Temperatures != nil {
		if len(*f.Temperatures) == 1 {
			filter = append(filter, bson.E{Key: "temperature", Value: bson.D{{Key: "$eq", Value: (*f.Temperatures)[0]}}})
		} else {
			if (*f.Temperatures)[0] >= (*f.Temperatures)[1] {
				return bson.D{}, ErrroInvalidFromToTemperature
			}

			filter = append(filter, bson.E{Key: "temperature", Value: bson.D{{Key: "$gt", Value: (*f.Temperatures)[0]}}})
			filter = append(filter, bson.E{Key: "temperature", Value: bson.D{{Key: "$lt", Value: (*f.Temperatures)[1]}}})
		}
	}

	if f.Dates != nil {
		if len(*f.Dates) == 1 {
			parsedDate, err := time.Parse("2006-01-02", (*f.Dates)[0])
			if err != nil {
				return bson.D{}, errors.New("invalid date format in filters")
			}
			filter = append(filter, bson.E{Key: "date", Value: bson.D{{Key: "$eq", Value: parsedDate}}})
		} else {
			fromDate, err1 := time.Parse("2006-01-02", (*f.Dates)[0])
			toDate, err2 := time.Parse("2006-01-02", (*f.Dates)[1])

			if err1 != nil || err2 != nil {
				return bson.D{}, errors.New("invalid date format in filters")
			}

			if fromDate.After(toDate) {
				return bson.D{}, errors.New("invalid date range: from >= to")
			}

			filter = append(filter, bson.E{
				Key: "date",
				Value: bson.D{
					{Key: "$gte", Value: fromDate},
					{Key: "$lt", Value: toDate},
				},
			})
		}
	}

	return filter, nil
}
