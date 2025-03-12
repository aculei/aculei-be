package models

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	animals := ctx.QueryArray("animal")
	moonPhases := ctx.QueryArray("moon_phase")
	temperatures := ctx.QueryArray("temperature")
	dates := ctx.QueryArray("date")

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
			var lowerTemp float64
			var upperTemp float64

			if (*f.Temperatures)[0] > (*f.Temperatures)[1] {
				lowerTemp = (*f.Temperatures)[1]
				upperTemp = (*f.Temperatures)[0]
			} else {
				lowerTemp = (*f.Temperatures)[0]
				upperTemp = (*f.Temperatures)[1]
			}

			filter = append(filter, bson.E{Key: "temperature", Value: bson.D{{Key: "$gte", Value: lowerTemp}}})
			filter = append(filter, bson.E{Key: "temperature", Value: bson.D{{Key: "$lte", Value: upperTemp}}})
		}
	}

	if f.Dates != nil {
		if len(*f.Dates) != 2 {
			return bson.D{}, ErrorDatesNotTwo
		}

		datePattern := "02/01/2006"
		date1, err1 := time.Parse(datePattern, (*f.Dates)[0])
		date2, err2 := time.Parse(datePattern, (*f.Dates)[1])

		if err1 != nil || err2 != nil {
			return bson.D{}, ErrorDateFormat
		}

		fromDate := date1
		toDate := date2
		if date1.After(date2) {
			toDate = date1
			fromDate = date2
		}

		mongoDateFrom := primitive.NewDateTimeFromTime(fromDate)
		mongoDateTo := primitive.NewDateTimeFromTime(toDate)

		filter = append(filter, bson.E{
			Key: "date",
			Value: bson.D{
				{Key: "$gte", Value: mongoDateFrom},
				{Key: "$lt", Value: mongoDateTo},
			},
		})
	}

	return filter, nil
}
