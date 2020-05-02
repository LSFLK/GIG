package controllers

// ControllerResponse
//
// The generic response model of the API
//
// swagger:model
type ControllerResponse struct {
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Content      interface{} `json:"content"`
}

func BuildErrResponse(status int, err error) ControllerResponse {
	return ControllerResponse{
		Status:       status,
		ErrorMessage: err.Error(),
	}
}

func BuildResponse(status int, content interface{}) ControllerResponse {
	return ControllerResponse{
		Status:  status,
		Content: content,
	}
}
