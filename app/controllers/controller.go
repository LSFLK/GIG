package controllers

// ControllerResponse
//
// The generic response model of the API
//
// swagger:model
type ControllerResponse struct {
	ErrorMessage string      `json:"message"`
}

func BuildErrResponse(err error) ControllerResponse {
	return ControllerResponse{
		ErrorMessage: err.Error(),
	}
}