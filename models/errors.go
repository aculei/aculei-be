package models

import "errors"

var ErrorInternalServerError = errors.New("internal server error")

type ErrorResponseModel struct {
	Message string   `json:"message" example:"invalid"`
	Param   string   `json:"param,omitempty" example:"param_name"`
	Params  []string `json:"params" example:"param1,param2"`
}

var ErrorInternalServerErrorResponseModel ErrorResponseModel = ErrorResponseModel{
	Message: ErrorInternalServerError.Error(),
}
