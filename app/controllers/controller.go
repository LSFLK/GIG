package controllers

type Controller map[string]interface{}

func BuildErrResponse( status int, err error) Controller {
	ctrlErr := Controller{}
	ctrlErr["error_message"] = err.Error()
	ctrlErr["status"] = status
	return ctrlErr
}

func BuildResponse(status int, content interface{}) Controller {
	ctrlResp := Controller{}
	ctrlResp["status"] = status
	ctrlResp["content"] = content
	return ctrlResp
}
