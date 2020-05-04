package controllers

// ErrorResponse
//
// The generic response model of the API
//
// swagger:model
type ErrorResponse struct {
	ErrorMessage string      `json:"error"`
}

func BuildErrResponse(err error) ErrorResponse {
	return ErrorResponse{
		ErrorMessage: err.Error(),
	}
}