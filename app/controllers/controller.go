package controllers

// ControllerResponse
//
// The generic response model of the API
//
// swagger:model
type ControllerResponse struct {
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message"`
}

func BuildErrResponse(status int, err error) ControllerResponse {
	return ControllerResponse{
		Status:       status,
		ErrorMessage: err.Error(),
	}
}