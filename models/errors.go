package models

import "errors"

var ErrorInternalServerError = errors.New("internal server error")

type ErrorResponseModel struct {
	Message string `json:"message" example:"invalid"`
}

var ErrorInternalServerErrorResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorInternalServerError.Error(),
}
