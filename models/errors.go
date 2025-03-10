package models

import "errors"

var ErrorInternalServerError = errors.New("internal server error")
var ErrorBadRequest = errors.New("bad request")
var ErrorInvalidFilterValues = errors.New("invalid filter values")

// Filter errors
type ErrorFilters error

var ErrorTooManyDates ErrorFilters = errors.New("invalid date range in filters: too many dates")
var ErrorInvalidTemperatureValues ErrorFilters = errors.New("invalid temperature value in filters: couldn't convert to float")
var ErrroInvalidFromToTemperature ErrorFilters = errors.New("invalid temperature range in filters: from >= to")

type ErrorResponseModel struct {
	Message string `json:"message" example:"invalid"`
	Detail  string `json:"detail,omitempty" example:"detail"`
}

var ErrorInternalServerErrorResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorInternalServerError.Error(),
}

func NewBadRequest(detail string) ErrorResponseModel {
	return ErrorResponseModel{
		Message: ErrorBadRequest.Error(),
		Detail:  detail,
	}
}
