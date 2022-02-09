package controllers

// swagger:model
type Response struct {
	PayLoad interface{} `json:"payload"`
	Status  int    `json:"status"`
}

func BuildErrResponse(err error, status int) Response {
	return Response{
		PayLoad: err.Error(),
		Status:  status,
	}
}

func BuildSuccessResponse(payload interface{}, status int) Response {
	return Response{
		PayLoad: payload,
		Status:  status,
	}
}
