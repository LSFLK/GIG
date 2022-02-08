package controllers

// swagger:model
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func BuildErrResponse(err error, status int) Response {
	return Response{
		Message: err.Error(),
		Status:  status,
	}
}

func BuildSuccessResponse(message string, status int) Response {
	return Response{
		Message: message,
		Status:  status,
	}
}
