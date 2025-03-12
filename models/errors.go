package models

import (
	"errors"
)

var ErrorInternalServerError = errors.New("internal server error")
var ErrorBadRequest = errors.New("bad request")
var ErrorInvalidFilterValues = errors.New("invalid filter values")

type ErrorFilter struct {
	msg string
}

func (e *ErrorFilter) Error() string {
	return e.msg
}

func NewErrorFilter(msg string) *ErrorFilter {
	return &ErrorFilter{msg: msg}
}

var (
	ErrorInvalidTemperatureValues = NewErrorFilter("filters errror: couldn't convert temperature to float")
	ErrorTooManyDates             = NewErrorFilter("filters error: too many dates")
	ErrorDatesNotTwo              = NewErrorFilter("filters error: two dates required")
	ErrorDateFormat               = NewErrorFilter("filters error: dates format incorrect, expected is dd/mm/yyyy")
)

type ErrorDatabase struct {
	msg string
}

func (e *ErrorDatabase) Error() string {
	return e.msg
}

func NewErrorDatabase(msg string) *ErrorDatabase {
	return &ErrorDatabase{msg: msg}
}

var (
	ErrorDatabaseFind         = NewErrorDatabase("database error: couldn't complete find operation")
	ErrorDatabaseImageDecoder = NewErrorDatabase("database error: couldn't decode image")
	ErrorDatabaseCursor       = NewErrorDatabase("database error: cursor error")
	ErrorDatabaseCount        = NewErrorDatabase("database error: couldn't complete count operation")
	ErrorDatabaseAggregate    = NewErrorDatabase("database error: couldn't complete aggregate operation")
)

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
