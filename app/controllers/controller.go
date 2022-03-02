package controllers

import "GIG/app/types"

func BuildErrorResponse(err error, status int) types.Response {
	return types.Response{
		PayLoad: err.Error(),
		Status:  status,
	}
}

func BuildSuccessResponse(payload interface{}, status int) types.Response {
	return types.Response{
		PayLoad: payload,
		Status:  status,
	}
}
