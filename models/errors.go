package models

var InternalServerErrorMessage = "Internal server error"

type Error struct {
	Message string `json:"message" example:"error message"`
}

type ErrorInternalServerError struct {
	Message string `json:"message" example:"internal server error"`
}

type ErrorResponseModel struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func InternalServerError() {

}
