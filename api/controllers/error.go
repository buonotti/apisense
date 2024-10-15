package controllers

type ErrorResponse struct {
	Message string `json:"message"`
}

func Err(err error) ErrorResponse {
	return ErrorResponse{Message: err.Error()}
}
