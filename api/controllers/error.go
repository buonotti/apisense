package controllers

// ErrorResponse is the message returned to the client if any type of error happened
type ErrorResponse struct {
	Message string `json:"message"` // Holds the information about what happened
}

// Err is a helper function to convert an `error` to an `ErrorResponse`
func Err(err error) ErrorResponse {
	return ErrorResponse{Message: err.Error()}
}
